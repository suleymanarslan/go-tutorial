package models

type Interview struct {
	Id     string `json:"uuid" form:"-"`
    Name string `json:"name" form:"name"`
    Category Categories `json:"category" form:"category"`
    Interviewer Interviewer `json:"interviewer" form:"interviewer"`
    IsFeatured string `json:"featured" form:"featured"`
    Description string `json:"description" form:"description"`
}
