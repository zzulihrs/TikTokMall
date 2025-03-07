// Code generated by hertz generator. DO NOT EDIT.

package agent

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	agent "github.com/tiktokmall/backend/app/frontend/biz/handler/agent"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_agent := root.Group("/agent", _agentMw()...)
		{
			_api := _agent.Group("/api", _apiMw()...)
			_api.POST("/chat", append(_handlechatMw(), agent.HandleChat)...)
		}
	}
}
