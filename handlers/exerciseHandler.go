package handlers

import (
	"encoding/json"
	"github/perasd9/MTWebServer/usecases/interfaces"
)

type exerciseHandler struct {
	exerciseUsecase interfaces.ExerciseUsecase
}

func NewExerciseHandler(usecase interfaces.ExerciseUsecase) *exerciseHandler {
	return &exerciseHandler{
		exerciseUsecase: usecase,
	}
}

func (p *exerciseHandler) GetAll(param string) string {

	types := p.exerciseUsecase.GetAll()

	v, _ := json.Marshal(types)

	return string(v)
}