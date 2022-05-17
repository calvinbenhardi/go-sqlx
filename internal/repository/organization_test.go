package repository

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/calvinbenhardi/go-sqlx/config"
	"github.com/calvinbenhardi/go-sqlx/internal/model"
	"github.com/stretchr/testify/require"
)

func TestCreateOrganization(t *testing.T) {
	gofakeit.Seed(0)
	arg := model.CreateOrganizationParams{
		Name: gofakeit.Company(),
	}

	repo := NewProductRepository(config.NewPostgres())
	organization, err := repo.Save(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, organization.ID)
	require.Equal(t, arg.Name, organization.Name)
	require.NotEmpty(t, organization.CreatedAt)
	require.NotEmpty(t, organization.UpdatedAt)
}
