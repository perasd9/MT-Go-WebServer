package usecases

import (
	"github/perasd9/MTWebServer/infrastructure/repos/interfaces"
	"github/perasd9/MTWebServer/types"
	usecase "github/perasd9/MTWebServer/usecases/interfaces"
)

type memberUsecase struct {
	memberRepository interfaces.MemberRepository
}

func NewMemberUsecase(repo interfaces.MemberRepository) usecase.MemberUsecase {
	return &memberUsecase{
		memberRepository: repo,
	}
}

func (p *memberUsecase) Login(member types.Member) types.Member {
	return p.memberRepository.Login(member)
}

func (p *memberUsecase) Add(member types.Member) {

	p.memberRepository.Add(member)
}
