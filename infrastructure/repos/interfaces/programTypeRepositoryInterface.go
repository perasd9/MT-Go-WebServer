package interfaces

import "github/perasd9/MTWebServer/types"

type ProgramTypeRepository interface {
	GetAll() []types.ProgramType
}
