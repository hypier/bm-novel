package novel

import (
	"bm-novel/internal/domain/novel/chapter"
	"bm-novel/internal/domain/novel/paragraph"
	"bm-novel/internal/http/web"
	"bufio"
	"bytes"
	"context"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/sirupsen/logrus"

	uuid "github.com/satori/go.uuid"
)

// Service 小说服务
type Service struct {
	Repo          INovelRepository
	ChapterRepo   IChapterRepository
	ParagraphRepo IParagraphRepository
}

// Create 创建小说
func (s Service) Create(ctx context.Context, novel *Novel) error {
	title := novel.NovelTitle

	dbNovel, err := s.Repo.FindByTitle(ctx, title)
	if err != nil {
		return err
	}

	if dbNovel != nil {
		return web.WriteErrLogWithField(logrus.Fields{
			"title":     title,
			"dbNovelID": dbNovel.NovelID,
		}, web.ErrConflict, "Create Novel, Duplicate novelTitle")
	}

	novel.NovelID = uuid.NewV4()
	return s.Repo.Create(ctx, novel)
}

// Delete 删除小说
func (s Service) Delete(ctx context.Context, novelID uuid.UUID) error {
	panic("implement me")
}

// AssignResponsibleEditor 指派责编
func (s Service) AssignResponsibleEditor(ctx context.Context, novelID uuid.UUID, editorID uuid.UUID) error {
	dbNovel, err := s.Repo.FindOne(ctx, novelID)

	if err != nil {
		return err
	}

	if dbNovel == nil {
		return web.WriteErrLogWithField(logrus.Fields{
			"novelID": novelID,
		}, web.ErrNotFound, "AssignResponsibleEditor, Novel Not Found")
	}

	dbNovel.ResponsibleEditorID = editorID
	return s.Repo.Update(ctx, dbNovel)

}

// UploadDraft 上传原文
func (s Service) UploadDraft(ctx context.Context, novelID uuid.UUID, file io.Reader) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1200*time.Second)
	defer cancel()

	dbNovel, err := s.Repo.FindOne(ctx, novelID)

	if err != nil {
		return err
	}

	if dbNovel == nil {
		return web.WriteErrLogWithField(logrus.Fields{
			"novelID": novelID,
		}, web.ErrNotFound, "UploadDraft, Novel Not Found")
	}

	r := bufio.NewReader(file)

	cs := &chapter.Chapters{}
	ps := &paragraph.Paragraphs{}
	var c *chapter.Chapter
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}

		dec := bytes.NewBuffer(line)

		if isChapter(dec) {
			c = &chapter.Chapter{NovelID: novelID, ChapterID: uuid.NewV4()}

			if parseChapter(dec, c) {
				*cs = append(*cs, c)
			} else {
				parseParagraph(dec, ps, c)
			}
		} else {
			parseParagraph(dec, ps, c)
		}

	}

	if err = s.ChapterRepo.BatchCreate(ctx, cs); err != nil {
		return err
	}

	if err = s.ParagraphRepo.BatchCreate(ctx, ps); err != nil {
		return err
	}

	return nil
}

func parseParagraph(dec *bytes.Buffer, ps *paragraph.Paragraphs, c *chapter.Chapter) {
	str := dec.String()

	s2 := regexp.MustCompile(`['"”]`)
	str = s2.ReplaceAllString(str, `“`)

	split := strings.Split(str, `“`)

	for _, v := range split {
		content := strings.TrimSpace(v)
		if len(content) == 0 {
			continue
		}

		wordsCount := utf8.RuneCountInString(content)

		p := paragraph.Paragraph{
			ParagraphID: uuid.NewV4(),
			Content:     content,
			WordsCount:  wordsCount,
			ChapterID:   c.ChapterID,
			NovelID:     c.NovelID,
		}

		c.WordsCount += wordsCount
		*ps = append(*ps, &p)
	}

}

func parseChapterWithVolume(dec *bytes.Buffer, c *chapter.Chapter) bool {
	p2 := `([零一二两三四五六七八九十百千万亿壹贰叁肆伍陆柒捌玖拾佰仟1-9]+)卷.+([零一二两三四五六七八九十百千万亿壹贰叁肆伍陆柒捌玖拾佰仟1-9]+).+[集章回话节 、]([\w\W].*)`

	s2 := regexp.MustCompile(p2)
	if !s2.Match(dec.Bytes()) {
		return false
	}

	all := s2.FindSubmatch(dec.Bytes())
	if all == nil || len(all) != 4 {
		return false
	}

	if index, ok := cNumberToInt(string(all[1])); ok {
		c.Volume = index
	}

	if index, ok := cNumberToInt(string(all[2])); ok {
		c.ChapterNo = index
	}

	c.ChapterTitle = string(all[3])

	return true
}

func parseChapterNoVolume(dec *bytes.Buffer, c *chapter.Chapter) bool {
	p2 := `([零一二两三四五六七八九十百千万亿壹贰叁肆伍陆柒捌玖拾佰仟1-9]+).+[集章回话节 、]([\w\W].*)`

	s2 := regexp.MustCompile(p2)

	all := s2.FindSubmatch(dec.Bytes())
	if all == nil || len(all) != 3 {
		return false
	}

	if index, ok := cNumberToInt(string(all[1])); ok {
		c.ChapterNo = index
	}

	c.ChapterTitle = string(all[2])

	return true
}

func isChapter(dec *bytes.Buffer) bool {
	s1 := regexp.MustCompile(`^[\w 　]`)
	if s1.Match(dec.Bytes()) || len(dec.Bytes()) == 0 {
		return false
	}

	p2 := `([零一二两三四五六七八九十百千万亿壹贰叁肆伍陆柒捌玖拾佰仟1-9]+).+[集章回话节 、]([\w\W].*)`
	s1 = regexp.MustCompile(p2)
	if !s1.Match(dec.Bytes()) {
		return false
	}

	return true
}

func parseChapter(dec *bytes.Buffer, c *chapter.Chapter) bool {

	if parseChapterWithVolume(dec, c) {
		return true
	}

	if parseChapterNoVolume(dec, c) {
		return true
	}

	return false
}

func cNumberToInt(s string) (int, bool) {
	// 1.全数字转换
	if numberInt, err := strconv.Atoi(s); err == nil {
		return numberInt, true
	}

	cNum := map[rune]int{'零': 0, '一': 1, '二': 2, '两': 2, '三': 3, '四': 4, '五': 5, '六': 6, '七': 7, '八': 8, '九': 9, '十': 10, '百': 100, '千': 1000, '万': 10000, '亿': 100000000, '壹': 1, '贰': 2, '叁': 3, '肆': 4, '伍': 5, '陆': 6, '柒': 7, '捌': 8, '玖': 9, '拾': 10, '佰': 100, '仟': 1000}
	total, temp := 0, 0

	for i, c := range s {
		// 2.判断是否是全中文
		if !unicode.Is(unicode.Han, c) {
			return 0, false
		}

		val := cNum[c]

		// 3.判断是否是单位
		if val >= 10 {
			// 判断首位是否是单位
			if i == 0 {
				total = val
			} else {
				total += temp * val
			}

			temp = 0
		} else {
			temp = val
		}
	}

	// 未尾数字处理
	if temp > 0 {
		total += temp
	}

	return total, true
}
