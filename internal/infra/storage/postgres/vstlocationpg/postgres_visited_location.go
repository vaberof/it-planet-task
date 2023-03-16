package vstlocationpg

import "time"

type VisitedLocation struct {
	Id                           int64
	DateTimeOfVisitLocationPoint time.Time `gorm:"autoCreateTime"`
	LocationPointId              int64
	ChipperId                    int32
	AnimalId                     int64 `gorm:"default:null"`
}
