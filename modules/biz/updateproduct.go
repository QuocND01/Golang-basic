package biz

import (
	"context"
	"myproject/modules/model"
)

// Tách tầng business khỏi tầng storage, quy định phương thức mà tầng storage phải có
type UpdateProductStogage interface {
	GetProductBase(ctx context.Context, cond map[string]interface{}) (*model.Product, error)
	UpdateProduct(ctx context.Context, cond map[string]interface{}, updata *model.UpdateProduct) error
}

// Cho phép bất kỳ struct nào implement interface UpdateProductStogage đều có thể truyền vào đây.
type updateProductBiz struct {
	store UpdateProductStogage
}

// Khởi tạo và truy cập struct updateProductBiz mà bản thân nó là private
func UpdateProductBiz(store UpdateProductStogage) *updateProductBiz {
	return &updateProductBiz{store: store}
}

// Phương thức chính của biz gắn với updateProductBiz struct gọi tới phương thức GetProductBase và UpdateProduct ở tầng storage
func (biz *updateProductBiz) UpdateProduct(ctx context.Context, id int, updata *model.UpdateProduct) error {
	data, err := biz.store.GetProductBase(ctx, map[string]interface{}{"productid": id})
	if err != nil {
		return err
	}
	if data.ProductStatus != nil && *data.ProductStatus == model.StatusInactive {
		return model.ErrBookIsDeleted
	}
	if err := biz.store.UpdateProduct(ctx, map[string]interface{}{"productid": id}, updata); err != nil {
		return err
	}

	return nil
}
