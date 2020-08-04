package novel

import (
	"context"

	uuid "github.com/satori/go.uuid"
)

type Service struct {
}

func (s Service) Create(ctx context.Context, novel Novel) (*Novel, error) {
	panic("implement me")
}

func (s Service) Delete(ctx context.Context, novelID uuid.UUID) error {
	panic("implement me")
}

func (s Service) AssignResponsibleEditor(ctx context.Context, novelID uuid.UUID, editorID uuid.UUID) error {
	panic("implement me")
}

func (s Service) SetFormat(ctx context.Context, novelID uuid.UUID, format Settings) error {
	panic("implement me")
}

func (s Service) UploadDraft(ctx context.Context, novelID uuid.UUID, draft string) error {
	panic("implement me")
}
