package service

import (
	"context"

	"github.com/tiktokmall/backend/app/user/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/user/biz/dal/redis"
	"github.com/tiktokmall/backend/app/user/biz/model"
	user "github.com/tiktokmall/backend/rpc_gen/kitex_gen/user"

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

	// 更新使用缓存的查询逻辑
	dao := model.NewUserDAO(s.ctx, mysql.DB, redis.RedisClient)
	userRow, err := dao.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userRow.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	return &user.LoginResp{
		UserId:   int64(userRow.ID),
		Username: userRow.Username,
		Email:    userRow.Email,
	}, nil
}
