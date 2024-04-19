package usecases

import (
	"github/perasd9/MTWebServer/infrastructure/repos/interfaces"
	"github/perasd9/MTWebServer/types"
	usecase "github/perasd9/MTWebServer/usecases/interfaces"

	"golang.org/x/crypto/bcrypt"
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

	mem := p.memberRepository.Login(member)

	if CheckPasswordHash(member.Lozinka, mem.Lozinka) {
		return mem
	}
	//clanId = 0
	return member
}

func (p *memberUsecase) Add(member types.Member) {
	member.Lozinka, _ = HashPassword(member.Lozinka)

	p.memberRepository.Add(member)
}

// hashing password func
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 6)
	return string(bytes), err
}

// checking hash password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
