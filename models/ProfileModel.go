package models

type Profile struct {
	ID       int64  `json:"-"`
	Username string `json:"username"`
}
