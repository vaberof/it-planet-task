package animal

type AnimalTypeStorage interface {
	FindAnimalTypes(animalTypes []int64) error
	FindAnimalTypeById(typeId int64) error
	IsAnimalTypeAssociatedWithAnimal(animalTypeId int64) error
}
