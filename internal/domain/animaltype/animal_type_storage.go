package animaltype

type AnimalTypeStorage interface {
	CreateAnimalType(chipperId int32, animalType string) (*AnimalType, error)
	GetAnimalTypeById(typeId int64) (*AnimalType, error)
	UpdateAnimalType(typeId int64, animalType string) (*AnimalType, error)
	DeleteAnimalType(typeId int64) error
	FindAnimalType(animalType string) error
	IsAnimalTypeAssociatedWithAnimal(animalTypeId int64) error
}
