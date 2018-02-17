package user


type User struct {

}

type LoginParams struct {
	User string `json:"user"`
	Password string `json:"password"`
	UserData bool `json:"userData"`
}