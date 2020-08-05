package paragraph

import (
	"bm-novel/internal/config"
	"bm-novel/internal/domain/novel/paragraph"
	"bm-novel/internal/infrastructure/postgres"
	"context"
	"fmt"
	"testing"

	uuid "github.com/satori/go.uuid"
)

func init() {
	config.LoadConfigForTest()
	postgres.InitDB()
}

func TestRepository_Create(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())
	repo := New()

	for i := 20; i < 1000; i += 20 {
		p := &paragraph.Paragraph{ParagraphID: uuid.NewV4(), Index: i, SubIndex: 0}

		err := repo.Create(ctx, p)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type pos struct {
	index    int
	subIndex int
}

func TestInsert(t *testing.T) {
	p := &paragraph.Paragraph{ParagraphID: uuid.NewV4(), Index: 0, SubIndex: 0}
	prev := pos{20, 3}

	// 取大于20的5条,排序取
	ctx, _ := context.WithCancel(context.Background())
	repo := New()
	list, _ := repo.FindList(ctx, prev.index, 5)

	pos, sub := getIndex(list, prev)
	//if pos > prev.index+1{
	//	fmt.Println("要移动", pos)
	//
	//	return
	//}

	p.Index = pos
	p.SubIndex = sub
	_ = repo.Create(ctx, p)

	fmt.Println(pos, sub)
}

func getIndex(list *paragraph.Paragraphs, prev pos) (int, int) {
	end := 0
	sub := 0
	for _, x := range *list {
		n := prev.index + 1
		if x.Index > n {
			end = n
			break
		} else if x.Index == n {
			end = n
			sub = x.SubIndex + 1
			break
		}
	}

	return end, sub
}
