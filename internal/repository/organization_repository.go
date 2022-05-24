package repository

import (
	"context"

	"github.com/calvinbenhardi/go-sqlx/internal/model"
	"github.com/google/uuid"
)

type OrganizationRepository interface {
	List(ctx context.Context, arg model.ListOrganizationParams) ([]model.Organization, error)
	Save(ctx context.Context, arg model.CreateOrganizationParams) (model.Organization, error)
	Get(ctx context.Context, id uuid.UUID) (model.Organization, error)
	Update(ctx context.Context, arg model.UpdateOrganizationParams) (model.Organization, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
