package constants

type MType int

const (
	TextMtype  MType = iota // 0 - 文本消息
	ImageMtype              // 1 - 图片消息
	FileMtype               // 2 - 文件消息
)

type ChatType int

const (
	GroupChatType ChatType = iota + 1
	SingleChatType
)

type ContentType int

const (
	ContentChatMsg ContentType = iota
	ContentMakeRead
)
