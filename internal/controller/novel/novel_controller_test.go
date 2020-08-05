package novel

import (
	"bm-novel/internal/config"
	"bm-novel/internal/domain/novel"
	rp "bm-novel/internal/infrastructure/persistence/novel"
	"bm-novel/internal/infrastructure/postgres"
	"context"
	"testing"

	uuid "github.com/satori/go.uuid"
)

func init() {
	config.LoadConfigForTest()
	postgres.InitDB()
}

func TestGetNovels(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())
	repo := rp.New()

	n := &novel.Novel{NovelTitle: "abc2", NovelID: uuid.NewV4(), ChiefEditorID: uuid.Nil}
	service := novel.Service{Repo: repo}

	err := service.Create(ctx, n)

	t.Log(err)
}
