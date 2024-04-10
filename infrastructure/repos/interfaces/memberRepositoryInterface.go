package interfaces

import "github/perasd9/MTWebServer/types"

type MemberRepository interface {
	Login(member types.Member) types.Member
	Add(member types.Member)
}
