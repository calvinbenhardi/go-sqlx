package repository

import (
	"context"

	"github.com/calvinbenhardi/go-sqlx/internal/model"
)

type OrganizationRepository interface {
	Save(ctx context.Context, arg model.CreateOrganizationParams) (model.Organization, error)
}
