package interfaces

import "github/perasd9/MTWebServer/types"

type ProgramUsecase interface {
	Add(types.Program)
	GetAll(string) []types.Program
	GetAllPrivatePrograms(parm types.Program) []types.Program
	Delete(param int)
}
