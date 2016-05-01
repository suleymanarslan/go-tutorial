package models

import(
    "time"
)

type Interviewer struct{
    Id string `json:"uuid" form:"-"`
    UserId string `json:"uuid" form:"user"`
    Ranking int `json:"uuid" form:"ranking"`
    DateJoined time.Time `json:"uuid" form:"joined"`
}
