package interfaces

import "github/perasd9/MTWebServer/types"

type ProgramUsecase interface {
	Add(types.Program)
	GetAll(string) []types.Program
	GetAllPrivatePrograms(datum string) []types.Program
}
