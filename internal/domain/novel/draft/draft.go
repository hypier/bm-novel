package draft

import (
	"bm-novel/internal/domain/novel"
	"bm-novel/internal/domain/novel/chapter"
	"bm-novel/internal/domain/novel/paragraph"
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"
)

var (
	PatternChapter   = `([零一二两三四五六七八九十百千万亿壹贰叁肆伍陆柒捌玖拾佰仟1-9]+)([集章回话节 、])([\w\W].*)\n`
	PatternParagraph = `[“”"]`

	ErrNoPosition = errors.New("no position")
	ErrNotMatched = errors.New("not matched")
)

type Draft struct {
	Paragraphs *paragraph.Paragraphs
	Chapters   *chapter.Chapters
	Novel      *novel.Novel
}

func (d *Draft) getSplitPosition(cp position, pp positions) (int, error) {

	if cp.isNull() {
		// 没有章节匹配
		return pp.getSpitePos()
	}

	i := compare(cp, pp.head)
	if i <= 0 {
		// 章节在前
		return cp.end, nil
	}

	i = compare(cp, pp.tail)
	if i <= 0 {
		// 章节位于中间
		return cp.begin, nil
	}

	// 章节在后
	return pp.getSpitePos()
}

func (d *Draft) chapterPosition(data []byte) position {
	pos := regexp.MustCompile(PatternChapter).FindIndex(data)
	if len(pos) < 2 {
		return *null()
	}

	if pos[0] > 9 {
		return *null()
	}

	return position{pos[0], pos[1]}
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

func (d *Draft) SetNovel(novel *novel.Novel) {
	// 确认当前章数
}

func (d *Draft) Parser(file io.Reader) {
	r := bufio.NewReader(file)

	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 5000), bufio.MaxScanTokenSize)
	scanner.Split(d.split)

	for scanner.Scan() {
		dec := bytes.NewBuffer(scanner.Bytes())
		content := dec.Bytes()

		if len(content) == 0 {
			continue
		}

		if strings.Index(string(content), "觉得记性太好也是麻烦") > 0 {
			fmt.Println(11111)
		}

		cp := d.chapterPosition(dec.Bytes())
		if !cp.isNull() {
			if c, err := d.parseChapter(dec); err == nil {
				d.addChapter(c)
				fmt.Println(c.ChapterNo, c.ChapterTitle)
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
		d.Chapters = &chapter.Chapters{}
	}

	*d.Chapters = append(*d.Chapters, c)
}

func (d *Draft) addParagraph(p *paragraph.Paragraph) {
	if d.Paragraphs == nil {
		d.Paragraphs = &paragraph.Paragraphs{}
	}

	*d.Paragraphs = append(*d.Paragraphs, p)
}

func (d *Draft) parseChapter(dec *bytes.Buffer) (*chapter.Chapter, error) {
	c := &chapter.Chapter{}

	s2 := regexp.MustCompile(PatternChapter)
	all := s2.FindSubmatch(dec.Bytes())
	if all == nil || len(all) < 4 {
		return c, ErrNotMatched
	}

	if index, ok := cNumberToInt(string(all[1])); ok {
		c.ChapterNo = index
	} else {
		return c, ErrNotMatched
	}

	c.ChapterTitle = string(all[3])

	return c, nil
}

func (d *Draft) parseParagraph(dec *bytes.Buffer) *paragraph.Paragraph {
	p := &paragraph.Paragraph{
		Content: strings.TrimSpace(dec.String())}
	return p
}
