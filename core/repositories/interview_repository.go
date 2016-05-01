package repositories

import (
    "hoditgo/core/mysql"
    "hoditgo/services/models"
    "hoditgo/api"
)

type InterviewRepository struct {
}

var util api.Utils

var interviewRepository *InterviewRepository = nil

func InitInterviewRepo() *InterviewRepository {
    if interviewRepository == nil {
        interviewRepository = &InterviewRepository{
        }
    }

    return interviewRepository
}


func (repo *InterviewRepository) CreateInterview(interview *models.Interview) {
    var err error
    dbConn := mysql.Connect()
    stmt, err := dbConn.Prepare("INSERT Interviews SET Id=?,CategoryId=?,Description=?,InterviewerId=?, IsFeatured=?")
    checkErr(err)
    _, err = stmt.Exec(util.GenerateUUID(), interview.CategoryId, interview.Description, interview.InterviewerId, interview.IsFeatured)
    checkErr(err)
}