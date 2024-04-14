package repos

import (
	database "github/perasd9/MTWebServer/infrastructure"
	"github/perasd9/MTWebServer/infrastructure/repos/interfaces"
	"github/perasd9/MTWebServer/types"
)

type programRepository struct {
	db *database.MysqlDb
}

func NewProgramRepository(db *database.MysqlDb) interfaces.ProgramRepository {
	return &programRepository{
		db: db,
	}
}

// TO FOLLOW FULLY CLEAN ARCHITECTURE AND SOLID PRIN. RIGHT WAY IS TO LIFT UP THIS WHOLE LOGIC IN USECASE
// AND JUST CALL REPOSITORY YOU NEED FOR LOGIC BUT NOT IN REPOSITORY WORK WITH 2 TABLES
// THIS CRASHING SINGLE RESPONSIBILITY IN SOLID PRINCIPLES
func (p *programRepository) Add(param types.Program) {
	// var tip_programa []types.Program

	// p.db.Db.Table("tip_programa").Find(&tip_programa)

	tx := p.db.Db.Begin()

	result := tx.Table("program").Create(&param)

	if result.Error != nil {
		print((result.Error))
	}
	for _, value := range param.ListaAktivnosti {

		value.ProgramId = param.ProgramId
		tx.Table("aktivnost").Create(&value)

		switch value.TipAktivnosti {
		case "Teretana":
			act := &types.Gym{ProgramId: param.ProgramId, Rb: value.Rb, BrojSerija: value.BrojSerija, Kilaza: value.Kilaza, VezbaId: value.VezbaId, TipAktivnosti: value.TipAktivnosti}
			result = tx.Table(act.TipAktivnosti).Create(act)
		case "Hrana":
			act := &types.Food{ProgramId: param.ProgramId, Rb: value.Rb, Naziv: value.Naziv, BrojKalorija: value.BrojKalorija, TipAktivnosti: value.TipAktivnosti}
			result = tx.Table(act.TipAktivnosti).Create(act)
		case "Suplement":
			act := &types.Supplement{ProgramId: param.ProgramId, Rb: value.Rb, Naziv: value.Naziv, Kolicina: value.Kolicina, TipAktivnosti: value.TipAktivnosti}
			result = tx.Table(act.TipAktivnosti).Create(act)
		default:
			tx.Rollback()
			return
		}

		if result.Error != nil {
			tx.Rollback()
			return
		}
	}

	tx.Commit()
}

func (p *programRepository) GetAll(datum string) []types.Program {
	var programs []types.Program

	// p.db.Db.Model(types.Program{}).Select("program.*, clan.*").Where("datum = ?", datum).
	// 	Joins("left join clan on program.clanId = clan.clanId").Preload("Clan").Find(&programs)

	p.db.Db.Table("program").Where("datum = ? and public = ?", datum, true).Find(&programs)

	for index, value := range programs {
		var clan types.Member
		p.db.Db.Table("clan").Where("clanId = ?", value.ClanId).Find(&clan)

		programs[index].Clan = clan
	}

	return programs
}

func (p *programRepository) GetAllPrivatePrograms(param types.Program) []types.Program {
	var programs []types.Program

	p.db.Db.Table("program").Where("datum = ? and clanId = ?", param.Datum.String(), param.ClanId).Find(&programs)

	return programs
}
