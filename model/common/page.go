package common

type BasePage struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	Order    string `json:"order"`
	Desc     bool   `json:"desc"`
}

func (s *BasePage) GetPage() (page int) {
	page = s.Page
	if page != 0 {
		page -= 1
	}
	return
}

func (s *BasePage) GetOffset() (offset int) {
	page := s.GetPage()
	offset = page * s.PageSize
	return
}
