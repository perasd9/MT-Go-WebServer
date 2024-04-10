package interfaces

import "github/perasd9/MTWebServer/types"

type ExerciseUsecase interface {
	GetAll() []types.Exercise
}
