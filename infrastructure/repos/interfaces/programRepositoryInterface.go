package interfaces

import "github/perasd9/MTWebServer/types"

type ProgramRepository interface {
	Add(types.Program)
	GetAll(datum string) []types.Program
	GetAllPrivatePrograms(datum string) []types.Program
}
