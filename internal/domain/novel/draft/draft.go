package draft

import (
	"bm-novel/internal/domain/novel/chapter"
	nc "bm-novel/internal/domain/novel/counter"
	"bm-novel/internal/domain/novel/paragraph"
	"bufio"
	"bytes"
	"errors"
	"io"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/sirupsen/logrus"

	uuid "github.com/satori/go.uuid"
)

var (
	// ErrNoPosition 没有切分位置
	ErrNoPosition = errors.New("no position")
	// ErrNotMatched 没有匹配内容
	ErrNotMatched = errors.New("not matched")
)

// Draft 草稿
type Draft struct {
	Paragraphs *paragraph.Paragraphs
	Chapters   chapter.Chapters
	Counter    *nc.NovelCounter

	isChapter bool
	pCounter  func() int
}

func (d *Draft) getLastChapter() *chapter.Chapter {
	if d.Chapters == nil {
		// 初始化
		return &chapter.Chapter{
			NovelID:      d.Counter.NovelID,
			ChapterID:    uuid.NewV4(),
			ChapterTitle: "<未命名>"}
	}

	return d.Chapters[len(d.Chapters)-1]
}

func (d *Draft) getSplitPosition(cp position, pp positions) (int, error) {

	if cp.isNull() {
		// 没有章节匹配
		d.isChapter = false
		return pp.getSpitePos()
	}

	if pp.head.isNull() {
		d.isChapter = true
		return cp.end, nil
	}

	i := compare(cp, pp.head)
	if i <= 0 {
		// 章节在前
		d.isChapter = true
		return cp.end, nil
	}

	i = compare(cp, pp.tail)
	if i <= 0 {
		// 章节位于中间
		d.isChapter = false
		return cp.begin, nil
	}

	// 章节在后
	d.isChapter = false
	return pp.getSpitePos()
}

// 可提取匹配表达式
func (d *Draft) chapterPosition(data []byte) position {

	if cp, err := chapterPositionAll(data); err == nil {
		return cp
	} else if !errors.Is(err, ErrNotMatched) {
		return *null()
	}

	return *null()
}

func (d *Draft) paragraphPosition(data []byte) positions {
	pos := regexp.MustCompile(PatternParagraph).FindAllIndex(data, 2)
	if len(pos) < 2 {
		return positions{*null(), *null()}
	}

	return positions{
		head: position{pos[0][0], pos[0][1]},
		tail: position{pos[1][0], pos[1][1]},
	}
}

func (d *Draft) split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	cp := d.chapterPosition(data)
	pp := d.paragraphPosition(data)

	if pos, err := d.getSplitPosition(cp, pp); err == nil {
		return pos, data[0:pos], nil
	}

	return 0, nil, nil
}

// Parser 解析
func (d *Draft) Parser(counter *nc.NovelCounter, file io.Reader) {
	d.Counter = counter
	r := bufio.NewReader(file)

	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 5000), bufio.MaxScanTokenSize)
	scanner.Split(d.split)

	for scanner.Scan() {
		dec := bytes.NewBuffer(scanner.Bytes())
		content := dec.Bytes()

		if len(bytes.TrimSpace(content)) == 0 {
			continue
		}

		if d.isChapter {
			if c, err := d.parseChapter(dec); err == nil {
				logrus.Debugf("volume: %d, no: %d, title: %s ", c.Volume, c.ChapterNo, c.ChapterTitle)
				d.addChapter(c)
				continue
			}
		} else {
			p := d.parseParagraph(dec)
			d.addParagraph(p)
		}

	}
}

func (d *Draft) addChapter(c *chapter.Chapter) {
	if d.Chapters == nil {
		d.Chapters = chapter.Chapters{}
	}

	c.ChapterID = uuid.NewV4()
	c.NovelID = d.Counter.NovelID
	d.Counter.ChaptersCount++
	d.Chapters = append(d.Chapters, c)
}

func (d *Draft) addParagraph(p *paragraph.Paragraph) {
	if d.Paragraphs == nil {
		d.Paragraphs = &paragraph.Paragraphs{}
		d.pCounter = counter(d.Counter.ChaptersCount, 1)
	}

	p.ParagraphID = uuid.NewV4()
	p.ChapterIndex = d.pCounter()

	p.WordsCount = utf8.RuneCountInString(p.Content)
	d.getLastChapter().WordsCount += p.WordsCount
	d.Counter.WordsCount += p.WordsCount

	p.NovelID = d.Counter.NovelID
	*d.Paragraphs = append(*d.Paragraphs, p)
}

// 可提取匹配表达式，
func (d *Draft) parseChapter(dec *bytes.Buffer) (*chapter.Chapter, error) {

	if cp, err := chapterParserAll(dec); err == nil {
		return cp, nil
	} else if !errors.Is(err, ErrNotMatched) {
		return nil, err
	}
	return nil, ErrNotMatched
}

func (d *Draft) parseParagraph(dec *bytes.Buffer) *paragraph.Paragraph {
	p := &paragraph.Paragraph{
		Content: strings.TrimSpace(dec.String())}
	return p
}
