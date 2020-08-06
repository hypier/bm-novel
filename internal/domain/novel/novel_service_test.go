package novel

import (
	"bm-novel/internal/domain/novel/chapter"
	"bm-novel/internal/domain/novel/paragraph"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"testing"
	"unicode/utf8"

	uuid "github.com/satori/go.uuid"
)

func openFile() io.Reader {
	f, _ := os.Open("C:\\Users\\yuepaidui20200612\\iCloudDrive\\Documents\\工作\\joyparty\\北冥有声\\间谍的战争-已完结.txt")
	//defer f.Close()
	return f
}

func TestService_UploadDraft(t *testing.T) {
	file := openFile()
	r := bufio.NewReader(file)

	novelID := uuid.NewV4()

	cs := &chapter.Chapters{}
	ps := &paragraph.Paragraphs{}
	var c *chapter.Chapter
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}

		dec := bytes.NewBuffer(line)

		if isChapter(dec) {
			c = &chapter.Chapter{NovelID: novelID, ChapterID: uuid.NewV4()}
			// todo 是章节但没有序号的处理
			parseChapter(dec, c)
			*cs = append(*cs, c)
		} else {
			parseParagraph(dec, ps, c)
		}

	}

	for _, p := range *ps {
		fmt.Println(p.ChapterID)
	}
	//fmt.Println(len(*ps))
	//for _, c2 := range *cs {
	//	fmt.Println(c2.ChapterID, c2.ChapterTitle)
	//}
}

func TestService_UploadDraft1(t *testing.T) {

	str := "　　杨逸现在最关心的就是这个，李凡微微一笑，侧身让出了身后的人，道："

	s2 := regexp.MustCompile(`['"”]`)
	str = s2.ReplaceAllString(str, `“`)

	cstr := strings.Split(str, `“`)

	for _, v := range cstr {

		fmt.Println(strings.TrimSpace(v))
	}
}

func BenchmarkService_UploadDraft(b *testing.B) {

	i := 0
	file := openFile()
	b.Run(b.Name(), func(b *testing.B) {
		r := bufio.NewReader(file)

		for {
			line, _, err := r.ReadLine()
			if err != nil {
				break
			}

			i += utf8.RuneCountInString(bytes.NewBuffer(line).String())
		}
	})

	fmt.Println(i)
}

func BenchmarkService_UploadDraft2(b *testing.B) {
	i := 0
	file := openFile()

	b.Run(b.Name(), func(b *testing.B) {
		r := bufio.NewReader(file)

		for {
			line, _, err := r.ReadLine()
			if err != nil {
				break
			}
			i += utf8.RuneCountInString(string(line))
		}
	})

	fmt.Println(i)
}

func BenchmarkService_UploadDraft3(b *testing.B) {
	i := 0
	file := openFile()
	b.Run(b.Name(), func(b *testing.B) {
		all, _ := ioutil.ReadAll(file)

		i += utf8.RuneCountInString(string(all))
	})

	fmt.Println(i)
}
