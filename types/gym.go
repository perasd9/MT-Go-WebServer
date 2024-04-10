package types

type Gym struct {
	Rb            int `gorm:"column:rb"`
	ProgramId     int `gorm:"column:programId"`
	Program       Program
	BrojSerija    string  `gorm:"column:brojSerija"`
	Kilaza        float64 `gorm:"column:kilaza"`
	VezbaId       int     `gorm:"column:vezbaId"`
	Vezba         Exercise
	Activity      Activity `gorm:"-"`
	TipAktivnosti string   `json:"activityType"`
}
