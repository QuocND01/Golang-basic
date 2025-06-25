package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

// Tạo hằng số đại diện cho trạng thái Product
type ProductStatus int

// Định nghĩa 2 giá trị hằng số thuộc kiểu ProductStatus bằng cách dùng iota cho tự tăng giá trị.
const (
	StatusActive ProductStatus = iota
	StatusInactive
)

// Khai báo các error
var (
	ErrProductnameIsBlank          = errors.New("Product name cannot be blank")
	ErrProductnameIsInvalid        = errors.New("Product name cannot contain special characters")
	ErrProductdescriptionIsBlank   = errors.New("Product description cannot be blank")
	ErrProductdescriptionIsInvalid = errors.New("Product description cannot contain special characters")
	ErrProductpriceIsInvalid       = errors.New("Product price must be a positive number")
	ErrProductcategoryIsBlank      = errors.New("Product category cannot be blank")
	ErrBookIsDeleted               = errors.New("this book is deleted")
)

// Tạo mảng chứa các trạng thái của Product
var Statuses = [2]string{"Active", "Inactive"}

// Chuyển string thành enum ProductStatus với giá trị int
func parseStr2Status(s string) (ProductStatus, error) {
	for i := range Statuses {
		if Statuses[i] == s {
			return ProductStatus(i), nil
		}
	}
	return ProductStatus(0), errors.New("Invalid status string")
}

// Lấy chuỗi tương ứng của giá trị ProductStatus, dựa trên mảng Statuses.
func (b *ProductStatus) String() string {
	return Statuses[*b]
}

// Custom phương thức Scan() để chuyển dữ liệu dạng string từ database thành enum ProductStatus
func (p *ProductStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("Fail to scan data from sql: %s", value))
	}
	v, err := parseStr2Status(string(bytes))
	if err != nil {
		return errors.New(fmt.Sprintf("Fail to scan data from sql: %s", value))
	}
	*p = v
	return nil

}

// Custom phương thức Value() để chuyển dữ liệu dạng enum thành string
func (p *ProductStatus) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return p.String(), nil
}

// Chuyển từ enum thành byte chứa json
func (p *ProductStatus) MarshalJSON() ([]byte, error) {
	if p == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", p.String())), nil
}

// Chuyển từ byte chứa json thành enum
func (p *ProductStatus) UnmarshalJSON(data []byte) error {
	s := strings.ReplaceAll(string(data), "\"", "")
	result, err := parseStr2Status(s)
	if err != nil {
		return err
	}
	*p = result
	return nil
}

// Định nghĩa struct model Product
type Product struct {
	Productid     int            `json:"productid"  gorm:"column:productid;primaryKey"`
	Productname   string         `json:"productname" gorm:"column:productname;" binding:"required"`
	Description   string         `json:"description" gorm:"column:description;" binding:"required"`
	Price         float32        `json:"price" gorm:"column:price;" binding:"required,gt=0"`
	ProductStatus *ProductStatus `json:"productstatus" gorm:"column:productstatus;"`
	Created_at    *time.Time     `json:"created_at" gorm:"created_at;"`
	Update_at     *time.Time     `json:"update_at" gorm:"update_at;"`
	Categoryid    int            `json:"categoryid" gorm:"categoryid;" binding:"required"`
}

// Cho GORM biết mapping struct này với table categories trong database.
func (Product) TableName() string {
	return "product"
}

// Định nghĩa struct model Product để in với Categoryname
type ProductWithCatename struct {
	Productid     int            `json:"productid"  gorm:"column:productid;primaryKey"`
	Productname   string         `json:"productname" gorm:"column:productname;" binding:"required"`
	Description   string         `json:"description" gorm:"column:description;" binding:"required"`
	Price         float32        `json:"price" gorm:"column:price;" binding:"required,gt=0"`
	ProductStatus *ProductStatus `json:"productstatus" gorm:"column:productstatus;"`
	Created_at    *time.Time     `json:"created_at" gorm:"created_at;"`
	Update_at     *time.Time     `json:"update_at" gorm:"update_at;"`
	Categoryname  string         `json:"categoryname"`
}

// Cho GORM biết mapping struct này với table categories trong database.
func (ProductWithCatename) TableName() string {
	return Product{}.TableName()
}

// Định nghĩa struct model Product để thêm Product mới
type AddProduct struct {
	Productid   int     `json:"productid"  gorm:"column:productid;"`
	Productname string  `json:"productname" gorm:"column:productname;" binding:"required"`
	Description string  `json:"description" gorm:"column:description;" binding:"required"`
	Price       float32 `json:"price" gorm:"column:price;" binding:"required,gt=0"`
	Categoryid  int     `json:"categoryid" gorm:"categoryid;" binding:"required"`
}

// Cho GORM biết mapping struct này với table categories trong database.
func (AddProduct) TableName() string {
	return Product{}.TableName()
}

// Định nghĩa struct model Product để update Product
type UpdateProduct struct {
	Productname   string         `json:"productname" gorm:"column:productname;" binding:"required"`
	Description   string         `json:"description" gorm:"column:description;" binding:"required"`
	Price         float32        `json:"price" gorm:"column:price;" binding:"required,gt=0"`
	ProductStatus *ProductStatus `json:"productstatus" gorm:"column:productstatus;"`
	Categoryid    int            `json:"categoryid" gorm:"categoryid;" binding:"required"`
	Update_at     *time.Time     `json:"update_at" gorm:"update_at;"`
}

// Cho GORM biết mapping struct này với table categories trong database.
func (UpdateProduct) TableName() string {
	return Product{}.TableName()
}
