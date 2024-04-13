package handlers

import (
	"encoding/json"
	handlers "github/perasd9/MTWebServer/handlers/serverHandlers"
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

	v, err := json.MarshalIndent(types, "", "   ")

	if err != nil {
		return handlers.NewResponse().BadRequest(string(err.Error()))
	}

	return handlers.NewResponse().Ok(string(v))
}
