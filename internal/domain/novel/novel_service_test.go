package novel

import (
	"container/list"
	"testing"
)

func TestService_Create(t *testing.T) {
	//ctx, _ := context.WithCancel(context.Background())
	//repo := novel.New()
	//
	//novel := &Novel{NovelTitle: "abc"}
	//service := Service{Repo: repo}
	//
	//err := service.Create(ctx, novel)
	//
	//t.Log(err)
	l := list.New()
	e4 := l.PushBack("a4")
	e1 := l.PushFront("a1")
	l.InsertBefore("a3", e4)
	l.InsertAfter("a2", e1)

	//sort.Sort(l)
}
