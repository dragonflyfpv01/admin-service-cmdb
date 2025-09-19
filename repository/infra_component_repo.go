package repository

import (
	"context"

	"sllpklls/admin-service/model"
)

type InfraComponentRepo interface {
	GetAllInfraComponents(ctx context.Context) ([]model.InfraComponent, error)
	GetInfraComponentsPaginated(ctx context.Context, pagination model.PaginationRequest) ([]model.InfraComponent, int64, error)
	GetInfraComponentsByStatus(ctx context.Context, status string) ([]model.InfraComponent, error)
	UpdateInfraComponentStatus(ctx context.Context, id int, hostname string, newStatus string) error
	UpdateInfraComponent(ctx context.Context, id int, component model.InfraComponent) error
}
