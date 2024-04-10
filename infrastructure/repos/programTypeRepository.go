package repos

import (
	database "github/perasd9/MTWebServer/infrastructure"
	"github/perasd9/MTWebServer/infrastructure/repos/interfaces"
	"github/perasd9/MTWebServer/types"
)

type programTypeRepository struct {
	db *database.MysqlDb
}

func NewProgramTypeRepository(db *database.MysqlDb) interfaces.ProgramTypeRepository {
	return &programTypeRepository{
		db: db,
	}
}

func (p *programTypeRepository) GetAll() []types.ProgramType {
	var tip_programa []types.ProgramType

	p.db.Db.Table("tip_programa").Find(&tip_programa)

	return tip_programa
}
