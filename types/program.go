package types

import "time"

type Program struct {
	ProgramId       int        `gorm:"column:programId;primaryKey;autoIncrement" json:"programId"`
	Naziv           string     `gorm:"column:naziv" json:"naziv"`
	ListaAktivnosti []Activity `gorm:"-" json:"listaAktivnosti"`
	Datum           time.Time  `gorm:"column:datum" json:"datum"`
	Public          bool       `gorm:"column:public" json:"public"`
	ClanId          int        `gorm:"column:clanId; foreignKey:clanId" json:"clanId"`
	Clan            Member     `gorm:"<-:false; foreignKey:clanId"`
	// TipProgramaId string `gorm:"column:tipProgramaId"`
	// TipPrograma   ProgramType
}

func (Program) TableName() string {
	return "program"
}
