package models

type Categories struct{
    Id string `json:"uuid" form:"-"`
    CategoryName string `json:"category" form:"category"`
    ParentId string `json:"parent" form:"parent"`
}