package service

import (
	"context"
	"fmt"

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

	// TODO：这里可以使用缓存优化，防止恶意通过此接口，导致数据库压力过大
	u1, err := model.GetById(mysql.DB, s.ctx, (req.UserId))
	if err != nil {
		return nil, err
	}
	if req.Username == "" {
		return nil, fmt.Errorf("username 不能为空")
	}
	if req.Username == u1.Username && req.Avator == u1.Avator {
		return
	}

	u1.Username = req.Username
	u1.Avator = req.Avator

	if err = model.UpdateUser(mysql.DB, s.ctx, u1); err != nil {
		return nil, err
	}
	return &user.UpdateUserResp{UserId: req.UserId}, nil
}

// Crypt Encrypt the password using crypto/bcrypt
func crypt(password string) (string, error) {
	// Generate "cost" factor for the bcrypt algorithm
	// Hash password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}
