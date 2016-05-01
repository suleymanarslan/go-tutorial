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
    stmt, err := dbConn.Prepare("INSERT Interviews SET Id=?, Name=? CategoryId=?,Description=?,InterviewerId=?, IsFeatured=?")
    util.CheckErr(err)
    _, err = stmt.Exec(util.GenerateUUID(), interview.CategoryId, interview.Description, interview.InterviewerId, interview.IsFeatured)
    util.CheckErr(err)
}

func (repo *InterviewRepository) UpdateInterview(interview *models.Interview) {
      var err error
    dbConn := mysql.Connect()
    stmt, err := dbConn.Prepare("UPDATE Interviews SET (Name, CategoryId, Description, IsFeatured, IsActive) VALUES(?,?,?,?,?) WHERE Id = ?")
    util.CheckErr(err)
    _, err = stmt.Exec(interview.Name, interview.CategoryId, interview.Description, interview.IsFeatured, interview.Id, true)
    util.CheckErr(err)
}


func (repo *InterviewRepository) GetInterviewById(id string) {
      var err error
    dbConn := mysql.Connect()
    stmt, err := dbConn.Prepare("SELECT Name, CategoryId, Description, InterviewerId, IsFeatured WHERE Id = ? AND IsActive = ?")
    util.CheckErr(err)
    _, err = stmt.Exec(id, true)
    util.CheckErr(err)
}

func (repo *InterviewRepository) GetInterviewByName(name string) {
      var err error
    dbConn := mysql.Connect()
    stmt, err := dbConn.Prepare("SELECT Name, CategoryId, Description, InterviewerId, IsFeatured WHERE Name = ? AND IsActive = ?")
    util.CheckErr(err)
    _, err = stmt.Exec(name, true)
    util.CheckErr(err)
}

func (repo *InterviewRepository) DeactivateInterview(id string) {
          var err error
    dbConn := mysql.Connect()
    stmt, err := dbConn.Prepare("UPDATE Interviews SET (IsActive) VALUES(?) WHERE Id = ? AND IsActive = ?")
    util.CheckErr(err)
    _, err = stmt.Exec(id, false, true)
    util.CheckErr(err)
}