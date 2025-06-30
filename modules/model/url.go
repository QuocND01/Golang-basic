package model

type Url struct {
	Id        int    `json:"id" gorm:"column:id"`
	Originurl string `json:"originurl" gorm:"column:originurl"`
	Sorturl   string `json:"sorturl" gorm:"column:sorturl"`
}

func (Url) TableName() string {
	return "url"
}

type Urladd struct {
	Originurl string `json:"originurl" gorm:"column:originurl"`
	Sorturl   string `json:"sorturl" gorm:"column:sorturl"`
}

func (Urladd) TableName() string {
	return Url{}.TableName()
}
