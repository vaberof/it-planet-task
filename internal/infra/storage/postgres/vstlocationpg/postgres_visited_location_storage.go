package vstlocationpg

import (
	"github.com/sirupsen/logrus"
	"github.com/vaberof/it-planet-task/internal/domain/vstlocation"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
	"gorm.io/gorm"
)

type PostgresVisitedLocationStorage struct {
	db *gorm.DB
}

func NewPostgresVisitedLocationStorage(db *gorm.DB) *PostgresVisitedLocationStorage {
	return &PostgresVisitedLocationStorage{db: db}
}

func (s *PostgresVisitedLocationStorage) CreateVisitedLocation(animalId, locationPointId int64, chipperId int32) (*vstlocation.VisitedLocation, error) {
	var postgresVisitedLocation VisitedLocation

	postgresVisitedLocation.LocationPointId = locationPointId
	postgresVisitedLocation.ChipperId = chipperId
	postgresVisitedLocation.AnimalId = animalId

	err := s.db.Table("visited_locations").Create(&postgresVisitedLocation).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "vstlocationpg",
			"func":    "CreateVisitedLocation",
		}).Error(err)

		return nil, xerror.NewErrorWrapper(500, "cannot create visited location", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "vstlocationpg",
		"func":    "CreateVisitedLocation",
	}).Info("created visited location")

	return BuildDomainVisitedLocation(&postgresVisitedLocation), nil
}

func (s *PostgresVisitedLocationStorage) GetVisitedLocationById(visitedLocationId int64) (*VisitedLocation, error) {
	var postgresVisitedLocation VisitedLocation

	err := s.db.Table("visited_locations").Where("id = ?", visitedLocationId).First(&postgresVisitedLocation).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "vstlocationpg",
			"func":    "GetVisitedLocationById",
		}).Error(err)

		return nil, xerror.NewErrorWrapper(404, "cannot find visited location with given id", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "vstlocationpg",
		"func":    "GetVisitedLocationById",
	}).Info("received visited location by id")

	return &postgresVisitedLocation, nil
}
