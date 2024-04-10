package types

type Food struct {
	Rb            int      `gorm:"column:rb"`
	ProgramId     int      `gorm:"column:programId"`
	Program       Program  `gorm:"-"`
	Naziv         string   `gorm:"column:naziv"`
	BrojKalorija  int      `gorm:"column:brojKalorija"`
	Activity      Activity `gorm:"-"`
	TipAktivnosti string   `gorm:"-" json:"activityType"`
}
