package interfaces

import "github/perasd9/MTWebServer/types"

type ExerciseRepository interface {
	GetAll() []types.Exercise
}
