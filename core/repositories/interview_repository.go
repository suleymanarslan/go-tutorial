package repositories

import (
    "hoditgo/core/mysql"
    "hoditgo/services/models"
    "hoditgo/api"
    "database/sql"

)

type InterviewRepository struct {
   dbConnection *sql.DB
}

var util api.Utils
var interviewRepository *InterviewRepository = nil

func InitInterviewRepository() *InterviewRepository {
    if interviewRepository == nil {
        interviewRepository = &InterviewRepository{
        }
    }
    
    interviewRepository.dbConnection = mysql.Connect()
    return interviewRepository
}


func (repo *InterviewRepository) CreateInterview(interview *models.Interview) {
    var err error
    stmt, err := interviewRepository.dbConnection.Prepare("INSERT INTO Interviews SET Id=?, Name=? CategoryId=?,Description=?,InterviewerId=?, IsFeatured=?")
    util.CheckErr(err)
    _, err = stmt.Exec(util.GenerateUUID(), interview.CategoryId, interview.Description, interview.InterviewerId, interview.IsFeatured)
    util.CheckErr(err)
}

func (repo *InterviewRepository) UpdateInterview(interview *models.Interview) {
      var err error
    stmt, err := interviewRepository.dbConnection.Prepare("UPDATE Interviews SET Name = ?, CategoryId = ?, Description = ?, IsFeatured = ?, IsActive = ? WHERE Id = ?")
    util.CheckErr(err)
    _, err = stmt.Exec(interview.Name, interview.CategoryId, interview.Description, interview.IsFeatured, interview.Id, true)
    util.CheckErr(err)
}


func (repo *InterviewRepository) GetInterviewById(id string) (models.Interview) {
    var result models.Interview
    
    stmt, err := interviewRepository.dbConnection.Prepare("SELECT Name, CategoryId, Description, InterviewerId, IsFeatured FROM Interviews WHERE Id = ? AND IsActive = ?")
    util.CheckErr(err)
    row := stmt.QueryRow(id, true)
    err = row.Scan(&result.Name, &result.CategoryId, &result.Description, &result.InterviewerId, &result.IsFeatured)
    util.CheckErr(err)
    return result
}

func (repo *InterviewRepository) GetInterviewByName(name string) ([]models.Interview) {
    var  results []models.Interview
    var result models.Interview
    
    stmt, err := interviewRepository.dbConnection.Prepare("SELECT Name, CategoryId, Description, InterviewerId, IsFeatured FROM Interviews WHERE Name = ? AND IsActive = ?")
    util.CheckErr(err)
    rows, err := stmt.Query(name, true)
    
    for rows.Next() {
        err = rows.Scan(&result.Name, &result.CategoryId, &result.Description, &result.InterviewerId, &result.IsFeatured)
        results = append(results, result)
    }
    
    util.CheckErr(err)
    return results
}

func (repo *InterviewRepository) GetAllInterviews(offset int) ([]models.Interview){
    var  results []models.Interview
    var result models.Interview

    stmt, err := interviewRepository.dbConnection.Prepare("SELECT Name, CategoryId, Description, InterviewerId, IsFeatured FROM Interviews WHERE AND IsActive = ?")
    util.CheckErr(err)
    rows, err := stmt.Query()
    for rows.Next() {
        err = rows.Scan(&result.Name, &result.CategoryId, &result.Description, &result.InterviewerId, &result.IsFeatured)
        results = append(results, result)
    }
    
    util.CheckErr(err)
    return results
}

func (repo *InterviewRepository) DeactivateInterview(id string) {
          var err error
    stmt, err := interviewRepository.dbConnection.Prepare("UPDATE Interviews SET IsActive = ? WHERE Id = ? AND IsActive = ?")
    util.CheckErr(err)
    _, err = stmt.Exec(id, false, true)
    util.CheckErr(err)
}
