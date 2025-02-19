package domain

type (
	News struct {
		ID         int32   `json:"Id"`
		Title      string  `reform:"Title"`
		Content    string  `json:"Content"`
		Categories []int32 `reform:"Categories"`
	}
)
