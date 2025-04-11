package conversation

import (
	"easy-chat/apps/im/ws/internal/svc"
	"easy-chat/apps/im/ws/websocket"
	"easy-chat/apps/im/ws/ws"
	"easy-chat/apps/task/mq/mq"
	"easy-chat/pkg/constants"
	"easy-chat/pkg/wuid"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"

	"time"
)

func Chat(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {
		// todo: 私聊
		var data ws.Chat
		if err := mapstructure.Decode(msg.Data, &data); err != nil {

			srv.Send(websocket.NewErrMessage(err), conn)
			return
		}
		if data.MType == constants.ImageMtype || data.MType == constants.FileMtype {
			if data.Content == "" || !strings.HasPrefix(data.Content, "http") {
				srv.Send(websocket.NewErrMessage(errors.New("无效的文件 URL")), conn)
				return
			}
		}

		if data.ConversationId == "" {
			switch data.ChatType {
			case constants.SingleChatType:
				data.ConversationId = wuid.CombineId(conn.Uid, data.RecvId)
			case constants.GroupChatType:
				data.ConversationId = data.RecvId
			}
		}

		//err := logic.NewConversation(context.Background(), srv, svc).SingleChat(&data, conn.Uid)
		//if err != nil {
		//	srv.Send(websocket.NewErrMessage(err), conn)
		//	return
		//}
		//srv.SendByUserId(websocket.NewMessage(conn.Uid, ws.Chat{
		//	ConversationId: data.ConversationId,
		//	ChatType:       data.ChatType,
		//	SendId:         conn.Uid,
		//	RecvId:         data.RecvId,
		//	SendTime:       time.Now().UnixMilli(),
		//	Msg:            data.Msg,
		//}), data.RecvId)

		err := svc.MsgChatTransferClient.Push(&mq.MsgChatTransfer{
			ConversationId: data.ConversationId,
			ChatType:       data.ChatType,
			SendId:         conn.Uid,
			RecvId:         data.RecvId,
			SendTime:       time.Now().UnixMilli(),
			MType:          data.Msg.MType,
			Content:        data.Msg.Content,
			FileName:       data.Msg.FileName,
			FileSize:       data.Msg.FileSize,
			MsgId:          msg.Id,
		})
		if err != nil {
			srv.Send(websocket.NewErrMessage(err), conn)
			return
		}
	}
}

func MarkRead(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {
		// todo: 已读未读处理
		var data ws.MarkRead
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			srv.Send(websocket.NewErrMessage(err), conn)
			return
		}
		err := svc.MsgReadTransferClient.Push(&mq.MsgMarkRead{
			ChatType:       data.ChatType,
			ConversationId: data.ConversationId,
			SendId:         conn.Uid,
			RecvId:         data.RecvId,
			MsgIds:         data.MsgIds,
		})
		if err != nil {
			srv.Send(websocket.NewErrMessage(err), conn)
			return
		}

	}
}

/*func saveFileLocally(fileContent string, fileName string) (string, error) {
	// 文件存储目录
	dir := "D:\\Gocode\\src\\easy-chat\\apps\\im\\ws\\uploads"

	// 如果目录不存在，则创建
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return "", fmt.Errorf("创建文件夹失败: %w", err)
		}
	}

	// 将文件内容从 Base64 转换为字节数组
	decodedData, err := os.ReadFile(fileContent) // 如果是 Base64 编码，解码即可
	if err != nil {
		return "", fmt.Errorf("读取文件内容失败: %w", err)
	}

	// 保存文件到本地
	filePath := filepath.Join(dir, fileName)
	err = os.WriteFile(filePath, decodedData, 0644)
	if err != nil {
		return "", fmt.Errorf("保存文件失败: %w", err)
	}

	// 返回文件访问的 URL
	fileURL := fmt.Sprintf("http://localhost:10010/uploads/%s", fileName)
	return fileURL, nil
}
*/
