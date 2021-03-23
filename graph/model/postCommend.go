package model

//PostCommend Post Commend
type PostCommend struct {
	ID        int    `gorm:"type:int;AUTO_INCREMENT;not null"`
	Body      string `gorm:"type:varchar(200);not null"`
	CreatedAt string `gorm:"type:timestamp;not null"`
	UserID    int    `gorm:"type:int;not null"`
	PostID    int    `gorm:"type:int;not null"`
}
