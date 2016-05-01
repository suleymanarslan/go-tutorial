package api


import (
    "crypto/rand"
    "fmt"
)

type Utils struct{}

func (util *Utils) GenerateUUID() (uuid string) {

    b := make([]byte, 16)
    _, err := rand.Read(b)
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }

    uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

    return
}

func (util *Utils) CheckErr(err error) {
    if err != nil {
        panic(err)
    }
}



