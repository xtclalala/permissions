package common

type BasePage struct {
	Page     int  `json:"page"    validate:"omitempty,min=0"   label:"页数"`
	PageSize int  `json:"pageSize"  validate:"omitempty,lt=50"   label:"分页大小"`
	Desc     bool `json:"desc" query:"desc"`
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
