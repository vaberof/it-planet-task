package locationpg

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/vaberof/it-planet-task/internal/domain/location"
	"github.com/vaberof/it-planet-task/internal/infra/storage/postgres/animalpg"
	"github.com/vaberof/it-planet-task/internal/infra/storage/postgres/vstlocationpg"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
	"gorm.io/gorm"
	"strings"
)

type PostgresLocationStorage struct {
	db *gorm.DB
}

func NewPostgresLocationStorage(db *gorm.DB) *PostgresLocationStorage {
	return &PostgresLocationStorage{db: db}
}

func (s *PostgresLocationStorage) CreateLocation(chipperId int32, latitude, longitude float64) (*location.Location, error) {
	var postgresLocation Location

	postgresLocation.ChipperId = chipperId
	postgresLocation.Latitude = latitude
	postgresLocation.Longitude = longitude

	err := s.db.Table("locations").Create(&postgresLocation).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "locationpg",
			"func":    "CreateLocation",
		}).Error(err)

		return nil, xerror.NewErrorWrapper(500, "cannot create location", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "locationpg",
		"func":    "CreateLocation",
	}).Info("created location")

	return BuildDomainLocation(&postgresLocation), nil
}

func (s *PostgresLocationStorage) GetLocation(pointId int64) (*location.Location, error) {
	var postgresLocation Location

	err := s.db.Table("locations").Where("id = ?", pointId).First(&postgresLocation).Error

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "locationpg",
			"func":    "GetLocation",
		}).Error(err)

		return nil, xerror.NewErrorWrapper(404, "cannot find location with given latitude id", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "locationpg",
		"func":    "GetLocation",
	}).Info("received location")

	return BuildDomainLocation(&postgresLocation), nil
}

func (s *PostgresLocationStorage) UpdateLocation(pointId int64, latitude, longitude float64) (*location.Location, error) {
	postgresLocation, err := s.getLocationById(pointId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "locationpg",
			"func":    "UpdateLocation",
		}).Error(err)

		return nil, err
	}

	postgresLocation.Latitude = latitude
	postgresLocation.Longitude = longitude

	err = s.db.Table("locations").Save(&postgresLocation).Error
	if err != nil {
		return nil, xerror.NewErrorWrapper(500, "cannot update location", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "locationpg",
		"func":    "UpdateLocation",
	}).Info("updated location")

	return BuildDomainLocation(postgresLocation), nil
}

func (s *PostgresLocationStorage) DeleteLocation(pointId int64) error {
	postgresLocation, err := s.getLocationById(pointId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "locationpg",
			"func":    "DeleteLocation",
		}).Error(err)

		return err
	}

	err = s.db.Table("locations").Delete(&postgresLocation).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "locationpg",
			"func":    "DeleteLocation",
		}).Error(err)

		return xerror.NewErrorWrapper(500, "cannot delete location", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "locationpg",
		"func":    "DeleteLocation",
	}).Info("deleted location")

	return nil
}

func (s *PostgresLocationStorage) FindLocationById(id int64) error {
	_, err := s.getLocationById(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresLocationStorage) FindLocation(latitude, longitude float64) error {
	var postgresLocation Location

	err := s.db.
		Table("locations").
		Where("latitude = ? AND longitude = ?", latitude, longitude).
		First(&postgresLocation).Error

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "locationpg",
			"func":    "FindLocation",
		}).Error(err)

		return xerror.NewErrorWrapper(404, "cannot find location with given latitude and longitude", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "locationpg",
		"func":    "FindLocation",
	}).Info("found location")

	return nil
}

func (s *PostgresLocationStorage) IsLocationAssociatedWithAnimal(pointId int64) error {
	var postgresAnimal animalpg.Animal
	var query strings.Builder

	query.WriteString(fmt.Sprintf("chipping_location_id = %d", pointId))

	err := s.db.Table("animals").Where(query.String()).First(&postgresAnimal).Error
	if err == nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "locationpg",
			"func":    "IsLocationAssociatedWithAnimal",
		}).Error("location associated with animal")

		return xerror.NewErrorWrapper(
			400,
			fmt.Sprintf("location associated with animal with %d id", postgresAnimal.Id),
			errors.New("cannot delete associated account"))
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "locationpg",
		"func":    "IsLocationAssociatedWithAnimal",
	}).Info("location is not associated with any of animals")

	return nil
}

func (s *PostgresLocationStorage) IsLocationAssociatedWithAnimalsVisitedLocation(pointId int64) error {
	var postgresVisitedLocation vstlocationpg.VisitedLocation
	var query strings.Builder

	query.WriteString(fmt.Sprintf("location_point_id = %d", pointId))

	err := s.db.Table("visited_locations").Where(query.String()).First(&postgresVisitedLocation).Error
	if err == nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "locationpg",
			"func":    "IsLocationAssociatedWithAnimalsVisitedLocation",
		}).Error("visited location associated with animal")

		return xerror.NewErrorWrapper(
			400,
			"location associated with animal",
			errors.New("cannot delete associated location"))
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "locationpg",
		"func":    "IsLocationAssociatedWithAnimalsVisitedLocation",
	}).Info("location is not associated with any of animals")

	return nil
}

func (s *PostgresLocationStorage) getLocationById(id int64) (*Location, error) {
	var postgresLocation Location

	err := s.db.Table("locations").Where("id = ?", id).First(&postgresLocation).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "locationpg",
			"func":    "getLocationById",
		}).Error(err)

		return nil, xerror.NewErrorWrapper(404, "cannot find location with given id", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "locationpg",
		"func":    "getLocationById",
	}).Info("received location by id")

	return &postgresLocation, nil
}
