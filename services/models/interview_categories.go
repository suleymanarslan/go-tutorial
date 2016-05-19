package models

type Categories struct{
    Id string `json:"id" form:"-"`
    CategoryName string `json:"category" form:"category"`
    ParentId string `json:"parent" form:"parent"`
}