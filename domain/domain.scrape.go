package domain

type Scrape struct {
	ID           	 uint   `gorm:"primarykey" json:"id"`
	Location     	 string `gorm:"location" json:"location"`
	Name         	 string `gorm:"name" json:"name"`
	Address      	 string `gorm:"address" json:"address"`
	Phone_Number 	 string `gorm:"phone_number" json:"phone_number"`
	Link_Google_Maps string `gorm:"link_google_maps" json:"link_google_maps"`
}