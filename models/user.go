package models

type User struct {
	ID       int    `json:"id" gorm:"primary_key:auto_increment"`
	Username string `json:"username" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"type:varchar(255)"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	Role     string `json:"role" gorm:"type:text"`
}

type UserRespon struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (UserRespon) TableName() string {
	return "users"
}
