package draft

import (
	nc "bm-novel/internal/domain/novel/counter"
	"io"
	"os"
	"testing"

	"github.com/sirupsen/logrus"

	uuid "github.com/satori/go.uuid"
)

func openFile() io.Reader {
	//f, _ := os.Open("C:\\Users\\yuepaidui20200612\\iCloudDrive\\Documents\\工作\\joyparty\\北冥有声\\庆余年-已完结.txt")
	f, _ := os.Open("C:\\Users\\yuepaidui20200612\\iCloudDrive\\Documents\\工作\\joyparty\\北冥有声\\间谍的战争-已完结.txt")
	//f, _ := os.Open("/Users/barry/go/src/bm-novel/docs/间谍的战争-已完结.txt")
	//defer f.Close()
	return f
}

func TestDraft_Parser(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	draft := &Draft{}
	c := &nc.NovelCounter{NovelID: uuid.NewV4(), CountID: uuid.NewV4()}
	draft.Parser(c, openFile())

	t.Log(len(draft.Chapters))
}
