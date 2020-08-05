package novel

import (
	"bm-novel/internal/domain/novel"
	"bm-novel/internal/http/web"
	rp "bm-novel/internal/infrastructure/persistence/novel"
	"net/http"
	"sync"

	"github.com/joyparty/httpkit"
)

var (
	ns   *novel.Service
	once sync.Once
)

func service() *novel.Service {
	once.Do(func() {
		ns = &novel.Service{Repo: rp.New()}
	})

	return ns
}

// PostNovels 添加小说
func PostNovels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	params := struct {
		Title string `json:"title"  valid:"required"`
	}{}

	httpkit.MustScanJSON(&params, r.Body)

	err := service().Create(r.Context(), &novel.Novel{NovelTitle: params.Title})
	if err == nil {
		w.WriteHeader(201)
		return
	}

	web.WriteHTTPStats(w, err)

}
