package types

type ProgramType struct {
	TipProgramaId int    `gorm:"column:tipProgramaId"`
	Naziv         string `gorm:"column:naziv"`
}
