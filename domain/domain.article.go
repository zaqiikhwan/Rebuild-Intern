package domain

type Article struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Image    string `json:"image"`
	Category string `json:"category"`
}
