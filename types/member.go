package types

type Member struct {
	ClanId  int    `gorm:"column:clanId;primaryKey;autoIncrement"`
	Ime     string `gorm:"column:ime"`
	Prezime string `gorm:"column:prezime"`
	Email   string `gorm:"column:email"`
	Lozinka string `gorm:"column:lozinka"`
}
