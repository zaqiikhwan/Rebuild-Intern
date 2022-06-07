package domain

type User struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}
