package models

type InterviewResult struct{
    Id string `json:"id" form:"-"`
    Pros string `json:"pros" form:"pros"`
    Cons string `json:"cons" form:"cons"`
    EvaluationReport string `json:"report" form:"report"`
    Point int `json:"point" form:"point"`
    InterviewSchedule InterviewSchedule `json:"interview" form:"interview"`
}