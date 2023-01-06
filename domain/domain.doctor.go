package domain

type Doctor struct {
	ID          uint    `gorm:"primarykey" json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Jadwal      string  `json:"jadwal"`
	LokasiKerja string  `json:"lokasi_kerja"`
	Picture     string  `json:"picture"`
	Meet        string  `json:"meet"`
	Price       float64 `json:"price"`
	Pengalaman  int     `json:"pengalaman"`
	Password    string  `json:"password"`
	Username    string  `json:"username"`
}