package biz

import (
	"context"
	"myproject/modules/model"
)

// Tách tầng business khỏi tầng storage, quy định phương thức mà tầng storage phải có
type GetProductStogage interface {
	GetProduct(ctx context.Context, cond map[string]interface{}) (*model.ProductWithCatename, error)
}

// Cho phép bất kỳ struct nào implement interface GetProductStogage đều có thể truyền vào đây.
type getProductBiz struct {
	store GetProductStogage
}

// Khởi tạo và truy cập struct getProductBiz mà bản thân nó là private
func GetProductBiz(store GetProductStogage) *getProductBiz {
	return &getProductBiz{store: store}
}

// Phương thức chính của biz gắn với getProductBiz struct gọi tới phương thức GetProduct ở tầng storage
func (biz *getProductBiz) GetProductByID(ctx context.Context, id int) (*model.ProductWithCatename, error) {
	data, err := biz.store.GetProduct(ctx, map[string]interface{}{"productid": id})

	if err != nil {
		return nil, err
	}
	return data, nil
}
