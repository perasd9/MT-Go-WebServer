package types

type Supplement struct {
	Rb            int `gorm:"column:rb"`
	ProgramId     int `gorm:"column:programId"`
	Program       Program
	Naziv         string   `gorm:"column:naziv"`
	Kolicina      float64  `gorm:"column:kolicina"`
	Activity      Activity `gorm:"-"`
	TipAktivnosti string   `gorm:"-" json:"activityType"`
}
