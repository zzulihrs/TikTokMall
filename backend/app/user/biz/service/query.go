package service

import (
	"context"

	"github.com/tiktokmall/backend/app/user/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/user/biz/model"
	user "github.com/tiktokmall/backend/rpc_gen/kitex_gen/user"
)

type QueryService struct {
	ctx context.Context
} // NewQueryService new QueryService
func NewQueryService(ctx context.Context) *QueryService {
	return &QueryService{ctx: ctx}
}

// Run create note info
func (s *QueryService) Run(req *user.QueryUserReq) (resp *user.QueryUserResp, err error) {
	// Finish your business logic.
	u := &model.User{}
	if u, err = model.QueryUser(mysql.DB, s.ctx, req.UserId); err != nil {
		return nil, err
	}
	resp = &user.QueryUserResp{
		User: &user.User{
			UserId:   int64(u.ID),
			Email:    u.Email,
			Username: u.Username,
			Avator:   u.Avator,
		},
	}
	return
}
