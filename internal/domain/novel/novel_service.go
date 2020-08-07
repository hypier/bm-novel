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

func counter(value, step int) func() int {
	num := value
	return func() int {
		num += step
		return num
	}
}

var pCounter func() int

func findChapterLine(data []byte) (begin, end int) {
	p1 := `([零一二两三四五六七八九十百千万亿壹贰叁肆伍陆柒捌玖拾佰仟1-9]+)([集章回话节 、])([\w\W].*)`
	pos := regexp.MustCompile(p1).FindSubmatchIndex(data)
	if len(pos) == 0 {
		return 0, 0
	}

	return pos[0], pos[1]
}

// UploadDraft 上传原文
func (s Service) UploadDraft(ctx context.Context, novelID uuid.UUID, file io.Reader) error {
	logrus.Debug("小说解析开始")
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
	scanner := bufio.NewScanner(r)
	buf := make([]byte, 5000)
	scanner.Buffer(buf, bufio.MaxScanTokenSize)

	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		begin, end := findChapterLine(data)

		s2 := regexp.MustCompile(`[“”"]`)
		pos := s2.FindAllIndex(data, 2)

		if end > 0 {

			if len(pos) > 0 {
				if begin < pos[0][0] {
					return end, data[0:end], nil
				} else {
					if len(pos) < 1 {
						if begin < pos[1][0] {
							return begin, data[0:begin], nil
						}
						return begin, data[0:begin], nil
					}

				}
			}

		}

		// todo 半个引号未处理
		if len(pos) > 1 {
			index := pos[0][0]

			// 首字为引号
			if index <= 3 {
				index = pos[1][1]
			}

			return index, data[0:index], nil
		}

		if !atEOF && !utf8.FullRune(data) {
			// Incomplete; get more bytes.
			return 0, nil, nil
		}

		return 0, nil, nil
	}
	scanner.Split(split)

	pCounter = counter(0, 1)
	cs := &chapter.Chapters{}
	ps := &paragraph.Paragraphs{}
	var c *chapter.Chapter
	for scanner.Scan() {
		dec := bytes.NewBuffer(scanner.Bytes())

		if len(bytes.TrimSpace(dec.Bytes())) == 0 {
			continue
		}

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

	logrus.Debug("小说解析完成")
	return nil
}

func parseParagraph(dec *bytes.Buffer, ps *paragraph.Paragraphs, c *chapter.Chapter) {
	if c == nil {
		return
	}

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
			Index:       pCounter(),
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

func parseChapterNoVolume2(dec *bytes.Buffer, c *chapter.Chapter) bool {
	p2 := `([零一二两三四五六七八九十百千万亿壹贰叁肆伍陆柒捌玖拾佰仟1-9]*)([\w\W].*)`

	s2 := regexp.MustCompile(p2)

	all := s2.FindSubmatch(dec.Bytes())
	if all == nil || len(all) != 2 {
		return false
	}

	if index, ok := cNumberToInt(string(all[1])); ok {
		c.ChapterNo = index
	}

	c.ChapterTitle = strings.TrimSpace(string(all[2]))

	return true
}

func isChapter(dec *bytes.Buffer) bool {
	s1 := regexp.MustCompile(`^[\w 　]`)
	if s1.Match(dec.Bytes()) || len(dec.Bytes()) == 0 {
		return false
	}

	p2 := `([零一二两三四五六七八九十百千万亿壹贰叁肆伍陆柒捌玖拾佰仟1-9]+).+[集章回话节 、]([\w\W].*)`
	s1 = regexp.MustCompile(p2)
	if s1.Match(dec.Bytes()) {
		return true
	}

	p2 = `([零一二两三四五六七八九十百千万亿壹贰叁肆伍陆柒捌玖拾佰仟1-9]+).+`
	s1 = regexp.MustCompile(p2)
	if s1.Match(dec.Bytes()) {
		return true
	}

	return false
}

func parseChapter(dec *bytes.Buffer, c *chapter.Chapter) bool {

	if parseChapterWithVolume(dec, c) {
		return true
	}

	if parseChapterNoVolume(dec, c) {
		return true
	}

	if parseChapterNoVolume2(dec, c) {
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
