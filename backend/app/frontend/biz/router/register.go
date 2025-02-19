// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	about "github.com/tiktokmall/backend/app/frontend/biz/router/about"
	agent "github.com/tiktokmall/backend/app/frontend/biz/router/agent"
	auth "github.com/tiktokmall/backend/app/frontend/biz/router/auth"
	cart "github.com/tiktokmall/backend/app/frontend/biz/router/cart"
	category "github.com/tiktokmall/backend/app/frontend/biz/router/category"
	checkout "github.com/tiktokmall/backend/app/frontend/biz/router/checkout"
	home "github.com/tiktokmall/backend/app/frontend/biz/router/home"
	merchant "github.com/tiktokmall/backend/app/frontend/biz/router/merchant"
	order "github.com/tiktokmall/backend/app/frontend/biz/router/order"
	oss "github.com/tiktokmall/backend/app/frontend/biz/router/oss"
	product "github.com/tiktokmall/backend/app/frontend/biz/router/product"
	user "github.com/tiktokmall/backend/app/frontend/biz/router/user"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	agent.Register(r)

	merchant.Register(r)

	oss.Register(r)
	user.Register(r)

	order.Register(r)

	checkout.Register(r)

	about.Register(r)

	cart.Register(r)

	category.Register(r)

	product.Register(r)

	auth.Register(r)

	home.Register(r)
}
