package repository

type Mahasiswa struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Avatar *string `json:"avatar"`
}

type Dosen struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Avatar *string `json:"avatar"`
}

type Log struct {
	Id   int    `json:"id"`
	desc string `json:"desc"`
}
