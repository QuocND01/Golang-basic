package biz

import (
	"context"
	"myproject/modules/model"
	"regexp"
	"strings"
	"unicode"
)

// Tách tầng business khỏi tầng storage, quy định phương thức mà tầng storage phải có
type AddProductStorage interface {
	AddProduct(ctx context.Context, data *model.AddProduct) error
}

// Cho phép bất kỳ struct nào implement interface AddProductStorage đều có thể truyền vào đây.
type addProductBiz struct {
	store AddProductStorage
}

// Khởi tạo và truy cập struct addProductBiz mà bản thân nó là private
func AddProductBiz(store AddProductStorage) *addProductBiz {
	return &addProductBiz{store: store}
}

// Phương thức chính của biz gắn với addProductBiz struct gọi tới phương thức AddProduct ở tầng storage
func (biz *addProductBiz) CreateNewItem(ctx context.Context, data *model.AddProduct) error {
	Productname := strings.TrimSpace(data.Productname)
	Description := strings.TrimSpace(data.Description)

	if matched, _ := regexp.MatchString("^[a-zA-Z0-9 ]+$", Productname); !matched {
		return model.ErrProductnameIsInvalid
	}

	valid := true
	for _, ch := range Description {
		if !(unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == ' ') {
			valid = false
			break
		}
	}
	if !valid {
		return model.ErrProductdescriptionIsInvalid
	}
	if err := biz.store.AddProduct(ctx, data); err != nil {
		return err
	}
	return nil
}
