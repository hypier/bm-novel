package draft

import (
	"bm-novel/internal/domain/novel/chapter"
	"bytes"
	"regexp"
	"strings"
)

var (

	// PatternChapter 章节匹配
	PatternChapter = `(?:^|\n)第([零一二两三四五六七八九十百千万亿壹贰叁肆伍陆柒捌玖拾佰仟1-9]+)([集章回话节 、])(.*)\n`
	// PatternChapterNoTitle 章节没有标题
	PatternChapterNoTitle = `(?:^|\n)第([零一二两三四五六七八九十百千万亿壹贰叁肆伍陆柒捌玖拾佰仟1-9]+)([集章回话节 、])\n`
	// PatternChapterWithVolume 章节带卷
	PatternChapterWithVolume = `(?:^|\n)第([零一二两三四五六七八九十百千万亿壹贰叁肆伍陆柒捌玖拾佰仟1-9]+)卷.+([零一二两三四五六七八九十百千万亿壹贰叁肆伍陆柒捌玖拾佰仟1-9]+).+([集章回话节 、])(.*)\n`
	// PatternChapterNoTitleWithVolume 章节带卷
	PatternChapterNoTitleWithVolume = `(?:^|\n)第([零一二两三四五六七八九十百千万亿壹贰叁肆伍陆柒捌玖拾佰仟1-9]+)卷.+([零一二两三四五六七八九十百千万亿壹贰叁肆伍陆柒捌玖拾佰仟1-9]+).+([集章回话节 、])\n`

	// PatternParagraph 段落匹配
	PatternParagraph = `[“”"]`
)

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
