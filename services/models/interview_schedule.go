package models

type InterviewSchedule struct{
    Id string `json:"uuid" form:"-"`
    UserId string `json:"user" form:"user"`
    Interviewerid string `json:"interviewer" form:"interviewer"`
    InterviewDate string `json:"date" form:"date"`
}
