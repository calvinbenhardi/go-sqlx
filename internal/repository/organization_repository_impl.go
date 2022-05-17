package repository

import (
	"context"

	"github.com/calvinbenhardi/go-sqlx/internal/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func NewProductRepository(db *sqlx.DB) *organizationRepositoryImpl {
	return &organizationRepositoryImpl{db: db}
}

type organizationRepositoryImpl struct {
	db *sqlx.DB
}

func (r *organizationRepositoryImpl) Save(ctx context.Context, arg model.CreateOrganizationParams) (model.Organization, error) {
	var organization model.Organization

	row := r.db.QueryRowContext(ctx, `INSERT INTO organizations (name) VALUES ($1) RETURNING *`, arg.Name)
	err := row.Scan(
		&organization.ID,
		&organization.Name,
		&organization.CreatedAt,
		&organization.UpdatedAt,
	)

	if err != nil {
		return organization, err
	}

	return organization, nil
}

func (r *organizationRepositoryImpl) Get(ctx context.Context, id uuid.UUID) (model.Organization, error) {
	var organization model.Organization

	row := r.db.QueryRowContext(ctx, `SELECT id, name, created_at, updated_at FROM organizations WHERE id=$1`, id)
	err := row.Scan(
		&organization.ID,
		&organization.Name,
		&organization.CreatedAt,
		&organization.UpdatedAt,
	)

	if err != nil {
		return organization, err
	}

	return organization, nil
}

func (r *organizationRepositoryImpl) Update(ctx context.Context, arg model.UpdateOrganizationParams) (model.Organization, error) {
	var organization model.Organization

	row := r.db.QueryRowContext(ctx, `UPDATE organizations SET name = $1 WHERE id = $2 RETURNING *`, arg.Name, arg.ID)
	err := row.Scan(
		&organization.ID,
		&organization.Name,
		&organization.CreatedAt,
		&organization.UpdatedAt,
	)

	if err != nil {
		return organization, err
	}

	return organization, nil
}
