package location

type LocationStorage interface {
	FindLocation(latitude, longitude float64) error
	CreateLocation(chipperId int32, latitude, longitude float64) (*Location, error)
	GetLocation(pointId int64) (*Location, error)
	UpdateLocation(pointId int64, latitude, longitude float64) (*Location, error)
	DeleteLocation(pointId int64) error
	IsLocationAssociatedWithAnimal(pointId int64) error
	IsLocationAssociatedWithAnimalsVisitedLocation(pointId int64) error
}
