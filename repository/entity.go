package repository

type User struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Role    string  `json:"role"`
	Program *string `json:"program"`
	Company *string `json:"company"`
	Batch   *int    `json:"batch"`
	Avatar  *string `json:"avatar"`
}
