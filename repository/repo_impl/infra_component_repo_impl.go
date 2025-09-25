package repo_impl

import (
	"context"
	"fmt"

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

func (i *InfraComponentRepoImpl) GetInfraComponentsByStatus(ctx context.Context, status string) ([]model.InfraComponent, error) {
	var components []model.InfraComponent
	statement := `
		SELECT id, hostname, dns, description, public_internet, class, ipaddress, 
		       subnet, site, it_component_type, request_type, appid, vlan, 
		       app_name, app_owner, level, ci_owners, im_cm, status, 
		       created_at, create_by 
		FROM infra_components 
		WHERE status = $1
		ORDER BY id DESC
	`
	err := i.sql.Db.SelectContext(ctx, &components, statement, status)
	if err != nil {
		log.Error("Error getting infra components by status:", err.Error())
		return nil, err
	}
	return components, nil
}

func (i *InfraComponentRepoImpl) UpdateInfraComponentStatus(ctx context.Context, id int, hostname string, newStatus string) error {
	// statement := `
	// 	UPDATE infra_components
	// 	SET status = $1
	// 	WHERE id = $2 AND hostname = $3 AND status = 'Đang chờ'
	// `

	statement := `
		UPDATE infra_components 
		SET status = $1 
		WHERE id = $2 AND hostname = $3
	`

	result, err := i.sql.Db.ExecContext(ctx, statement, newStatus, id, hostname)
	if err != nil {
		log.Error("Error updating infra component status:", err.Error())
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error("Error getting rows affected:", err.Error())
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no infra component found with ID %d and hostname '%s' or status is not 'Đang chờ'", id, hostname)
	}

	return nil
}

func (i *InfraComponentRepoImpl) UpdateInfraComponent(ctx context.Context, id int, component model.InfraComponent) error {
	statement := `
		UPDATE infra_components 
		SET hostname = $1, dns = $2, description = $3, public_internet = $4, 
		    class = $5, ipaddress = $6, subnet = $7, site = $8, 
		    it_component_type = $9, request_type = $10, appid = $11, vlan = $12, 
		    app_name = $13, app_owner = $14, level = $15, ci_owners = $16, 
		    im_cm = $17, created_at = $18
		WHERE id = $19
	`

	result, err := i.sql.Db.ExecContext(ctx, statement,
		component.Hostname, component.DNS, component.Description, component.PublicInternet,
		component.Class, component.IPAddress, component.Subnet, component.Site,
		component.ITComponentType, component.RequestType, component.AppID, component.VLAN,
		component.AppName, component.AppOwner, component.Level, component.CIOwners,
		component.IMCM, component.CreatedAt, id)

	if err != nil {
		log.Error("Error updating infra component:", err.Error())
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error("Error getting rows affected:", err.Error())
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no infra component found with ID %d", id)
	}

	return nil
}

func (i *InfraComponentRepoImpl) CreateInfraComponent(ctx context.Context, component model.InfraComponent) (*model.InfraComponent, error) {
	statement := `
		INSERT INTO infra_components (
			hostname, dns, description, public_internet, class, ipaddress, 
			subnet, site, it_component_type, request_type, appid, vlan, 
			app_name, app_owner, level, ci_owners, im_cm, status, 
			created_at, create_by
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, 
			$11, $12, $13, $14, $15, $16, $17, $18, $19, $20
		) RETURNING id, created_at`

	var returnedID int
	var returnedCreatedAt string

	err := i.sql.Db.QueryRowContext(ctx, statement,
		component.Hostname, component.DNS, component.Description, component.PublicInternet,
		component.Class, component.IPAddress, component.Subnet, component.Site,
		component.ITComponentType, component.RequestType, component.AppID, component.VLAN,
		component.AppName, component.AppOwner, component.Level, component.CIOwners,
		component.IMCM, component.Status, component.CreatedAt, component.CreateBy,
	).Scan(&returnedID, &returnedCreatedAt)

	if err != nil {
		log.Error("Error creating infra component:", err.Error())
		return nil, err
	}

	// Set returned values to the component
	component.ID = returnedID
	component.CreatedAt = returnedCreatedAt

	return &component, nil
}
