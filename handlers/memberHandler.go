package handlers

import (
	"bytes"
	"encoding/json"
	"github/perasd9/MTWebServer/types"
	"github/perasd9/MTWebServer/usecases/interfaces"
)

type memberHandler struct {
	memberUsecase interfaces.MemberUsecase
}

func NewMemberHandler(usecase interfaces.MemberUsecase) *memberHandler {
	return &memberHandler{
		memberUsecase: usecase,
	}
}

func (p *memberHandler) Login(param string) string {

	var member types.Member

	byted := []byte(param)

	byted = bytes.Trim(byted, "\x00")

	err := json.Unmarshal(byted, &member)

	if err != nil {
		return err.Error()
	}

	mem := p.memberUsecase.Login(member)

	if mem.ClanId != 0 {
		v, _ := json.Marshal(mem)
		return string(v)

	} else {
		return ""
	}

}

func (p *memberHandler) Add(param string) string {

	// types := p.memberUsecase.Add()

	// v, _ := json.Marshal(types)

	// return string(v)
	return ""
}
