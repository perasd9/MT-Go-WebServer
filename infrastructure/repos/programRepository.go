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

	p.db.Db.Table("program").Where("datum = ?", datum).Find(&programs)

	return programs
}
