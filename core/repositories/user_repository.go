package repositories

import (
    "hoditgo/core/mysql"
    "hoditgo/services/models"
    "golang.org/x/crypto/bcrypt"
    "time"
    "database/sql"
)

type UserRepository struct {
    
     dbConnection *sql.DB
}

var userRepository *UserRepository = nil

func InitUserRepository() *UserRepository {
    if userRepository == nil {
        userRepository = &UserRepository{
        }
    }
    
    userRepository.dbConnection = mysql.Connect()
    return userRepository
}


func (repo *UserRepository) CreateUser(user *models.User) {
    var err error
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
    stmt, err :=  userRepository.dbConnection.Prepare("INSERT INTO Users SET Id=?,Username=?,Password=?,Email=?, Name=?, Surname=?, DateJoined=?, IsActive=?")
    util.CheckErr(err)
    _, err = stmt.Exec(util.GenerateUUID(), user.Username, hashedPassword, user.Email, user.Name, user.Surname, time.Now().Format(time.RFC3339), true)
    util.CheckErr(err)
}

func (repo *UserRepository) CheckUser(email string, password string) (exists bool){
    var Email string 
    var err error
    rows, err :=  userRepository.dbConnection.Query("Select Email from Users Where Email = ?", email)
    util.CheckErr(err)
    defer rows.Close()
    for rows.Next(){
        err = rows.Scan(&Email)
        util.CheckErr(err)
    }

    return Email == email    
}


