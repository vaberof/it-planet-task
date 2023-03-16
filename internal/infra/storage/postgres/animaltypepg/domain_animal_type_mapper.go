package animaltypepg

import (
	"github.com/vaberof/it-planet-task/internal/domain/animaltype"
)

func BuildDomainAnimalType(postgresAnimalType *AnimalType) *animaltype.AnimalType {
	return &animaltype.AnimalType{
		Id:   postgresAnimalType.Id,
		Type: postgresAnimalType.Type,
	}
}
