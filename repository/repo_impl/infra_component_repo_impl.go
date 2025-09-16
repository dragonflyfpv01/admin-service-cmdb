package repo_impl

import (
	"context"

	"sllpklls/admin-service/db"
	"sllpklls/admin-service/model"
	"sllpklls/admin-service/repository"

	"github.com/labstack/gommon/log"
)

type InfraComponentRepoImpl struct {
	sql *db.Sql
}

func NewInfraComponentRepo(sql *db.Sql) repository.InfraComponentRepo {
	return &InfraComponentRepoImpl{
		sql: sql,
	}
}

func (i *InfraComponentRepoImpl) GetAllInfraComponents(ctx context.Context) ([]model.InfraComponent, error) {
	var components []model.InfraComponent
	statement := `
		SELECT id, hostname, dns, description, public_internet, class, ipaddress, 
		       subnet, site, it_component_type, request_type, appid, vlan, 
		       app_name, app_owner, level, ci_owners, im_cm, status, 
		       created_at, create_by 
		FROM infra_components 
		ORDER BY id DESC
	`
	err := i.sql.Db.SelectContext(ctx, &components, statement)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return components, nil
}

func (i *InfraComponentRepoImpl) GetInfraComponentsPaginated(ctx context.Context, pagination model.PaginationRequest) ([]model.InfraComponent, int64, error) {
	var components []model.InfraComponent
	var totalCount int64

	// Đếm tổng số bản ghi
	countStatement := `SELECT COUNT(*) FROM infra_components`
	err := i.sql.Db.GetContext(ctx, &totalCount, countStatement)
	if err != nil {
		log.Error("Error counting infra components:", err.Error())
		return nil, 0, err
	}

	// Lấy dữ liệu với phân trang
	statement := `
		SELECT id, hostname, dns, description, public_internet, class, ipaddress, 
		       subnet, site, it_component_type, request_type, appid, vlan, 
		       app_name, app_owner, level, ci_owners, im_cm, status, 
		       created_at, create_by 
		FROM infra_components 
		ORDER BY id DESC
		LIMIT $1 OFFSET $2
	`

	err = i.sql.Db.SelectContext(ctx, &components, statement, pagination.Limit, pagination.GetOffset())
	if err != nil {
		log.Error("Error getting paginated infra components:", err.Error())
		return nil, 0, err
	}

	return components, totalCount, nil
}
