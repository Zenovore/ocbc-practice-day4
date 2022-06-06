package model

type Employee struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Dob      string `json:"dob"`
	Address  string `json:"address"`
	Jobs     string `json:"jobs"`
	JoinDate string `json:"join_date"`
}
