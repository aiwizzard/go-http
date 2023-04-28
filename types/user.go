package types

type User struct {
    ID int `json:"id"`
    Name string `json:"name"`
}

func VBalidateUser(u *User) bool { return true }
