package biz

import (
	"context"
	"myproject/common/paging"
	"myproject/modules/model"
)

// Tách tầng business khỏi tầng storage, quy định phương thức mà tầng storage phải có
type GetListProductStogage interface {
	GetProducts(ctx context.Context, filter *model.Filter, paging *paging.Paging, moreKeys ...string) ([]model.ProductWithCatename, error)
}

// Cho phép bất kỳ struct nào implement interface GetListProductStogage đều có thể truyền vào đây.
type getListProductBiz struct {
	store GetListProductStogage
}

// Khởi tạo và truy cập struct getListProductBiz mà bản thân nó là private
func GetListProductBiz(store GetListProductStogage) *getListProductBiz {
	return &getListProductBiz{store: store}
}

// Phương thức chính của biz gắn với getListProductBiz struct gọi tới phương thức GetProducts ở tầng storage
func (biz *getListProductBiz) GetProducts(ctx context.Context, filter *model.Filter, paging *paging.Paging, moreKeys ...string) ([]model.ProductWithCatename, error) {
	data, err := biz.store.GetProducts(ctx, filter, paging)

	if err != nil {
		return nil, err
	}
	return data, nil
}
