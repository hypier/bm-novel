package draft

import (
	"bm-novel/internal/domain/novel/chapter"
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

var (
	patternNumber = `([零一二两三四五六七八九十百千万亿壹贰叁肆伍陆柒捌玖拾佰仟1-9]+)`
	patternUnit   = `([集章回话节 、])`
	patternTitle  = `(.*)`
	// PatternChapterOnlyNo 只有序号
	PatternChapterOnlyNo = fmt.Sprintf(`(?:^|\n)%s\n`, patternNumber)
	// PatternChapter 章节匹配
	PatternChapter = fmt.Sprintf(`(?:^|\n)第%s%s%s\n`, patternNumber, patternUnit, patternTitle)
	// PatternChapterNoTitle 章节没有标题
	PatternChapterNoTitle = fmt.Sprintf(`(?:^|\n)第%s%s\n`, patternNumber, patternUnit)
	// PatternChapterWithVolume 章节带卷
	PatternChapterWithVolume = fmt.Sprintf(`(?:^|\n)第%s卷.+%s.+%s%s\n`, patternNumber, patternNumber, patternUnit, patternTitle)
	// PatternChapterNoTitleWithVolume 章节带卷
	PatternChapterNoTitleWithVolume = fmt.Sprintf(`(?:^|\n)第%s卷.+%s.+%s\n`, patternNumber, patternNumber, patternUnit)
	// PatternAll 全部匹配
	PatternAll = fmt.Sprintf("%s|%s|%s|%s|%s", PatternChapterOnlyNo, PatternChapter, PatternChapterNoTitle, PatternChapterWithVolume, PatternChapterNoTitleWithVolume)
	// PatternParagraph 段落匹配
	PatternParagraph = `[“"]|”(?:[。\.]\n)?`
)

func chapterPositionAll(data []byte) (position, error) {
	pos := regexp.MustCompile(PatternAll).FindIndex(data)
	if len(pos) < 2 {
		// 没有匹配到章节内容
		return *null(), ErrNotMatched
	}

	return position{pos[0], pos[1]}, nil
}

func chapterParserAll(dec *bytes.Buffer) (*chapter.Chapter, error) {
	c := &chapter.Chapter{}

	s2 := regexp.MustCompile(PatternAll)
	all := s2.FindSubmatch(dec.Bytes())
	if all == nil || len(all) < 2 {
		return c, ErrNotMatched
	}

	if index, ok := cNumberToInt(string(all[1])); ok {
		c.ChapterNo = index
	} else {
		return c, ErrNotMatched
	}

	if len(all) > 3 {
		c.ChapterTitle = strings.TrimSpace(string(all[3]))
	} else {
		c.ChapterTitle = "<未命名>"
	}

	return c, nil
}

func chapterPositionOnlyNo(data []byte) (position, error) {
	pos := regexp.MustCompile(PatternChapterOnlyNo).FindIndex(data)
	if len(pos) < 2 {
		// 没有匹配到章节内容
		return *null(), ErrNotMatched
	}

	return position{pos[0], pos[1]}, nil
}

func chapterPosition(data []byte) (position, error) {
	pos := regexp.MustCompile(PatternChapter).FindIndex(data)
	if len(pos) < 2 {
		// 没有匹配到章节内容
		return *null(), ErrNotMatched
	}

	return position{pos[0], pos[1]}, nil
}

func chapterPositionNoTitle(data []byte) (position, error) {
	pos := regexp.MustCompile(PatternChapterNoTitle).FindIndex(data)
	if len(pos) < 2 {
		// 没有匹配到章节内容
		return *null(), ErrNotMatched
	}

	return position{pos[0], pos[1]}, nil
}

func chapterPositionWithVolume(data []byte) (position, error) {
	pos := regexp.MustCompile(PatternChapterWithVolume).FindIndex(data)
	if len(pos) < 2 {
		// 没有匹配到章节内容
		return *null(), ErrNotMatched
	}

	return position{pos[0], pos[1]}, nil
}

func chapterPositionNoTitleWithVolume(data []byte) (position, error) {
	pos := regexp.MustCompile(PatternChapterNoTitleWithVolume).FindIndex(data)
	if len(pos) < 2 {
		// 没有匹配到章节内容
		return *null(), ErrNotMatched
	}

	return position{pos[0], pos[1]}, nil
}

func chapterParserOnlyNo(dec *bytes.Buffer) (*chapter.Chapter, error) {
	c := &chapter.Chapter{}

	s2 := regexp.MustCompile(PatternChapterOnlyNo)
	all := s2.FindSubmatch(dec.Bytes())
	if all == nil || len(all) < 2 {
		return c, ErrNotMatched
	}

	if index, ok := cNumberToInt(string(all[1])); ok {
		c.ChapterNo = index
	} else {
		return c, ErrNotMatched
	}

	c.ChapterTitle = "<未命名>"

	return c, nil
}

func chapterParser(dec *bytes.Buffer) (*chapter.Chapter, error) {
	c := &chapter.Chapter{}

	s2 := regexp.MustCompile(PatternChapter)
	all := s2.FindSubmatch(dec.Bytes())
	if all == nil || len(all) < 2 {
		return c, ErrNotMatched
	}

	if index, ok := cNumberToInt(string(all[1])); ok {
		c.ChapterNo = index
	} else {
		return c, ErrNotMatched
	}

	if len(all) > 3 {
		c.ChapterTitle = strings.TrimSpace(string(all[3]))
	} else {
		c.ChapterTitle = "<未命名>"
	}

	return c, nil
}

func chapterParserNoTitle(dec *bytes.Buffer) (*chapter.Chapter, error) {
	c := &chapter.Chapter{}

	s2 := regexp.MustCompile(PatternChapterNoTitle)
	all := s2.FindSubmatch(dec.Bytes())
	if all == nil || len(all) < 2 {
		return c, ErrNotMatched
	}

	if index, ok := cNumberToInt(string(all[1])); ok {
		c.ChapterNo = index
	} else {
		return c, ErrNotMatched
	}

	c.ChapterTitle = "<未命名>"

	return c, nil
}

func chapterParserWithVolume(dec *bytes.Buffer) (*chapter.Chapter, error) {
	c := &chapter.Chapter{}

	s2 := regexp.MustCompile(PatternChapterWithVolume)
	all := s2.FindSubmatch(dec.Bytes())
	if all == nil || len(all) < 2 {
		return c, ErrNotMatched
	}

	if index, ok := cNumberToInt(string(all[1])); ok {
		c.Volume = index
	} else {
		return c, ErrNotMatched
	}

	if index, ok := cNumberToInt(string(all[2])); ok {
		c.ChapterNo = index
	} else {
		return c, ErrNotMatched
	}

	if len(all) > 3 {
		c.ChapterTitle = strings.TrimSpace(string(all[4]))
	} else {
		c.ChapterTitle = "<未命名>"
	}

	return c, nil
}

func chapterParserNoTitleWithVolume(dec *bytes.Buffer) (*chapter.Chapter, error) {
	c := &chapter.Chapter{}

	s2 := regexp.MustCompile(PatternChapterNoTitleWithVolume)
	all := s2.FindSubmatch(dec.Bytes())
	if all == nil || len(all) < 2 {
		return c, ErrNotMatched
	}

	if index, ok := cNumberToInt(string(all[1])); ok {
		c.ChapterNo = index
	} else {
		return c, ErrNotMatched
	}

	if len(all) > 3 {
		c.ChapterTitle = strings.TrimSpace(string(all[3]))
	} else {
		c.ChapterTitle = "<未命名>"
	}

	return c, nil
}
