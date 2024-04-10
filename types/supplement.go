package types

type Supplement struct {
	Rb            int `gorm:"column:rb"`
	ProgramId     int `gorm:"column:programId"`
	Program       Program
	Naziv         string   `gorm:"column:naziv"`
	Kolicina      string   `gorm:"column:kolicina"`
	Activity      Activity `gorm:"-"`
	TipAktivnosti string   `json:"activityType"`
}
