package draft

import (
	"bm-novel/internal/domain/novel"
	"bm-novel/internal/domain/novel/chapter"
	"bm-novel/internal/domain/novel/paragraph"
	"bytes"
)

type Draft struct {
	Paragraphs *paragraph.Paragraphs
	Chapters   *chapter.Chapters
	Novel      *novel.Novel
}

type Parser interface {
	Exec(dec *bytes.Buffer, d *Draft)
}

type ParserFunc func(dec *bytes.Buffer, d *Draft)

func (p ParserFunc) Exec(dec *bytes.Buffer, d *Draft) {
	p.Exec(dec, d)
}
