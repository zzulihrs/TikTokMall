package service

import (
	"context"

	"github.com/tiktokmall/backend/app/user/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/user/biz/model"
	user "github.com/tiktokmall/backend/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type UpdateService struct {
	ctx context.Context
} // NewUpdateService new UpdateService
func NewUpdateService(ctx context.Context) *UpdateService {
	return &UpdateService{ctx: ctx}
}

// Run create note info
func (s *UpdateService) Run(req *user.UpdateUserReq) (resp *user.UpdateUserResp, err error) {
	// Finish your business logic.
	_, err = model.GetById(mysql.DB, s.ctx, (req.UserId))
	if err != nil {
		return nil, err
	}
	u := &model.User{}
	u.ID = uint(req.UserId)
	u.Email = req.Email
	u.PasswordHashed, err = Crypt(req.Password)
	u.Gender = int32(req.Gender)
	u.Nickname = req.Nickname
	u.Signature = req.Signature
	if err != nil {
		return nil, err
	}
	if err = model.UpdateUser(mysql.DB, s.ctx, u); err != nil {
		return nil, err
	}
	return &user.UpdateUserResp{UserId: req.UserId}, nil
}

// Crypt Encrypt the password using crypto/bcrypt
func Crypt(password string) (string, error) {
	// Generate "cost" factor for the bcrypt algorithm
	// Hash password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}
