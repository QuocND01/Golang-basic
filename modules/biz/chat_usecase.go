package biz

import (
	"context"
	"myproject/modules/model"
)

// Tách tầng business khỏi tầng storage, quy định phương thức mà tầng storage phải có
type MessageStorage interface {
	SaveMessage(ctx context.Context, msg *model.Message) error
	ListMessages(ctx context.Context) ([]model.Message, error)
}

// Cho phép bất kỳ struct nào implement interface MessageStorage đều có thể truyền vào đây.
type MessageBiz struct {
	store MessageStorage
}

// Khởi tạo và truy cập struct MessageBiz mà bản thân nó là private
func NewMessageBiz(store MessageStorage) *MessageBiz {
	return &MessageBiz{store: store}
}

// Phương thức chính của biz gắn với MessageBiz struct gọi tới phương thức SaveMessage ở tầng storage
func (biz *MessageBiz) SaveMessage(ctx context.Context, msg *model.Message) error {
	return biz.store.SaveMessage(ctx, msg)
}

// Phương thức chính của biz gắn với MessageBiz struct gọi tới phương thức ListMessages ở tầng storage
func (biz *MessageBiz) ListMessages(ctx context.Context) ([]model.Message, error) {
	return biz.store.ListMessages(ctx)
}
