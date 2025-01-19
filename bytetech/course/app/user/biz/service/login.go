package service

import (
	"bytetech/course/app/user/biz/dal/mysql"
	"bytetech/course/app/user/biz/model"
	user "bytetech/course/rpc_gen/kitex_gen/user"
	"context"

	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.

	userRow, err := model.GetByEmail(mysql.DB, s.ctx, req.Email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userRow.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	return &user.LoginResp{
		UserId: int32(userRow.ID),
	}, nil
}
