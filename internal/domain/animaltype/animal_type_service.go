package animaltype

import (
	"errors"
	"fmt"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
)

type AnimalTypeService struct {
	animalTypeStorage AnimalTypeStorage
}

func NewAnimalTypeService(animalTypeStorage AnimalTypeStorage) *AnimalTypeService {
	return &AnimalTypeService{
		animalTypeStorage: animalTypeStorage,
	}
}

func (s *AnimalTypeService) CreateAnimalType(chipperId int32, animalType string) (*AnimalType, error) {
	return s.createAnimalTypeImpl(chipperId, animalType)
}

func (s *AnimalTypeService) GetAnimalTypeById(typeId int64) (*AnimalType, error) {
	return s.getAnimalTypeByIdImpl(typeId)
}

func (s *AnimalTypeService) UpdateAnimalType(typeId int64, animalType string) (*AnimalType, error) {
	return s.updateAnimalTypeImpl(typeId, animalType)
}

func (s *AnimalTypeService) DeleteAnimalType(typeId int64) error {
	return s.deleteAnimalTypeImpl(typeId)
}

func (s *AnimalTypeService) createAnimalTypeImpl(chipperId int32, animalType string) (*AnimalType, error) {
	err := s.animalTypeStorage.FindAnimalType(animalType)
	if err == nil {
		return nil, xerror.NewErrorWrapper(409,
			"animal type with given type already exists",
			errors.New("animal type with given type already exists"))
	}

	return s.animalTypeStorage.CreateAnimalType(chipperId, animalType)
}

func (s *AnimalTypeService) getAnimalTypeByIdImpl(typeId int64) (*AnimalType, error) {
	return s.animalTypeStorage.GetAnimalTypeById(typeId)
}

func (s *AnimalTypeService) updateAnimalTypeImpl(typeId int64, animalType string) (*AnimalType, error) {
	return s.animalTypeStorage.UpdateAnimalType(typeId, animalType)
}

func (s *AnimalTypeService) deleteAnimalTypeImpl(typeId int64) error {
	err := s.animalTypeStorage.IsAnimalTypeAssociatedWithAnimal(typeId)
	if err != nil {
		return xerror.NewErrorWrapper(
			400,
			fmt.Sprintf("there are existing animals with %d id", typeId),
			errors.New("trying to delete associated animal type id"))
	}

	return s.animalTypeStorage.DeleteAnimalType(typeId)
}
