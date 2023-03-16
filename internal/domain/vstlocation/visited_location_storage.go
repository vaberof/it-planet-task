package vstlocation

type VisitedLocationStorage interface {
	CreateVisitedLocation(animalId, locationPointId int64, chipperId int32) (*VisitedLocation, error)
}
