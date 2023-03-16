package vstlocation

type VisitedLocationService struct {
	visitedLocationStorage VisitedLocationStorage
	locationStorage        LocationStorage
}

func NewVisitedLocationService(
	visitedLocationStorage VisitedLocationStorage,
	locationStorage LocationStorage) *VisitedLocationService {

	return &VisitedLocationService{
		visitedLocationStorage: visitedLocationStorage,
		locationStorage:        locationStorage,
	}
}

func (s *VisitedLocationService) CreateVisitedLocation(animalId, locationPointId int64, chipperId int32) (*VisitedLocation, error) {
	return s.createVisitedLocationImpl(animalId, locationPointId, chipperId)
}

func (s *VisitedLocationService) createVisitedLocationImpl(animalId, locationPointId int64, chipperId int32) (*VisitedLocation, error) {
	err := s.locationStorage.FindLocationById(locationPointId)
	if err != nil {
		return nil, err
	}

	return s.visitedLocationStorage.CreateVisitedLocation(animalId, locationPointId, chipperId)
}
