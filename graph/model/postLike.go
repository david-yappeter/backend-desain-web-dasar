package model

//PostLike Post Like
type PostLike struct {
	ID        int    `gorm:"type:int;AUTO_INCREMENT;not null"`
	CreatedAt string `gorm:"type:timestamp;not null"`
	UserID    int    `gorm:"type:int;not null"`
	PostID    int    `gorm:"type:int;not null"`
}
