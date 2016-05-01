package repositories

import (
    "hoditgo/core/mysql"
    "hoditgo/services/models"
    "golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
}

var userRepository *UserRepository = nil

func InitUserRepository() *UserRepository {
    if userRepository == nil {
        userRepository = &UserRepository{
        }
    }

    return userRepository
}


func (repo *UserRepository) CreateUser(user *models.User) {
    var err error
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
    dbConn := mysql.Connect()
    stmt, err := dbConn.Prepare("INSERT Users SET Id=?,Username=?,Password=?,Email=?")
    checkErr(err)
    _, err = stmt.Exec(util.GenerateUUID(), user.Username, hashedPassword, user.Email)
    checkErr(err)
}

func (repo *UserRepository) CheckUser(email string, password string) (exists bool){
    var Email string 
    var err error
    dbConn := mysql.Connect()
    rows, err := dbConn.Query("Select Email from Users Where Email = ?", email)
    checkErr(err)
    defer rows.Close()
    for rows.Next(){
        err = rows.Scan(&Email)
        checkErr(err)
    }

    return Email == email    
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

