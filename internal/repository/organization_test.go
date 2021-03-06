package repository

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/calvinbenhardi/go-sqlx/config"
	"github.com/calvinbenhardi/go-sqlx/internal/model"
	"github.com/stretchr/testify/require"
)

var repo = NewProductRepository(config.NewPostgres())

func CreateOrganization(t *testing.T) model.Organization {
	gofakeit.Seed(0)
	arg := model.CreateOrganizationParams{
		Name: gofakeit.Company(),
	}

	organization, err := repo.Save(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, organization.ID)
	require.Equal(t, arg.Name, organization.Name)
	require.NotEmpty(t, organization.CreatedAt)
	require.NotEmpty(t, organization.UpdatedAt)

	return organization
}

func TestCreateOrganization(t *testing.T) {
	CreateOrganization(t)
}

func TestListOrganization(t *testing.T) {
	for i := 0; i <= 10; i++ {
		CreateOrganization(t)
	}

	arg := model.ListOrganizationParams{
		Limit:  5,
		Offset: 0,
	}

	organizations, err := repo.List(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, organizations)

	require.Equal(t, arg.Limit, len(organizations))
}

func TestGetOrganization(t *testing.T) {
	organization1 := CreateOrganization(t)
	organization2, err := repo.Get(context.Background(), organization1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, organization2)

	require.Equal(t, organization1.ID, organization2.ID)
	require.Equal(t, organization1.Name, organization2.Name)
	require.NotEmpty(t, organization2.CreatedAt)
	require.NotEmpty(t, organization2.UpdatedAt)

}

func TestUpdateOrganization(t *testing.T) {
	organization1 := CreateOrganization(t)

	arg := model.UpdateOrganizationParams{
		ID:   organization1.ID,
		Name: gofakeit.Name(),
	}

	organization2, err := repo.Update(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, organization2)

	require.Equal(t, arg.Name, organization2.Name)
}

func TestDeleteOrganization(t *testing.T) {
	organization1 := CreateOrganization(t)

	err := repo.Delete(context.Background(), organization1.ID)
	require.NoError(t, err)

	organization2, err := repo.Get(context.Background(), organization1.ID)
	require.Error(t, err)
	require.Empty(t, organization2)
}
