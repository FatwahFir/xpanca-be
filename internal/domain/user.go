package domain

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"size:64;uniqueIndex"`
	Password string `gorm:"size:255"`
	Role     string `gorm:"size:32"`
}
