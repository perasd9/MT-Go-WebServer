package repos

import (
	database "github/perasd9/MTWebServer/infrastructure"
	"github/perasd9/MTWebServer/infrastructure/repos/interfaces"
	"github/perasd9/MTWebServer/types"
)

type memberRepository struct {
	db *database.MysqlDb
}

func NewMemberRepository(db *database.MysqlDb) interfaces.MemberRepository {
	return &memberRepository{
		db: db,
	}
}

func (p *memberRepository) Add(member types.Member) {
	p.db.Db.Table("clan").Create(&member)
}

func (p *memberRepository) Login(member types.Member) types.Member {
	var mem types.Member

	p.db.Db.Table("clan").Where("email = ?", member.Email).Find(&mem)

	return mem
}
