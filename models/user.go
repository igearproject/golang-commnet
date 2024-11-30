package models

type User struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Name     string    `json:"name" validate:"required,min=3,max=50"`
	Email    string    `gorm:"unique" json:"email" validate:"required,email"`
	Password string    `json:"password" validate:"required,min=8"`
	Comments []Comment `gorm:"foreignKey:UserId" json:"comments"`
}
