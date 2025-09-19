package repo_impl

import (
	"context"
	"database/sql"
	"time"

	"sllpklls/admin-service/db"
	"sllpklls/admin-service/errors"
	"sllpklls/admin-service/model"
	"sllpklls/admin-service/model/req"
	"sllpklls/admin-service/repository"

	"github.com/labstack/gommon/log"
	"github.com/lib/pq"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) repository.UserRepo {
	return &UserRepoImpl{
		sql: sql,
	}
}
func (u UserRepoImpl) SaveUser(context context.Context, user model.User) (model.User, error) {
	statement := `
		INSERT INTO users_admin(user_id, email, password, role, full_name, created_at, updated_at)
		VALUES(:user_id, :email, :password, :role, :full_name, :created_at, :updated_at)
	`
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	if _, err := u.sql.Db.NamedExecContext(context, statement, user); err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, errors.UserConflict
			}
		}
		return user, err
	}
	return user, nil
}
func (u *UserRepoImpl) CheckLogin(context context.Context, loginReq req.ReqLogin) (model.User, error) {
	var user = model.User{}
	statement := `SELECT * FROM users_admin WHERE email =$1`
	err := u.sql.Db.GetContext(context, &user, statement, loginReq.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.UserNotFound
		}
		log.Error(err.Error())
		return user, err
	}
	return user, nil
}

func (u *UserRepoImpl) GetAllUsers(context context.Context) ([]model.UserList, error) {
	var users []model.UserList
	statement := `SELECT id, username, user_id, role FROM users ORDER BY id`

	err := u.sql.Db.SelectContext(context, &users, statement)
	if err != nil {
		log.Error("Error fetching users: ", err.Error())
		return nil, err
	}

	return users, nil
}
