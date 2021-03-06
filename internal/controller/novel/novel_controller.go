package novel

import (
	"bm-novel/internal/domain/novel"
	"bm-novel/internal/http/auth"
	"bm-novel/internal/http/web"
	rc "bm-novel/internal/infrastructure/persistence/chapter"
	rnc "bm-novel/internal/infrastructure/persistence/couter"
	rn "bm-novel/internal/infrastructure/persistence/novel"
	rp "bm-novel/internal/infrastructure/persistence/paragraph"
	"net/http"
	"sync"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"

	"github.com/joyparty/httpkit"
)

var (
	ns   *novel.Service
	once sync.Once
)

func service() *novel.Service {
	once.Do(func() {
		ns = &novel.Service{Repo: rn.New(), ChapterRepo: rc.New(), ParagraphRepo: rp.New(), Counter: rnc.New()}
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

	userID, err := auth.GetVisitorUserIDFromJWT(r)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	err = service().Create(r.Context(), &novel.Novel{NovelTitle: params.Title, ChiefEditorID: userID})
	if err == nil {
		w.WriteHeader(201)
		return
	}

	web.WriteHTTPStats(w, err)
}

// PutResponsibleEditor 指派责编
func PutResponsibleEditor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	params := struct {
		EditorID string `json:"editor_id"  valid:"required"`
	}{}

	httpkit.MustScanJSON(&params, r.Body)
	novelID, err := getNovelIDForURLParam(r)
	if err != nil {
		web.WriteHTTPStats(w, err)
		return
	}

	editorID, err := uuid.FromString(params.EditorID)
	if err != nil {
		web.WriteHTTPStats(w, web.ErrNotFound)
		return
	}

	err = service().AssignResponsibleEditor(r.Context(), novelID, editorID)
	if err != nil {
		web.WriteHTTPStats(w, err)
		return
	}
}

// PostDraft 上传原文
func PostDraft(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	novelID, err := getNovelIDForURLParam(r)
	if err != nil {
		web.WriteHTTPStats(w, err)
		return
	}

	file, _, err := r.FormFile("draft")
	if err != nil {
		web.WriteHTTPStats(w, err)
		return
	}
	defer file.Close()

	err = service().UploadDraft(r.Context(), novelID, file)
	if err != nil {
		web.WriteHTTPStats(w, err)
		return
	}
}

// getNovelIDForURLParam NovelID
func getNovelIDForURLParam(r *http.Request) (novelID uuid.UUID, err error) {
	id := chi.URLParam(r, "novel_id")

	if id == "" {
		return novelID, web.ErrNotFound
	}

	novelID, err = uuid.FromString(id)
	if err != nil {
		return novelID, web.ErrServerError
	}

	return novelID, nil
}
