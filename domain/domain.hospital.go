package domain

type Hospital struct {
	ID           uint   `gorm:"primarykey" json:"id"`
	NameCity     string `json:"name_city"`
	HospitalName string `json:"hospital_name"`
	Contact      string `json:"contact"`
	Address      string `json:"address"`
	Link         string `json:"link"`
}
