package types

type Activity struct {
	Rb        int     `gorm:"column:rb;primaryKey;autoIncrement"`
	ProgramId int     `gorm:"column:programId;ForeignKey:ProgramId"`
	Program   Program `gorm:"foreignKey:ProgramId" json:"-"`

	//food
	Naziv        string `gorm:"<-:false;column:naziv"`
	BrojKalorija int    `gorm:"<-:false;column:brojKalorija"`

	//gym
	BrojSerija int      `gorm:"<-:false;column:brojSerija"`
	Kilaza     float64  `gorm:"<-:false;column:kilaza"`
	VezbaId    int      `gorm:"<-:false;column:vezbaId"`
	Vezba      Exercise `gorm:"-"`
	//supplement
	//Naziv     string `gorm:"column:naziv"`
	Kolicina float64 `gorm:"<-:false;column:kolicina"`

	//string helper
	TipAktivnosti string `gorm:"column:activityType" json:"activityType"`
}
