package interfaces

import (
	"github/perasd9/MTWebServer/types"
)

type ProgramRepository interface {
	Add(types.Program)
	GetAll(datum string) []types.Program
	GetAllPrivatePrograms(param types.Program) []types.Program
	Delete(param int)
	Update(param types.Program)
	BeginTransaction()
	Commit()
	Rollback()
}
