package draft

import (
	"io"
	"os"
	"testing"
)

func openFile() io.Reader {
	//f, _ := os.Open("C:\\Users\\yuepaidui20200612\\iCloudDrive\\Documents\\工作\\joyparty\\北冥有声\\间谍的战争-已完结.txt")
	f, _ := os.Open("/Users/barry/go/src/bm-novel/docs/间谍的战争-已完结.txt")
	//defer f.Close()
	return f
}

func TestDraft_Parser(t *testing.T) {
	draft := &Draft{}

	draft.Parser(openFile())
}
