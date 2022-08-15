package models

type Post struct {
	ID      int64  `json:"-";sql:"type:string REFERENCES users(username)"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
