package models

type Comment struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	UserId uint   `json:"userId" validate:"required"`
	Title  string `json:"title" validate:"required,min=3"`
	Body   string `json:"body" validate:"required,min=3"`
	User   User   `gorm:"foreignKey:UserId" json:"user"`
}

type DataGetComment struct {
	ID     uint   `json:"id"`
	UserId uint   `json:"userId" validate:"required"`
	Title  string `json:"title" validate:"required,min=3"`
	Body   string `json:"body" validate:"required,min=3"`
}
