package resources

type Pagination struct {
	Total       int         `json:"total"`
	PerPage     int         `json:"per_page"`
	CurrentPage int         `json:"current_page"`
	LastPage    int         `json:"last_page"`
	NextPageURL string      `json:"next_page_url"`
	PrevPageURL string      `json:"prev_page_url"`
	From        int         `json:"from"`
	To          int         `json:"to"`
	Data        interface{} `json:"data"`
}

func (p Pagination) GetOffset() int {
	return (p.CurrentPage - 1) * p.PerPage
}
