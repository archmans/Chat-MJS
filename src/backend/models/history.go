package models

type History struct {
	ID         int64  `gorm:"primaryKey" json:"id"`
	Pertanyaan string `gorm:"type:text" json:"pertanyaan"`
	Jawaban    string `gorm:"type:text" json:"jawaban"`
}
