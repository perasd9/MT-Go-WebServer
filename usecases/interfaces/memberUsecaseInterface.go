package interfaces

import "github/perasd9/MTWebServer/types"

type MemberUsecase interface {
	Login(member types.Member) types.Member
	Add(member types.Member)
}
