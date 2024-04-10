package interfaces

import "github/perasd9/MTWebServer/types"

type ProgramTypeUsecase interface {
	GetAll() []types.ProgramType
}
