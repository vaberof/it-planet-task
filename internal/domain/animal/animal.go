package animal

import (
	"github.com/vaberof/it-planet-task/internal/domain/vstlocation"
	"time"
)

type Animal struct {
	Id                 int64
	AnimalTypes        []int64
	Weight             float32
	Length             float32
	Height             float32
	Gender             string
	LifeStatus         string
	ChippingDateTime   time.Time
	ChipperId          int32
	ChippingLocationId int64
	VisitedLocations   []*vstlocation.VisitedLocation
	DeathDateTime      *time.Time
}
