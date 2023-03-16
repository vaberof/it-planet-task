package location

import (
	"errors"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
)

type LocationService struct {
	locationStorage LocationStorage
}

func NewLocationService(locationStorage LocationStorage) *LocationService {
	return &LocationService{
		locationStorage: locationStorage,
	}
}

func (s *LocationService) CreateLocation(chipperId int32, latitude, longitude float64) (*Location, error) {
	return s.createLocationImpl(chipperId, latitude, longitude)
}

func (s *LocationService) GetLocation(pointId int64) (*Location, error) {
	return s.getLocationImpl(pointId)
}

func (s *LocationService) UpdateLocation(pointId int64, latitude, longitude float64) (*Location, error) {
	err := s.locationStorage.FindLocation(latitude, longitude)
	if err == nil {
		return nil, xerror.NewErrorWrapper(409,
			"animal type with given type already exists",
			errors.New("animal type with given type already exists"))
	}

	return s.updateLocationImpl(pointId, latitude, longitude)
}

func (s *LocationService) DeleteLocation(pointId int64) error {
	return s.deleteLocationImpl(pointId)
}

func (s *LocationService) createLocationImpl(chipperId int32, latitude, longitude float64) (*Location, error) {
	err := s.locationStorage.FindLocation(latitude, longitude)
	if err == nil {
		return nil, xerror.NewErrorWrapper(
			409,
			"location point with given latitude and longitude already exists",
			errors.New("location already exists"))
	}

	return s.locationStorage.CreateLocation(chipperId, latitude, longitude)
}

func (s *LocationService) getLocationImpl(pointId int64) (*Location, error) {
	return s.locationStorage.GetLocation(pointId)
}

func (s *LocationService) updateLocationImpl(pointId int64, latitude, longitude float64) (*Location, error) {
	return s.locationStorage.UpdateLocation(pointId, latitude, longitude)
}

func (s *LocationService) deleteLocationImpl(pointId int64) error {
	err := s.locationStorage.IsLocationAssociatedWithAnimal(pointId)
	if err != nil {
		return err
	}

	err = s.locationStorage.IsLocationAssociatedWithAnimalsVisitedLocation(pointId)
	if err != nil {
		return err
	}

	return s.locationStorage.DeleteLocation(pointId)
}
