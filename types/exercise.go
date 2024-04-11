package types

type Exercise struct {
	VezbaId int    `gorm:"column:vezbaId"`
	Naziv   string `gorm:"column:naziv"`
}

func (Exercise) TableName() string {
	return "vezba"
}
