package models

import(
    "time"
)

type Interviewer struct{
    Id string `json:"id" form:"-"`
    UserId string `json:"user" form:"user"`
    Summary string `json:"summary" form:"summary"`
    Ranking int `json:"ranking" form:"ranking"`
    DateJoined time.Time `json:"uuid" form:"joined"`
}
