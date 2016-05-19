package repositories

import (
    "hoditgo/core/mysql"
    "database/sql"
    "hoditgo/services/models"
    "time"
)

type InterviewerRepository struct {
       dbConnection *sql.DB
}


var interviewerRepository *InterviewerRepository = nil

func InitInterviewerRepo() *InterviewerRepository {
    if interviewerRepository == nil {
        interviewerRepository = &InterviewerRepository{
        }
    }
    interviewerRepository.dbConnection = mysql.Connect()
    return interviewerRepository
}


func (repo *InterviewerRepository) CreateInterviewer(interviewer *models.Interviewer) {
    var err error
    stmt, err := interviewerRepository.dbConnection.Prepare("INSERT INTO Interviewers SET Id=?, DateJoined=?, UserId=?,Ranking=?, Summary=?")
    util.CheckErr(err)
    _, err = stmt.Exec(util.GenerateUUID(), time.Now().Format(time.RFC3339), interviewer.User.Id, -1, interviewer.Summary)
    util.CheckErr(err)
}

func (repo *InterviewerRepository) UpdateInterviewerSummary(summary, id string) {
      var err error
    stmt, err := interviewerRepository.dbConnection.Prepare("UPDATE Interviewers SET Summary = ? WHERE Id = ?")
    util.CheckErr(err)
    _, err = stmt.Exec(summary, id)
    util.CheckErr(err)
}

func (repo *InterviewerRepository) UpdateInterviewerRanking(ranking int, id string) {
      var err error
    stmt, err := interviewerRepository.dbConnection.Prepare("UPDATE Interviewers SET Ranking = ? WHERE Id = ?")
    util.CheckErr(err)
    _, err = stmt.Exec(ranking, id)
    util.CheckErr(err)
}


func (repo *InterviewerRepository) GetInterviewerById(id string) (models.Interviewer){
    
    var result models.Interviewer
    stmt, err := interviewerRepository.dbConnection.Prepare(`SELECT Interviewer.Id, 
                                                            User.Username, 
                                                            User.Id,
                                                            User.Name,
                                                            User.Surname,
                                                            User.Datejoined,
                                                            Interviewer.Ranking, 
                                                            Interviewer.DateJoined, 
                                                            Interviewer.Summary 
                                                            FROM Interviewers as Interviewer
                                                            INNER JOIN Users as User on 
                                                            User.Id = Interviewer.UserId
                                                            WHERE Interviewer.Id = ?`)
    util.CheckErr(err)
    
    row := stmt.QueryRow(id)
    
    err = row.Scan(&result.Id, 
                    &result.User.Username,  
                    &result.User.Id,
                    &result.User.Name, 
                    &result.User.Surname, 
                    &result.User.DateJoined, 
                    &result.Ranking, 
                    &result.DateJoined, 
                    &result.Summary)
                    
    util.CheckErr(err)
    return result
}

func (repo *InterviewerRepository) GetInterviewerByName(name string) ([]models.Interviewer){
    var results  []models.Interviewer
        stmt, err := interviewerRepository.dbConnection.Prepare(`SELECT Interviewer.Id, 
                                                            User.Username, 
                                                            User.Id,
                                                            User.Name,
                                                            User.Surname,
                                                            User.Datejoined,
                                                            Interviewer.Ranking, 
                                                            Interviewer.DateJoined, 
                                                            Interviewer.Summary 
                                                            FROM Interviewers as Interviewer
                                                            INNER JOIN Users as User on 
                                                            User.Id = Interviewer.UserId
                                                            WHERE User.Username = ?`)
    util.CheckErr(err)
    rows, err := stmt.Query(name)
    
    for rows.Next(){
        var result models.Interviewer
        err = rows.Scan(&result.Id, 
                    &result.User.Username,  
                    &result.User.Id,
                    &result.User.Name, 
                    &result.User.Surname, 
                    &result.User.DateJoined, 
                    &result.Ranking, 
                    &result.DateJoined, 
                    &result.Summary)
        
        results = append(results, result)
    }
    
    util.CheckErr(err)
    
    return results
}

