package handlers

import (
	"bytes"
	"encoding/json"
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
		return err.Error()
	}
	// _, _ := json.Marshal(program)

	p.programUsecase.Add(program)

	return "ok"
}

func (p *programHandler) GetAll(param string) string {

	var datum string
	for i := 0; i < len(param); i++ {
		if param[i] == '2' && param[i+1] == '0' {
			for j := i; j < len(param); j++ {
				if param[j] != '"' && param[j] != '}' {
					datum += string(param[j])
				}
			}
			break
		}
	}

	programs := p.programUsecase.GetAll(datum)

	jsonPrograms, _ := json.Marshal(programs)

	return string(jsonPrograms)
}
