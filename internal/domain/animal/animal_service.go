package animal

import (
	"errors"
	"fmt"
	"github.com/vaberof/it-planet-task/internal/domain/vstlocation"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
	"time"
)

type AnimalService struct {
	animalStorage      AnimalStorage
	animalTypeStorage  AnimalTypeStorage
	accountStorage     AccountStorage
	locationStorage    LocationStorage
	vstLocationService VisitedLocationService
}

func NewAnimalService(
	animalStorage AnimalStorage,
	animalTypeStorage AnimalTypeStorage,
	accountStorage AccountStorage,
	locationStorage LocationStorage,
	visitedLocationStorage VisitedLocationService) *AnimalService {

	return &AnimalService{
		animalStorage:      animalStorage,
		animalTypeStorage:  animalTypeStorage,
		accountStorage:     accountStorage,
		locationStorage:    locationStorage,
		vstLocationService: visitedLocationStorage,
	}
}

func (s *AnimalService) CreateAnimal(
	chipperId int32,
	chippingLocationId int64,
	animalTypes []int64,
	weight float32,
	length float32,
	height float32,
	gender string) (*Animal, error) {
	return s.createAnimalImpl(chipperId, chippingLocationId, animalTypes, weight, length, height, gender)
}

func (s *AnimalService) GetAnimalById(id int64) (*Animal, error) {
	return s.animalStorage.GetAnimalById(id)
}

func (s *AnimalService) SearchAnimals(
	startDateTime *time.Time,
	endDateTime *time.Time,
	chipperId int32,
	chippingLocationId int64,
	lifeStatus string,
	gender string,
	from int32,
	size int32) ([]*Animal, error) {
	return s.searchAnimalsImpl(startDateTime,
		endDateTime,
		chipperId,
		chippingLocationId,
		lifeStatus,
		gender,
		from,
		size)
}

func (s *AnimalService) UpdateAnimal(
	animalId int64,
	weight float32,
	length float32,
	height float32,
	gender string,
	lifeStatus string,
	chipperId int32,
	chippingLocationId int64) (*Animal, error) {
	return s.updateAnimalImpl(animalId, weight, length, height, gender, lifeStatus, chipperId, chippingLocationId)
}

func (s *AnimalService) DeleteAnimal(animalId int64) error {
	return s.deleteAnimalImpl(animalId)
}

func (s *AnimalService) AddAnimalsType(animalId, animalTypeId int64) (*Animal, error) {
	return s.addAnimalsTypeImpl(animalId, animalTypeId)
}

func (s *AnimalService) UpdateAnimalsType(animalId, oldTypeId, newTypeId int64) (*Animal, error) {
	return s.updateAnimalsTypeImpl(animalId, oldTypeId, newTypeId)
}

func (s *AnimalService) DeleteAnimalsType(animalId, animalTypeId int64) (*Animal, error) {
	return s.deleteAnimalsTypeImpl(animalId, animalTypeId)
}

func (s *AnimalService) GetAnimalsVisitedLocations(
	animalId int64,
	startDateTime *time.Time,
	endDateTime *time.Time,
	from int32,
	size int32) ([]*vstlocation.VisitedLocation, error) {
	return s.getAnimalsVisitedLocationImpl(animalId, startDateTime, endDateTime, from, size)
}

func (s *AnimalService) AddAnimalsVisitedLocation(animalId, locationPointId int64, chipperId int32) (*vstlocation.VisitedLocation, error) {
	return s.addAnimalsVisitedLocationImpl(animalId, locationPointId, chipperId)
}

func (s *AnimalService) UpdateAnimalsVisitedLocation(animalId, visitedLocationPointId, locationPointId int64) (*vstlocation.VisitedLocation, error) {
	return s.updateAnimalsVisitedLocationImpl(animalId, visitedLocationPointId, locationPointId)
}

func (s *AnimalService) DeleteAnimalsVisitedLocation(animalId, visitedLocationId int64) error {
	return s.deleteAnimalsVisitedLocationImpl(animalId, visitedLocationId)
}

func (s *AnimalService) createAnimalImpl(
	chipperId int32,
	chippingLocationId int64,
	animalTypes []int64,
	weight float32,
	length float32,
	height float32,
	gender string) (*Animal, error) {

	if gender != GenderMale && gender != GenderFemale && gender != GenderOther {
		return nil, xerror.NewErrorWrapper(
			400,
			"incorrect gender",
			errors.New("invalid request body"))
	}

	if s.hasDuplicates(animalTypes) {
		return nil, xerror.NewErrorWrapper(
			409,
			"animal types contain duplicates",
			errors.New("duplicates in animal types"))
	}

	err := s.animalTypeStorage.FindAnimalTypes(animalTypes)
	if err != nil {
		return nil, err
	}

	err = s.accountStorage.FindAccountById(chipperId)
	if err != nil {
		return nil, err
	}

	err = s.locationStorage.FindLocationById(chippingLocationId)
	if err != nil {
		return nil, err
	}

	return s.animalStorage.CreateAnimal(chipperId, chippingLocationId, animalTypes, weight, length, height, gender)
}

func (s *AnimalService) searchAnimalsImpl(
	startDateTime *time.Time,
	endDateTime *time.Time,
	chipperId int32,
	chippingLocationId int64,
	lifeStatus string,
	gender string,
	from int32,
	size int32) ([]*Animal, error) {

	if lifeStatus != "" && lifeStatus != LifeStatusAlive && lifeStatus != LifeStatusDead {
		return nil, xerror.NewErrorWrapper(
			400,
			"incorrect life status",
			errors.New("incorrect life status"))
	}

	if gender != "" && gender != GenderMale && gender != GenderFemale && gender != GenderOther {
		return nil, xerror.NewErrorWrapper(
			400,
			"incorrect gender",
			errors.New("incorrect gender"))
	}

	return s.animalStorage.SearchAnimals(
		startDateTime,
		endDateTime,
		chipperId,
		chippingLocationId,
		lifeStatus,
		gender,
		from,
		size)
}

func (s *AnimalService) updateAnimalImpl(
	animalId int64,
	weight float32,
	length float32,
	height float32,
	gender string,
	lifeStatus string,
	chipperId int32,
	chippingLocationId int64) (*Animal, error) {

	if lifeStatus != LifeStatusAlive && lifeStatus != LifeStatusDead {
		return nil, xerror.NewErrorWrapper(
			400,
			"incorrect life status",
			errors.New("incorrect life status"))
	}

	if gender != GenderMale && gender != GenderFemale && gender != GenderOther {
		return nil, xerror.NewErrorWrapper(
			400,
			"incorrect gender",
			errors.New("incorrect gender"))
	}

	animal, err := s.animalStorage.GetAnimalById(animalId)
	if err != nil {
		return nil, err
	}

	if animal.LifeStatus == LifeStatusDead && lifeStatus == LifeStatusAlive {
		return nil, xerror.NewErrorWrapper(
			400,
			"incorrect life status",
			errors.New("incorrect life status"))
	}

	if len(animal.VisitedLocations) > 0 && animal.VisitedLocations[0].LocationPointId == chippingLocationId {
		return nil, xerror.NewErrorWrapper(
			400,
			"new chipping location id matches with first visited location id",
			errors.New("invalid chipping location id"))
	}

	err = s.accountStorage.FindAccountById(chipperId)
	if err != nil {
		return nil, err
	}

	err = s.locationStorage.FindLocationById(chippingLocationId)
	if err != nil {
		return nil, err
	}

	return s.animalStorage.UpdateAnimal(
		animalId,
		weight,
		length,
		height,
		gender,
		lifeStatus,
		chipperId,
		chippingLocationId)
}

func (s *AnimalService) deleteAnimalImpl(animalId int64) error {
	domainAnimal, err := s.animalStorage.GetAnimalById(animalId)
	if err != nil {
		return err
	}

	if len(domainAnimal.VisitedLocations) > 0 {
		return xerror.NewErrorWrapper(
			400,
			"animal has visited locations",
			errors.New("not empty visited locations"))
	}

	return s.animalStorage.DeleteAnimal(animalId)
}

func (s *AnimalService) addAnimalsTypeImpl(animalId, typeId int64) (*Animal, error) {
	err := s.animalTypeStorage.FindAnimalTypeById(typeId)
	if err != nil {
		return nil, err
	}

	err = s.animalStorage.FindAnimalById(animalId)
	if err != nil {
		return nil, err
	}

	err = s.animalTypeStorage.IsAnimalTypeAssociatedWithAnimal(typeId)
	if err != nil {
		return nil, err
	}

	return s.animalStorage.AddAnimalsType(animalId, typeId)
}

func (s *AnimalService) updateAnimalsTypeImpl(animalId, oldTypeId, newTypeId int64) (*Animal, error) {
	domainAnimal, err := s.animalStorage.GetAnimalById(animalId)
	if err != nil {
		return nil, err
	}

	err = s.animalTypeStorage.FindAnimalTypeById(oldTypeId)
	if err != nil {
		return nil, err
	}

	err = s.animalTypeStorage.FindAnimalTypeById(newTypeId)
	if err != nil {
		return nil, err
	}

	if !s.hasType(oldTypeId, domainAnimal.AnimalTypes) {
		return nil, xerror.NewErrorWrapper(
			404,
			fmt.Sprintf("animal with %d id has no type with %d type", domainAnimal.Id, oldTypeId),
			errors.New("not found animal type with given id among animal`s type"))
	}

	if s.hasType(oldTypeId, domainAnimal.AnimalTypes) && s.hasType(newTypeId, domainAnimal.AnimalTypes) {
		return nil, xerror.NewErrorWrapper(
			409,
			fmt.Sprintf("animal with %d id already has types %d and %d", domainAnimal.Id, oldTypeId, newTypeId),
			errors.New("animal already has given types"))
	}

	err = s.animalTypeStorage.IsAnimalTypeAssociatedWithAnimal(newTypeId)
	if err != nil {
		return nil, err
	}

	return s.animalStorage.UpdateAnimalsType(animalId, oldTypeId, newTypeId)
}

func (s *AnimalService) deleteAnimalsTypeImpl(animalId, typeId int64) (*Animal, error) {
	domainAnimal, err := s.animalStorage.GetAnimalById(animalId)
	if err != nil {
		return nil, err
	}

	if len(domainAnimal.AnimalTypes) == 1 && domainAnimal.AnimalTypes[0] == typeId {
		return nil, xerror.NewErrorWrapper(
			400,
			fmt.Sprintf("animal has only %d type", typeId),
			errors.New("animal has only one type"))
	}

	err = s.animalStorage.FindAnimalById(animalId)
	if err != nil {
		return nil, err
	}

	err = s.animalTypeStorage.FindAnimalTypeById(typeId)
	if err != nil {
		return nil, err
	}

	if !s.hasType(typeId, domainAnimal.AnimalTypes) {
		return nil, xerror.NewErrorWrapper(
			404,
			fmt.Sprintf("animal with %d id has no type with %d type", domainAnimal.Id, typeId),
			errors.New("not found animal type with given id among animal`s type"))
	}

	return s.animalStorage.DeleteAnimalsType(animalId, typeId)
}

func (s *AnimalService) getAnimalsVisitedLocationImpl(
	animalId int64,
	startDateTime *time.Time,
	endDateTime *time.Time,
	from int32,
	size int32) ([]*vstlocation.VisitedLocation, error) {

	err := s.animalStorage.FindAnimalById(animalId)
	if err != nil {
		return nil, err
	}

	return s.animalStorage.GetAnimalsVisitedLocations(animalId, startDateTime, endDateTime, from, size)
}

func (s *AnimalService) addAnimalsVisitedLocationImpl(animalId, locationId int64, chipperId int32) (*vstlocation.VisitedLocation, error) {
	domainAnimal, err := s.animalStorage.GetAnimalById(animalId)
	if err != nil {
		return nil, err
	}

	if domainAnimal.LifeStatus == LifeStatusDead {
		return nil, xerror.NewErrorWrapper(
			400, "animal has 'DEAD' life status",
			errors.New("animal has 'DEAD' life status"))
	}

	if len(domainAnimal.VisitedLocations) == 0 && domainAnimal.ChippingLocationId == locationId {
		return nil, xerror.NewErrorWrapper(
			400,
			"trying to add location that equal to chipping location",
			errors.New("cannot add location that equal to chipping location"))
	}

	if len(domainAnimal.VisitedLocations) > 0 &&
		domainAnimal.VisitedLocations[len(domainAnimal.VisitedLocations)-1].LocationPointId == locationId {
		return nil, xerror.NewErrorWrapper(
			400,
			"trying to add location that equal to current location",
			errors.New("cannot add location that equal to current location"))
	}

	visitedLocation, err := s.vstLocationService.CreateVisitedLocation(animalId, locationId, chipperId)
	if err != nil {
		return nil, err
	}

	return visitedLocation, nil
}

func (s *AnimalService) updateAnimalsVisitedLocationImpl(animalId, visitedLocationId, locationId int64) (*vstlocation.VisitedLocation, error) {
	domainAnimal, err := s.animalStorage.GetAnimalById(animalId)
	if err != nil {
		return nil, err
	}

	visitedLocationIndex := s.getVisitedLocationIndex(visitedLocationId, domainAnimal.VisitedLocations)
	if visitedLocationIndex == -1 {
		return nil, xerror.NewErrorWrapper(
			404,
			"animal does not have visited location with given id",
			errors.New("not found visited location among animal`s locations"))
	}

	err = s.locationStorage.FindLocationById(locationId)
	if err != nil {
		return nil, err
	}

	if visitedLocationIndex == 0 && locationId == domainAnimal.ChippingLocationId {
		return nil, xerror.NewErrorWrapper(
			400,
			"trying to update first visited location on chipping location",
			errors.New("cannot update first visited location to chipping location"))
	}

	if domainAnimal.VisitedLocations[visitedLocationIndex].LocationPointId == locationId {
		return nil, xerror.NewErrorWrapper(
			400,
			"trying to update location to the same location",
			errors.New("cannot update location to the same location"))
	}

	if s.matchWithPreviousLocation(locationId, visitedLocationIndex, domainAnimal.VisitedLocations) ||
		s.matchWithNextLocation(locationId, visitedLocationIndex, domainAnimal.VisitedLocations) {
		return nil, xerror.NewErrorWrapper(
			400,
			"previous or/and next location matches with given location",
			errors.New("cannot update previous or/and next matching location with given location"))
	}

	return s.animalStorage.UpdateAnimalsVisitedLocation(visitedLocationId, locationId)
}

func (s *AnimalService) deleteAnimalsVisitedLocationImpl(animalId, visitedLocationId int64) error {
	domainAnimal, err := s.animalStorage.GetAnimalById(animalId)
	if err != nil {
		return err
	}

	visitedLocationIndex := s.getVisitedLocationIndex(visitedLocationId, domainAnimal.VisitedLocations)
	if visitedLocationIndex == -1 {
		return xerror.NewErrorWrapper(
			404,
			"animal does not have visited location with given id",
			errors.New("not found visited location among animal`s locations"))
	}

	return s.animalStorage.DeleteAnimalsVisitedLocation(animalId, visitedLocationIndex)
}

func (s *AnimalService) hasDuplicates(animalTypes []int64) bool {
	visited := make(map[int64]bool)

	for _, animalType := range animalTypes {
		if visited[animalType] {
			return true
		}
		visited[animalType] = true
	}
	return false
}

func (s *AnimalService) hasType(animalType int64, animalTypes []int64) bool {
	for _, animType := range animalTypes {
		if animType == animalType {
			return true
		}
	}
	return false
}

func (s *AnimalService) getVisitedLocationIndex(visitedLocationId int64, visitedLocations []*vstlocation.VisitedLocation) int {
	for i := 0; i < len(visitedLocations); i++ {
		if visitedLocations[i].Id == visitedLocationId {
			return i
		}
	}
	return -1
}

func (s *AnimalService) matchWithPreviousLocation(locationId int64, index int, visitedLocations []*vstlocation.VisitedLocation) bool {
	return index-1 >= 0 && visitedLocations[index-1].LocationPointId == locationId
}

func (s *AnimalService) matchWithNextLocation(locationId int64, index int, visitedLocations []*vstlocation.VisitedLocation) bool {
	return index+1 < len(visitedLocations) && visitedLocations[index+1].LocationPointId == locationId
}
