package interfaces

import "github/perasd9/MTWebServer/types"

type ActivityRepository interface {
	GetAll(param types.Activity) []types.Activity
	GetAllDistinct(param string) []types.Activity
}
