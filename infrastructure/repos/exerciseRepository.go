package repos

import (
	database "github/perasd9/MTWebServer/infrastructure"
	"github/perasd9/MTWebServer/infrastructure/repos/interfaces"
	"github/perasd9/MTWebServer/types"
)

type exerciseRepository struct {
	db *database.MysqlDb
}

func NewExerciseRepository(db *database.MysqlDb) interfaces.ExerciseRepository {
	return &exerciseRepository{
		db: db,
	}
}

func (p *exerciseRepository) GetAll() []types.Exercise {
	var vezbe []types.Exercise

	p.db.Db.Table("vezba").Find(&vezbe)

	return vezbe
}
