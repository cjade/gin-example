package models

type Books struct {
	BookId       uint `gorm:"primaryKey"`
	TypeId       uint
	BookName     string `gorm:"size:50"`
	Author       string `gorm:"size:50"`
	Intro        string
	CoverPicture string
	Model
}
