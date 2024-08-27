package models

type Blog struct {
	Id        uint   `json:"id`
	Title string `json:title`
	Body  string `json:body`
	Image     string `json:image`
	UserID  string `json:user_id`
	User User `json:"user";gorm:foreignkey:UserID`
}