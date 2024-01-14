package jsondemo

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title,omitempty"`
	Completed *bool  `json:"completed,omitempty"`
}
