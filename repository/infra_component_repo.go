package repository

import (
	"context"

	"sllpklls/admin-service/model"
)

type InfraComponentRepo interface {
	GetAllInfraComponents(ctx context.Context) ([]model.InfraComponent, error)
}
