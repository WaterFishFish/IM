package logic

import (
	"context"
	"easy-chat/pkg/xerr"
	"github.com/pkg/errors"
	"log"

	"easy-chat/apps/im/rpc/im"
	"easy-chat/apps/im/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetChatLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatLogLogic {
	return &GetChatLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取会话记录
func (l *GetChatLogLogic) GetChatLog(in *im.GetChatLogReq) (*im.GetChatLogResp, error) {
	// todo: add your logic here and delete this line

	// 根据id
	if in.MsgId != "" {
		chatLog, err := l.svcCtx.ChatLogModel.FindOne(l.ctx, in.MsgId)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewDBErr(), "find chatlog by msgId err %v, req %v", err, in.MsgId)
		}

		return &im.GetChatLogResp{
			List: []*im.ChatLog{
				{
					Id:             chatLog.ID.Hex(),
					ConversationId: chatLog.ConversationId,
					SendId:         chatLog.SendId,
					RecvId:         chatLog.RecvId,
					MsgType:        int32(chatLog.MsgType),
					MsgContent:     chatLog.MsgContent,
					ChatType:       int32(chatLog.ChatType),
					SendTime:       chatLog.SendTime,
					ReadRecords:    chatLog.ReadRecords,
				},
			},
		}, nil
	}
	// 时间段分段查询
	data, err := l.svcCtx.ChatLogModel.ListBySendTime(l.ctx, in.ConversationId, in.StartSendTime, in.EndSendTime, in.Count)
	if err != nil {
		log.Printf("数据库查询失败: %v", err)
		return nil, errors.Wrapf(xerr.NewDBErr(), "ListBySendTime err %v, req %v", err, in)
	}
	log.Printf("数据库查询结果: %+v", data) // 确认数据库是否有数据
	res := make([]*im.ChatLog, 0, len(data))
	for _, datum := range data {
		res = append(res, &im.ChatLog{
			Id:             datum.ID.Hex(),
			ConversationId: datum.ConversationId,
			SendId:         datum.SendId,
			RecvId:         datum.RecvId,
			MsgType:        int32(datum.MsgType),
			MsgContent:     datum.MsgContent,
			ChatType:       int32(datum.ChatType),
			SendTime:       datum.SendTime,
			ReadRecords:    datum.ReadRecords,
		})
	}

	return &im.GetChatLogResp{
		List: res,
	}, nil
}
