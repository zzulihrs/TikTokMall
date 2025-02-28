package service

import (
	"context"

	"github.com/tiktokmall/backend/app/user/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/user/biz/dal/redis"
	"github.com/tiktokmall/backend/app/user/biz/model"

	user "github.com/tiktokmall/backend/rpc_gen/kitex_gen/user"
)

type DeleteService struct {
	ctx context.Context
} // NewDeleteService new DeleteService
func NewDeleteService(ctx context.Context) *DeleteService {
	return &DeleteService{ctx: ctx}
}

// Run create note info
func (s *DeleteService) Run(req *user.DeleteUserReq) (resp *user.DeleteUserResp, err error) {
	// Finish your business logic.
	if err = model.DeleteUser(mysql.DB, s.ctx, req.UserId); err != nil {
		return nil, err
	}
	err = model.NewUserDAO(s.ctx, mysql.DB, redis.RedisClient).DeleteByID(int(req.UserId))

	return
}
