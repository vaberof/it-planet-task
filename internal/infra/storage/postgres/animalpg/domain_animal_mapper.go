package animalpg

import (
	"github.com/vaberof/it-planet-task/internal/domain/animal"
	"github.com/vaberof/it-planet-task/internal/infra/storage/postgres/vstlocationpg"
)

func BuildDomainAnimal(postgresAnimal *Animal) *animal.Animal {
	return &animal.Animal{
		Id:                 postgresAnimal.Id,
		AnimalTypes:        postgresAnimal.AnimalTypes,
		Weight:             postgresAnimal.Weight,
		Length:             postgresAnimal.Length,
		Height:             postgresAnimal.Height,
		Gender:             postgresAnimal.Gender,
		LifeStatus:         postgresAnimal.LifeStatus,
		ChippingDateTime:   postgresAnimal.ChippingDateTime,
		ChipperId:          postgresAnimal.ChipperId,
		ChippingLocationId: postgresAnimal.ChippingLocationId,
		VisitedLocations:   vstlocationpg.BuildDomainVisitedLocations(postgresAnimal.VisitedLocations),
		DeathDateTime:      postgresAnimal.DeathDateTime,
	}
}

func BuildDomainAnimals(postgresAnimals []*Animal) []*animal.Animal {
	domainAnimals := make([]*animal.Animal, len(postgresAnimals))

	for i := 0; i < len(domainAnimals); i++ {
		domainAnimals[i] = BuildDomainAnimal(postgresAnimals[i])
	}

	return domainAnimals
}
