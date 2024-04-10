package usecases

import (
	"github/perasd9/MTWebServer/infrastructure/repos/interfaces"
	"github/perasd9/MTWebServer/types"
	usecase "github/perasd9/MTWebServer/usecases/interfaces"
)

type exerciseUsecase struct {
	exerciseRepository interfaces.ExerciseRepository
}

func NewExerciseUsecase(repo interfaces.ExerciseRepository) usecase.ExerciseUsecase {
	return &exerciseUsecase{
		exerciseRepository: repo,
	}
}

func (p *exerciseUsecase) GetAll() []types.Exercise {
	return p.exerciseRepository.GetAll()
}
