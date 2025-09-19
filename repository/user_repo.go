package repository

import (
	"context"

	"sllpklls/admin-service/model"
	"sllpklls/admin-service/model/req"
)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
	CheckLogin(context context.Context, loginReq req.ReqLogin) (model.User, error)
	GetAllUsers(context context.Context) ([]model.UserList, error)
}
