package handlers

import (
	"encoding/json"
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

	v, _ := json.Marshal(types)

	return string(v)
}
