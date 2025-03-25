package models

type User struct {
	ID         int    `gorm:"index;type:int" json:"id"`
	UserName   string `gorm:"type:text" json:"username"`
	Password   string `gorm:"type:text" json:"password"`
	FirstName  string `gorm:"type:text" json:"first_name"`
	LastName   string `gorm:"type:text" json:"last_name"`
	Email      string `gorm:"type:text" json:"email"`
	Phone      string `gorm:"type:text" json:"phone"`
	CategoryID int    `gorm:"type:int" json:"categoryId"`
	Created    int    `gorm:"type:text" json:"created"`
	Updated    int    `gorm:"type:text" json:"updated"`
	Active     int    `gorm:"index;type:int" json:"active"`
}

func (s User) TableName() string { return "users" }

type LoginParams struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
