package logic

import (
	"context"
	"easy-chat/apps/im/rpc/imclient"
	"github.com/jinzhu/copier"
	"log"

	"easy-chat/apps/im/api/internal/svc"
	"easy-chat/apps/im/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 根据用户获取聊天记录
func NewGetChatLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatLogLogic {
	return &GetChatLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChatLogLogic) GetChatLog(req *types.ChatLogReq) (resp *types.ChatLogResp, err error) {
	// 打印传入参数
	log.Printf("API 请求参数: %+v", req)

	// 转换纳秒到毫秒
	req.StartSendTime /= 1e6
	req.EndSendTime /= 1e6

	// 调用 RPC 获取数据
	data, err := l.svcCtx.GetChatLog(l.ctx, &imclient.GetChatLogReq{
		ConversationId: req.ConversationId,
		StartSendTime:  req.StartSendTime,
		EndSendTime:    req.EndSendTime,
		Count:          req.Count,
	})
	if err != nil {
		log.Printf("RPC 调用失败: %v", err)
		return nil, err
	}

	log.Printf("RPC 返回数据: %+v", data)

	var res types.ChatLogResp
	copier.Copy(&res, data)

	return &res, err
}
