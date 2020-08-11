package draft

import (
	"bm-novel/internal/domain/novel/chapter"
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

var (
	// 未解决数字中间有空格问题
	patternNumber = `([零一二两三四五六七八九十百千万亿壹贰叁肆伍陆柒捌玖拾佰仟1-9]+)`
	patternUnit   = `(?:[集章回话节 、])`
	patternTitle  = `(.*)`

	// ChapterPatternAll 全部匹配
	ChapterPatternAll = fmt.Sprintf(`\n[^“]??(?:%s卷.*?)?%s%s?%s\n`, patternNumber, patternNumber, patternUnit, patternTitle)
	// ChapterPatternHead 首行匹配
	ChapterPatternHead = fmt.Sprintf(`^[^“]??(?:%s卷.*?)?%s%s?%s\n`, patternNumber, patternNumber, patternUnit, patternTitle)
	// PatternParagraph 段落匹配
	PatternParagraph = `[“"]|”(?:[。\.]\n)?`
)

func chapterPositionAll(data []byte, pattern string) (position, error) {

	pos := regexp.MustCompile(pattern).FindIndex(data)
	if len(pos) < 2 {
		// 没有匹配到章节内容
		return *null(), ErrNotMatched
	}

	return position{pos[0], pos[1]}, nil
}

func chapterParserAll(dec *bytes.Buffer, pattern string) (*chapter.Chapter, error) {
	c := &chapter.Chapter{}

	s2 := regexp.MustCompile(pattern)
	all := s2.FindSubmatch(dec.Bytes())
	if all == nil || len(all) < 2 {
		return nil, ErrNotMatched
	}

	if all[1] != nil {
		if index, ok := cNumberToInt(string(all[1])); ok {
			c.Volume = index
		}
	}

	if all[2] != nil {
		if index, ok := cNumberToInt(string(all[2])); ok {
			c.ChapterNo = index
		}
	}

	if all[3] != nil {
		c.ChapterTitle = strings.TrimSpace(string(all[3]))
	} else {
		c.ChapterTitle = "<未命名>"
	}

	return c, nil
}
