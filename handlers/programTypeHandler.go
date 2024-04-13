package handlers

import (
	"encoding/json"
	handlers "github/perasd9/MTWebServer/handlers/serverHandlers"
	"github/perasd9/MTWebServer/usecases/interfaces"
)

type programTypeHandler struct {
	programTypeUsecase interfaces.ProgramTypeUsecase
}

func NewProgramTypeHandler(usecase interfaces.ProgramTypeUsecase) *programTypeHandler {
	return &programTypeHandler{
		programTypeUsecase: usecase,
	}
}

func (p *programTypeHandler) GetAll(param string) string {

	types := p.programTypeUsecase.GetAll()

	v, err := json.MarshalIndent(types, "", "   ")

	if err != nil {
		return handlers.NewResponse().BadRequest(string(err.Error()))
	}

	return handlers.NewResponse().Ok(string(v))
}
