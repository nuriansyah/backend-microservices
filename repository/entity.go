package repository

type User struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Role    string  `json:"role"`
	Nrp     string  `json:"nrp"`
	Prodi   string  `json:"prodi"`
	Program string  `json:"institute"`
	Company *string `json:"major"`
	Batch   *int    `json:"batch"`
	Avatar  *string `json:"avatar"`
}
