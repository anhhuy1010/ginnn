package getlisttodo

type (
	GetListReq struct {
		Task      *string `form:"task"`
		SortBy    *string `form:"sortby"`
		SortOrder *string `form:"sortorder"`
		Limit     *int64  `form:"limit"`
		Skip      *int64  `form:"skip"`
	}
	GetListRes struct {
		Uuid string `json:"uuid"`
		Task string `json:"task"`
	}
)
