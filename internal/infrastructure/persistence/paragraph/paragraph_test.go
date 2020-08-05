package paragraph

import (
	"bm-novel/internal/config"
	"bm-novel/internal/domain/novel/paragraph"
	"bm-novel/internal/infrastructure/postgres"
	"container/list"
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

type v struct {
	index int
	code  string
	prev  string
	next  string
}

type vl []*v

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

	vlist := &vl{
		{7, "0d6fbec4-a0c5-468f-9343-5ad945ae2c5f", "a8b797f7-a4c6-41ec-b13f-ea1718ae4a17", "283c0927-629a-4c96-81ed-8c9a912789fc"},
		{26, "0de0cf98-0e42-40b3-9dae-07c6f52f629e", "48135fe4-6237-40f3-b322-7ca05801c873", "34123c49-f7e5-49c8-a12e-76a8362838ec"},
		{8, "283c0927-629a-4c96-81ed-8c9a912789fc", "0d6fbec4-a0c5-468f-9343-5ad945ae2c5f", "492242ae-5ebd-4ca9-a33f-a5d30897423e"},
		{27, "34123c49-f7e5-49c8-a12e-76a8362838ec", "0de0cf98-0e42-40b3-9dae-07c6f52f629e", "36c99a65-ed90-40fd-a797-c4e61bac9c1d"},
		{28, "36c99a65-ed90-40fd-a797-c4e61bac9c1d", "34123c49-f7e5-49c8-a12e-76a8362838ec", "42d780fc-fec3-447f-813b-a539c6362d1c"},
		{24, "36e8c568-668d-45e3-b622-8eb9d5ff5e17", "ffdfae25-41f2-4920-b18b-6fcce648cdda", "48135fe4-6237-40f3-b322-7ca05801c873"},
		{12, "3e8c8f37-ab42-4677-98f4-d1087183b09e", "d29fb940-2836-4bff-a585-e024303ade71", "5bb7ab95-f111-43ef-a32e-5169291917a4"},
		{2, "42a2cd05-434f-497d-bab3-b7564e8cf0f9", "4ba0be72-3681-4a85-95fc-cbb2c2f20fe3", "9195e5a2-942d-44f1-94f9-ce0021c6b11a"},
		{29, "42d780fc-fec3-447f-813b-a539c6362d1c", "36c99a65-ed90-40fd-a797-c4e61bac9c1d", "fd35ce7d-904c-4263-a29c-87c9214917af"},
		{17, "46eb4192-94e2-4811-a020-8b994406fe28", "588c5f8c-1fa8-418e-b447-3a1ec01a0c72", "9d123c5e-b719-4848-83a3-3ad2f4b4b212"},
		{25, "48135fe4-6237-40f3-b322-7ca05801c873", "36e8c568-668d-45e3-b622-8eb9d5ff5e17", "0de0cf98-0e42-40b3-9dae-07c6f52f629e"},
		{9, "492242ae-5ebd-4ca9-a33f-a5d30897423e", "283c0927-629a-4c96-81ed-8c9a912789fc", "f3034c67-7d02-468a-a119-e60b5b69c434"},
		{1, "4ba0be72-3681-4a85-95fc-cbb2c2f20fe3", "", "42a2cd05-434f-497d-bab3-b7564e8cf0f9"},
		{21, "5041801c-f1fa-4255-8de3-3bf4be216c99", "d36e5f27-3e8e-4efd-b89d-e038ece1b437", "c49d5de7-7056-4b8a-b8bb-fe919c2ec97a"},
		{16, "588c5f8c-1fa8-418e-b447-3a1ec01a0c72", "f318421f-09c0-4b2f-9032-6e99c21de3c1", "46eb4192-94e2-4811-a020-8b994406fe28"},
		{13, "5bb7ab95-f111-43ef-a32e-5169291917a4", "3e8c8f37-ab42-4677-98f4-d1087183b09e", "b80526c1-c94d-4728-8972-6584bc3c0b15"},
		{4, "8b7de0a4-8063-4e34-bbff-d2526b016f59", "9195e5a2-942d-44f1-94f9-ce0021c6b11a", "e6b1fcea-b424-40c1-a27e-8e218e03132b"},
		{3, "9195e5a2-942d-44f1-94f9-ce0021c6b11a", "42a2cd05-434f-497d-bab3-b7564e8cf0f9", "8b7de0a4-8063-4e34-bbff-d2526b016f59"},
		{18, "9d123c5e-b719-4848-83a3-3ad2f4b4b212", "46eb4192-94e2-4811-a020-8b994406fe28", "dcaf3bd6-6beb-46e6-ba73-5c0a3654a944"},
		{6, "a8b797f7-a4c6-41ec-b13f-ea1718ae4a17", "e6b1fcea-b424-40c1-a27e-8e218e03132b", "0d6fbec4-a0c5-468f-9343-5ad945ae2c5f"},
		{14, "b80526c1-c94d-4728-8972-6584bc3c0b15", "5bb7ab95-f111-43ef-a32e-5169291917a4", "f318421f-09c0-4b2f-9032-6e99c21de3c1"},
		{22, "c49d5de7-7056-4b8a-b8bb-fe919c2ec97a", "5041801c-f1fa-4255-8de3-3bf4be216c99", "ffdfae25-41f2-4920-b18b-6fcce648cdda"},
		{11, "d29fb940-2836-4bff-a585-e024303ade71", "f3034c67-7d02-468a-a119-e60b5b69c434", "3e8c8f37-ab42-4677-98f4-d1087183b09e"},
		{20, "d36e5f27-3e8e-4efd-b89d-e038ece1b437", "dcaf3bd6-6beb-46e6-ba73-5c0a3654a944", "5041801c-f1fa-4255-8de3-3bf4be216c99"},
		{19, "dcaf3bd6-6beb-46e6-ba73-5c0a3654a944", "9d123c5e-b719-4848-83a3-3ad2f4b4b212", "d36e5f27-3e8e-4efd-b89d-e038ece1b437"},
		{5, "e6b1fcea-b424-40c1-a27e-8e218e03132b", "8b7de0a4-8063-4e34-bbff-d2526b016f59", "a8b797f7-a4c6-41ec-b13f-ea1718ae4a17"},
		{10, "f3034c67-7d02-468a-a119-e60b5b69c434", "492242ae-5ebd-4ca9-a33f-a5d30897423e", "d29fb940-2836-4bff-a585-e024303ade71"},
		{15, "f318421f-09c0-4b2f-9032-6e99c21de3c1", "b80526c1-c94d-4728-8972-6584bc3c0b15", "588c5f8c-1fa8-418e-b447-3a1ec01a0c72"},
		{30, "fd35ce7d-904c-4263-a29c-87c9214917af", "42d780fc-fec3-447f-813b-a539c6362d1c", ""},
		{23, "ffdfae25-41f2-4920-b18b-6fcce648cdda", "c49d5de7-7056-4b8a-b8bb-fe919c2ec97a", "36e8c568-668d-45e3-b622-8eb9d5ff5e17"},
	}

	data := make(map[string]*v)

	for _, w := range *vlist {
		data[w.code] = w
	}

	l := list.New()

	for _, w := range data {

		//fmt.Println(s.index)
		if l.Back() == nil {
			l.PushBack(w)
		} else {
			if s1, ok1 := l.Back().Value.(*v); ok1 {
				if s1.next != "" {
					l.PushBack(data[s1.next])
				} else {
					if s2, ok2 := l.Front().Value.(*v); ok2 {
						l.PushFront(data[s2.prev])
					}
				}
			}

		}

	}

	for item := l.Front(); nil != item; item = item.Next() {
		if s1, ok1 := item.Value.(*v); ok1 {
			fmt.Printf("%d, ", s1.index)
		}
	}

}

type v1 struct {
	index int
	code  string
}

func Test1(t *testing.T) {
	vl := []v1{
		{20, "a"},
		{23, "b"},
		{22, "e"},
		{23, "22"},
		{37, "jjj"},
		{38, "ss"},
		{39, "sd"},
		{40, "jh"},
		{41, "34"},
		{50, "89"}}

	// 需要插入值
	w := "iii"
	// 插入位置
	prev := 20
	// 取大于20的5条,排序取
	vl2 := vl[1:6]

	end := 0
	for i, x := range vl2 {
		n := prev + i + 1
		if x.index > n {
			end = n
			break
		}
	}

	fmt.Println(w, end, vl2)
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
