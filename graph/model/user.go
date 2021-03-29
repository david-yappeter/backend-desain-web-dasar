package model

//User User
type User struct {
	ID        int     `gorm:"type:int;AUTO_INCREMENT;not null"`
	Name      string  `gorm:"type:varchar(100);not null"`
	Password  string  `gorm:"type:varchar(150);not null"`
	Email     string  `gorm:"type:varchar(100);not null"`
	Avatar    *string `gorm:"type:varchar(150);null;default:null"`
	CreatedAt string  `gorm:"type:timestamp;not null"`
	UpdatedAt *string `gorm:"type:timestamp;null;default:NULL"`
}
