package usecases

import (
	"github/perasd9/MTWebServer/infrastructure/repos/interfaces"
	"github/perasd9/MTWebServer/types"
	usecase "github/perasd9/MTWebServer/usecases/interfaces"
)

// NOT USING IN BEGGINING
type programTypeUsecase struct {
	programTypeRepository interfaces.ProgramTypeRepository
}

func NewProgramTypeUsecase(repo interfaces.ProgramTypeRepository) usecase.ProgramTypeUsecase {
	return &programTypeUsecase{
		programTypeRepository: repo,
	}
}

func (p *programTypeUsecase) GetAll() []types.ProgramType {
	return p.programTypeRepository.GetAll()
}
