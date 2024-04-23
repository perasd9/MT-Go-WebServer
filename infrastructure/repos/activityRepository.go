package repos

import (
	database "github/perasd9/MTWebServer/infrastructure"
	"github/perasd9/MTWebServer/infrastructure/repos/interfaces"
	"github/perasd9/MTWebServer/types"
)

type activityRepository struct {
	db *database.MysqlDb
}

func NewActivityRepository(db *database.MysqlDb) interfaces.ActivityRepository {
	return &activityRepository{
		db: db,
	}
}

func (a *activityRepository) GetAll(param types.Activity) []types.Activity {
	var activities []types.Activity

	if param.TipAktivnosti == "Teretana" {
		a.db.Db.Table(param.TipAktivnosti).Select("teretana.*, vezba.*").
			Where("programId = ?", param.ProgramId).Joins("join vezba on teretana.vezbaId = vezba.vezbaId").
			Scan(&activities)

		return activities
	}
	a.db.Db.Table(param.TipAktivnosti).Where("programId = ?", param.ProgramId).Find(&activities)

	return activities
}

func (a *activityRepository) GetAllDistinct(param string) []types.Activity {
	var activities []types.Activity

	a.db.Db.Table("aktivnost").Where("programId = ?", param).Distinct("programId, activityType").Find(&activities)

	return activities
}

func (a *activityRepository) Delete(param int, param2 int) {
	var activities []types.Activity

	a.db.Db.Table("aktivnost").Where("programId = ?", param).Distinct("programId, activityType").Find(&activities)
	if param2 != 0 {
		a.db.Db.Table(activities[0].TipAktivnosti).Where("programId = ? and rb = ?", param, param2).Delete(param)
		a.db.Db.Table("aktivnost").Where("programId = ? and rb = ?", param, param2).Delete(param)

	} else {
		a.db.Db.Table(activities[0].TipAktivnosti).Where("programId = ?", param).Delete(param)
		a.db.Db.Table("aktivnost").Where("programId = ?", param).Delete(param)
	}
}

func (a *activityRepository) Add(param types.Activity) error {

	result := a.db.Db.Table("aktivnost").Create(&param)

	if result.Error != nil {
		return result.Error
	}

	switch param.TipAktivnosti {
	case "Teretana":
		act := &types.Gym{ProgramId: param.ProgramId, Rb: param.Rb, BrojSerija: param.BrojSerija, Kilaza: param.Kilaza, VezbaId: param.VezbaId, TipAktivnosti: param.TipAktivnosti}
		result = a.db.Db.Table(act.TipAktivnosti).Create(act)
	case "Hrana":
		act := &types.Food{ProgramId: param.ProgramId, Rb: param.Rb, Naziv: param.Naziv, BrojKalorija: param.BrojKalorija, TipAktivnosti: param.TipAktivnosti}
		result = a.db.Db.Table(act.TipAktivnosti).Create(act)
	case "Suplement":
		act := &types.Supplement{ProgramId: param.ProgramId, Rb: param.Rb, Naziv: param.Naziv, Kolicina: param.Kolicina, TipAktivnosti: param.TipAktivnosti}
		result = a.db.Db.Table(act.TipAktivnosti).Create(act)
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}
