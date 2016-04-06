package repositories

import (
    "hoditgo/core/mysql"
    "hoditgo/services/models"
    "golang.org/x/crypto/bcrypt"
    "fmt"
    "crypto/rand"
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
    stmt, err := dbConn.Prepare("INSERT users SET UUID=?,Username=?,Password=?,Email=?")
    checkErr(err)
    _, err = stmt.Exec(pseudo_uuid(), user.Username, hashedPassword, user.Email)
    checkErr(err)
}

func (repo *UserRepository) CheckUser(email string, password string) (exists bool){
    var Email string 
    var err error
    dbConn := mysql.Connect()
    rows, err := dbConn.Query("Select Email from users Where Email = ?", email)
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

func pseudo_uuid() (uuid string) {

    b := make([]byte, 16)
    _, err := rand.Read(b)
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }

    uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

    return
}



