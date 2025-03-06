package service

import (
	"context"
	"errors"
	"io"
	"log"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/sse"
	"github.com/tiktokmall/backend/app/agent/cmd/einoagent/agent"
	common "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/common"
	frontendUtils "github.com/tiktokmall/backend/app/frontend/utils"
)

type HandleChatService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHandleChatService(Context context.Context, RequestContext *app.RequestContext) *HandleChatService {
	return &HandleChatService{RequestContext: RequestContext, Context: Context}
}

func (h *HandleChatService) Run(req *common.Empty) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	id := strconv.Itoa(int(frontendUtils.GetUserIdFromCtx(h.Context)))
	message := h.RequestContext.Query("message")
	if id == "" || message == "" {
		h.RequestContext.JSON(consts.StatusBadRequest, map[string]string{
			"status": "error",
			"error":  "missing id or message parameter",
		})
		return
	}
	log.Printf("[Chat] Starting chat with ID: %s, Message: %s\n", id, message)

	// core function
	sr, err := agent.RunAgent(h.Context, id, message)
	if err != nil {
		log.Printf("[Chat] Error running agent: %v\n", err)
		h.RequestContext.JSON(consts.StatusInternalServerError, map[string]string{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	//  如何获取这个sse流？
	s := sse.NewStream(h.RequestContext)
	defer func() {
		sr.Close()
		h.RequestContext.Flush()

		log.Printf("[Chat] Finished chat with ID: %s\n", id)
	}()

outer:
	for {
		select {
		case <-h.Context.Done():
			log.Printf("[Chat] Context done for chat ID: %s\n", id)
			return
		default:
			msg, err := sr.Recv()
			if errors.Is(err, io.EOF) {
				log.Printf("[Chat] EOF received for chat ID: %s\n", id)
				break outer
			}
			if err != nil {
				log.Printf("[Chat] Error receiving message: %v\n", err)
				break outer
			}

			err = s.Publish(&sse.Event{
				Data: []byte(msg.Content),
			})
			if err != nil {
				log.Printf("[Chat] Error publishing message: %v\n", err)
				break outer
			}
		}
	}
	return
}
