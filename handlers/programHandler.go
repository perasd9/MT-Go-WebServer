package handlers

import (
	"bytes"
	"encoding/json"
	handlers "github/perasd9/MTWebServer/handlers/serverHandlers"
	"github/perasd9/MTWebServer/types"
	"github/perasd9/MTWebServer/usecases/interfaces"
)

type programHandler struct {
	programUsecase interfaces.ProgramUsecase
}

func NewProgramHandler(usecase interfaces.ProgramUsecase) *programHandler {
	return &programHandler{
		programUsecase: usecase,
	}
}

func (p *programHandler) Add(param string) string {
	var program types.Program

	byted := []byte(param)

	byted = bytes.Trim(byted, "\x00")

	err := json.Unmarshal(byted, &program)

	if err != nil {
		return handlers.NewResponse().BadRequest(err.Error())
	}

	p.programUsecase.Add(program)

	return handlers.NewResponse().Created("")
}

func (p *programHandler) GetAll(param string) string {
	var program types.Program

	byted := []byte(param)

	byted = bytes.Trim(byted, "\x00")

	err := json.Unmarshal(byted, &program)

	if err != nil {
		return handlers.NewResponse().BadRequest(err.Error())
	}

	programs := p.programUsecase.GetAll(program.Datum.String())

	jsonPrograms, err := json.MarshalIndent(programs, "", "   ")

	if err != nil {
		return handlers.NewResponse().BadRequest(string(err.Error()))
	}

	return handlers.NewResponse().Ok(string(jsonPrograms))
}
