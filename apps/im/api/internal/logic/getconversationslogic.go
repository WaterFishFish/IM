package logic

import (
	"context"
	"easy-chat/apps/im/rpc/imclient"
	"easy-chat/pkg/ctxdata"
	"github.com/jinzhu/copier"
	"log"

	"easy-chat/apps/im/api/internal/svc"
	"easy-chat/apps/im/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConversationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取会话
func NewGetConversationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConversationsLogic {
	return &GetConversationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConversationsLogic) GetConversations(req *types.GetConversationsReq) (resp *types.GetConversationsResp, err error) {
	// todo: add your logic here and delete this line
	uid := ctxdata.GetUid(l.ctx)
	log.Printf("UID: %s", uid)
	data, err := l.svcCtx.GetConversations(l.ctx, &imclient.GetConversationsReq{
		UserId: uid,
	})
	log.Printf("GetConversations response data: %+v", data)

	if err != nil {
		return nil, err
	}

	var res types.GetConversationsResp
	copier.Copy(&res, &data)

	return &res, err

}
