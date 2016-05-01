package models

type Interview struct {
	Id     string `json:"uuid" form:"-"`
    CategoryId string `json:"categoryId" form:"categoryId"`
    InterviewerId string `json:"interviewerId" form:"interviewerId"`
    IsFeatured string `json:"featured" form:"featured"`
    Description string `json:"description" form:"description"`
}
