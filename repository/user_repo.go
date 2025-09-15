package repository

import (
	"context"

	"github.com/sllpklls/template-backend-go/model"
	"github.com/sllpklls/template-backend-go/model/req"
)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
	CheckLogin(context context.Context, loginReq req.ReqSignIn) (model.User, error)
}
