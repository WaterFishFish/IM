package ws

import "easy-chat/pkg/constants"

type (
	Msg struct {
		MsgId           string            `mapstructure:"msgId"`
		ReadRecords     map[string]string `mapstructure:"readRecords"`
		constants.MType `mapstructure:"mType"`
		FileName        string `mapstructure:"fileName,omitempty"` // 文件名（图片或文件）
		FileSize        int64  `mapstructure:"fileSize,omitempty"` // 文件大小（单位: byte）
		Content         string `mapstructure:"content"`
	}

	Chat struct {
		ConversationId     string `mapstructure:"conversationId"`
		constants.ChatType `mapstructure:"chatType"`
		SendId             string `mapstructure:"sendId"`
		RecvId             string `mapstructure:"recvId"`
		SendTime           int64  `mapstructure:"sendTime"`
		Msg                `mapstructure:"msg"`
	}

	Push struct {
		ConversationId     string `mapstructure:"conversationId"`
		constants.ChatType `mapstructure:"chatType"`
		SendId             string   `mapstructure:"sendId"`
		RecvId             string   `mapstructure:"recvId"`
		RecvIds            []string `mapstructure:"recvIds"`
		SendTime           int64    `mapstructure:"sendTime"`

		MsgId           string                `mapstructure:"msgId"`
		ReadRecords     map[string]string     `mapstructure:"readRecords"`
		ContentType     constants.ContentType `mapstructure:"contentType"`
		constants.MType `mapstructure:"mType"`
		Content         string `mapstructure:"content"`
		FileName        string `json:"fileName,omitempty" mapstructure:"fileName,omitempty"`
		FileSize        int64  `json:"fileSize,omitempty" mapstructure:"fileSize,omitempty"`
	}
	MarkRead struct {
		constants.ChatType `mapstructure:"chatType"`
		RecvId             string   `mapstructure:"recvId"`
		ConversationId     string   `mapstructure:"conversationId"`
		MsgIds             []string `mapstructure:"msgIds"`
	}
)
