package draft

import (
	"bm-novel/internal/domain/novel/chapter"
	"bytes"
)

func ChapterParser(next Parser) Parser {
	return ParserFunc(func(dec *bytes.Buffer, d *Draft) {

		next.Exec(dec, d)
	})
}

type Pattern interface {
	Exec(dec *bytes.Buffer, c *chapter.Chapter)
}
