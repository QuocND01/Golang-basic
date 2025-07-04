package paging

// Để đảm bảo rằng thông tin phân trang luôn hợp lệ, không bị sai
type Paging struct {
	Page  int   `json:"page" form: "page"`
	Limit int   `json:"limit" form: "limit"`
	Total int64 `json:"total" form:"-"`
}

func (p *Paging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 || p.Limit >= 100 {
		p.Limit = 10
	}
}
