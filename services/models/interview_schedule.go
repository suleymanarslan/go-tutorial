package models

type InterviewSchedule struct{
    Id string `json:"uuid" form:"-"`
    User User `json:"user" form:"user"`
    Interviewer Interviewer `json:"interviewer" form:"interviewer"`
    InterviewDate string `json:"date" form:"date"`
}
