package novel

import (
	"bm-novel/internal/config"
	"bm-novel/internal/domain/novel/paragraph"
	pr "bm-novel/internal/infrastructure/persistence/paragraph"
	"bm-novel/internal/infrastructure/postgres"
	"bufio"
	"bytes"
	"context"
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

func init() {
	config.LoadConfigForTest()
	postgres.InitDB()
}

func openFile() io.Reader {
	//f, _ := os.Open("C:\\Users\\yuepaidui20200612\\iCloudDrive\\Documents\\工作\\joyparty\\北冥有声\\间谍的战争-已完结.txt")
	f, _ := os.Open("/Users/barry/go/src/bm-novel/docs/间谍的战争-已完结.txt")
	//defer f.Close()
	return f
}

func TestService_UploadDraft(t *testing.T) {

}

func TestService_UploadDraft2(t *testing.T) {
	p := &paragraph.Paragraph{NovelID: uuid.NewV4()}
	ps := &paragraph.Paragraphs{}
	*ps = append(*ps, p)

	ctx, _ := context.WithCancel(context.Background())
	err := pr.New().BatchCreate(ctx, ps)

	fmt.Println(err)
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
