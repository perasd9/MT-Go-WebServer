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
