package models

type User struct {
	ID         int    `gorm:"index;type:int" json:"id"`
	Username   string `gorm:"type:text" json:"username"`
	Password   string `gorm:"type:text" json:"password"`
	FirstName  string `gorm:"type:text" json:"first_name"`
	LastName   string `gorm:"type:text" json:"last_name"`
	Email      string `gorm:"type:text" json:"email"`
	Phone      string `gorm:"type:text" json:"phone"`
	CategoryID int    `gorm:"type:int" json:"category_id"`
	Created    string `gorm:"type:text" json:"created"`
	Updated    string `gorm:"type:text" json:"updated"`
	Active     int    `gorm:"index;type:int" json:"active"`
}

func (s User) TableName() string { return "users" }

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
