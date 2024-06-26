package usecases

import (
	"github/perasd9/MTWebServer/infrastructure/repos/interfaces"
	"github/perasd9/MTWebServer/types"
	usecase "github/perasd9/MTWebServer/usecases/interfaces"
	"strconv"
)

type programUsecase struct {
	programRepository  interfaces.ProgramRepository
	activityRepository interfaces.ActivityRepository
}

func NewProgramUsecase(repo interfaces.ProgramRepository, repo1 interfaces.ActivityRepository) usecase.ProgramUsecase {
	return &programUsecase{
		programRepository:  repo,
		activityRepository: repo1,
	}
}

func (p *programUsecase) Add(program types.Program) {
	p.programRepository.Add(program)
}

func (p *programUsecase) GetAll(datum string) []types.Program {

	programs := p.programRepository.GetAll(datum)

	for i, value := range programs {
		var activities []types.Activity
		var concreteActivities []types.Activity

		activities = p.activityRepository.GetAllDistinct(strconv.Itoa(value.ProgramId))

		for _, activityValue := range activities {
			concreteActivities = p.activityRepository.GetAll(activityValue)

		}
		programs[i].ListaAktivnosti = concreteActivities
	}

	return programs
}

func (p *programUsecase) GetAllPrivatePrograms(param types.Program) []types.Program {

	programs := p.programRepository.GetAllPrivatePrograms(param)

	for i, value := range programs {
		var activities []types.Activity
		var concreteActivities []types.Activity

		activities = p.activityRepository.GetAllDistinct(strconv.Itoa(value.ProgramId))

		for _, activityValue := range activities {
			concreteActivities = p.activityRepository.GetAll(activityValue)

		}
		programs[i].ListaAktivnosti = concreteActivities
	}

	return programs
}

func (p *programUsecase) Delete(param int) {
	p.activityRepository.Delete(param, 0)
	p.programRepository.Delete(param)
}

func (p *programUsecase) Update(param types.Program) error {
	p.programRepository.Update(param)

	var result error

	for _, value := range param.ListaAktivnosti {
		if value.IsDeleted {
			p.activityRepository.Delete(value.ProgramId, value.Rb)
		}
		if value.IsAdded {
			value.ProgramId = param.ProgramId

			result = p.activityRepository.Add(value)
			if result != nil {
				p.programRepository.Rollback()
				return result
			}
		}
	}
	return result
}
