package sortstudent

type Req struct {
	SortBy    string `json:"sortby" binding:"required"`
	SortOrder string `json:"sortorder" binding:"required"`
}
