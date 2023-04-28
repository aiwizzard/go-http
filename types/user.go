package types

type User struct {
    ID int `json:"id"`
    name string `json:"name"`
}

func VBalidateUser(u *User) bool { return true }
