package draft

import (
	"bm-novel/internal/domain/novel/chapter"
	nc "bm-novel/internal/domain/novel/counter"
	"bm-novel/internal/domain/novel/paragraph"
	"io"
	"os"
	"testing"

	"github.com/sirupsen/logrus"

	uuid "github.com/satori/go.uuid"
)

func openFile() io.Reader {
	//f, _ := os.Open("C:\\Users\\yuepaidui20200612\\iCloudDrive\\Documents\\工作\\joyparty\\北冥有声\\庆余年-已完结.txt")
	//f, _ := os.Open("C:\\Users\\yuepaidui20200612\\iCloudDrive\\Documents\\工作\\joyparty\\北冥有声\\002.txt")
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

func BenchmarkDraft_Parser(b *testing.B) {
	logrus.SetLevel(logrus.DebugLevel)
	draft := &Draft{}
	c := &nc.NovelCounter{NovelID: uuid.NewV4(), CountID: uuid.NewV4()}
	draft.Parser(c, openFile())

	b.Log(len(draft.Chapters))
}

func TestDraft_Parser1(t *testing.T) {
	counter := &nc.NovelCounter{NovelID: uuid.NewV4(), CountID: uuid.NewV4()}
	type fields struct {
		Paragraphs *paragraph.Paragraphs
		Chapters   chapter.Chapters
		Counter    *nc.NovelCounter
		isChapter  bool
		pCounter   func() int
	}
	type args struct {
		counter *nc.NovelCounter
		//file    io.Reader
		fineName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   struct {
			wordsCount      int
			chaptersCount   int
			paragraphsCount int
		}
	}{
		{
			name: "间谍的战争",
			args: args{fineName: "C:\\Users\\yuepaidui20200612\\iCloudDrive\\Documents\\工作\\joyparty\\北冥有声\\间谍的战争-已完结.txt",
				counter: counter},
			want: struct {
				wordsCount      int
				chaptersCount   int
				paragraphsCount int
			}{wordsCount: 3140187, chaptersCount: 1475, paragraphsCount: 74928},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Draft{
				Paragraphs: tt.fields.Paragraphs,
				Chapters:   tt.fields.Chapters,
				Counter:    tt.fields.Counter,
				isChapter:  tt.fields.isChapter,
				pCounter:   tt.fields.pCounter,
			}

			f, _ := os.Open(tt.args.fineName)
			defer f.Close()

			d.Parser(tt.args.counter, f)

			if tt.want.wordsCount != d.Counter.WordsCount {
				t.Errorf("draft.Parser WordsCount %v, want %v", d.Counter.WordsCount, tt.want.wordsCount)
			}

			if tt.want.chaptersCount != d.Counter.ChaptersCount {
				t.Errorf("draft.Parser ChaptersCount %v, want %v", d.Counter.ChaptersCount, tt.want.chaptersCount)
			}

			paragraphsCount := len(*d.Paragraphs)
			if tt.want.paragraphsCount != paragraphsCount {
				t.Errorf("draft.Parser paragraphsCount %v, want %v", paragraphsCount, tt.want.paragraphsCount)
			}
		})
	}
}
