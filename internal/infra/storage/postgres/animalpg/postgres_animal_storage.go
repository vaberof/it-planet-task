package animalpg

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/vaberof/it-planet-task/internal/domain/animal"
	"github.com/vaberof/it-planet-task/internal/domain/vstlocation"
	"github.com/vaberof/it-planet-task/internal/infra/storage/postgres/vstlocationpg"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
	"gorm.io/gorm"
	"strings"
	"time"
)

type PostgresAnimalStorage struct {
	db                 *gorm.DB
	vstLocationStorage VisitedLocationStorage
}

func NewPostgresAnimalStorage(db *gorm.DB, vstLocationStorage VisitedLocationStorage) *PostgresAnimalStorage {
	return &PostgresAnimalStorage{
		db:                 db,
		vstLocationStorage: vstLocationStorage,
	}
}

func (s *PostgresAnimalStorage) CreateAnimal(
	chipperId int32,
	chippingLocationId int64,
	animalTypes []int64,
	weight float32,
	length float32,
	height float32,
	gender string) (*animal.Animal, error) {

	var postgresAnimal Animal

	postgresAnimal.AnimalTypes = animalTypes
	postgresAnimal.Weight = weight
	postgresAnimal.Length = length
	postgresAnimal.Height = height
	postgresAnimal.Gender = gender
	postgresAnimal.LifeStatus = animal.LifeStatusAlive
	postgresAnimal.ChipperId = chipperId
	postgresAnimal.ChippingLocationId = chippingLocationId

	err := s.db.Table("animals").Create(&postgresAnimal).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "animalpg",
			"func":    "CreateAnimal",
		}).Error(err)

		return nil, xerror.NewErrorWrapper(500, "cannot create animal", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "animalpg",
		"func":    "CreateAnimal",
	}).Info("created animal")

	return BuildDomainAnimal(&postgresAnimal), nil
}

func (s *PostgresAnimalStorage) GetAnimalById(id int64) (*animal.Animal, error) {
	postgresAnimal, err := s.getAnimalById(id)
	if err != nil {
		return nil, err
	}

	return BuildDomainAnimal(postgresAnimal), nil
}

func (s *PostgresAnimalStorage) SearchAnimals(
	startDateTime *time.Time,
	endDateTime *time.Time,
	chipperId int32,
	chippingLocationId int64,
	lifeStatus string,
	gender string,
	from int32,
	size int32) ([]*animal.Animal, error) {

	var postgresAnimals []*Animal

	var query strings.Builder

	query.WriteString("1=1")

	if startDateTime != nil {
		query.WriteString(fmt.Sprint(" AND chipping_date_time >= ", "'"+(*startDateTime).Format("2006-01-02T15:04:05-07:00")) + "'")
	}

	if endDateTime != nil {
		query.WriteString(fmt.Sprint(" AND chipping_date_time <= ", "'"+(*endDateTime).Format("2006-01-02T15:04:05-07:00")) + "'")
	}

	if chipperId != -1 {
		query.WriteString(fmt.Sprint(" AND chipper_id = ", "'"+fmt.Sprint(chipperId)+"'"))
	}

	if chippingLocationId != -1 {
		query.WriteString(fmt.Sprint(" AND chipping_location_id = ", "'"+fmt.Sprint(chippingLocationId)+"'"))
	}

	if lifeStatus != "" {
		query.WriteString(fmt.Sprint(" AND life_status = ", "'"+lifeStatus+"'"))
	}

	if gender != "" {
		query.WriteString(fmt.Sprint(" AND gender = ", "'"+gender+"'"))
	}

	err := s.db.
		Where(query.String()).
		Order("id ASC").
		Offset(int(from)).
		Limit(int(size)).
		Preload("VisitedLocations").
		Find(&postgresAnimals).Error

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "animalpg",
			"func":    "SearchAnimals",
		}).Error(err)

		return nil, xerror.NewErrorWrapper(404, "cannot find filtered animals", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "animalpg",
		"func":    "SearchAnimals",
	}).Info("received filtered animals")

	return BuildDomainAnimals(postgresAnimals), nil
}

func (s *PostgresAnimalStorage) UpdateAnimal(
	animalId int64,
	weight float32,
	length float32,
	height float32,
	gender string,
	lifeStatus string,
	chipperId int32,
	chippingLocationId int64) (*animal.Animal, error) {

	postgresAnimal, err := s.getAnimalById(animalId)
	if err != nil {
		return nil, err
	}

	postgresAnimal.Weight = weight
	postgresAnimal.Length = length
	postgresAnimal.Height = height
	postgresAnimal.Gender = gender
	postgresAnimal.LifeStatus = lifeStatus
	postgresAnimal.ChipperId = chipperId
	postgresAnimal.ChippingLocationId = chippingLocationId

	if lifeStatus == animal.LifeStatusDead {
		parsedDeathDateTime, err := time.Parse("2006-01-02T15:04:05-07:00", time.Now().Format("2006-01-02T15:04:05-07:00"))
		if err != nil {
			return nil, xerror.NewErrorWrapper(500, "cannot parse death time", err)
		}
		postgresAnimal.DeathDateTime = &parsedDeathDateTime
	}

	err = s.db.Table("animals").Save(&postgresAnimal).Error
	if err != nil {
		return nil, xerror.NewErrorWrapper(500, "cannot update animal", err)
	}

	return BuildDomainAnimal(postgresAnimal), nil
}

func (s *PostgresAnimalStorage) DeleteAnimal(animalId int64) error {
	postgresAnimal, err := s.getAnimalById(animalId)
	if err != nil {
		return err
	}

	err = s.db.Table("animals").Where("id = ?", animalId).Delete(&postgresAnimal).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "animalpg",
			"func":    "DeleteAnimal",
		}).Error(err)

		return xerror.NewErrorWrapper(500, "cannot delete animal", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "animalpg",
		"func":    "DeleteAnimal",
	}).Info("deleted animal")

	return nil
}

func (s *PostgresAnimalStorage) FindAnimalById(id int64) error {
	var postgresAnimal Animal

	err := s.db.Table("animals").Where("id = ?", id).First(&postgresAnimal).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "animalpg",
			"func":    "FindAnimalById",
		}).Error(err)

		return xerror.NewErrorWrapper(404, "cannot find animal with given id", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "animalpg",
		"func":    "FindAnimalById",
	}).Info("found animal with given id")

	return nil
}

func (s *PostgresAnimalStorage) AddAnimalsType(animalId, animalTypeId int64) (*animal.Animal, error) {
	postgresAnimal, err := s.getAnimalById(animalId)
	if err != nil {
		return nil, err
	}

	postgresAnimal.AnimalTypes = append(postgresAnimal.AnimalTypes, animalTypeId)

	err = s.db.Table("animals").Save(&postgresAnimal).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "animalpg",
			"func":    "AddAnimalsType",
		}).Error(err)

		return nil, xerror.NewErrorWrapper(500, "cannot save updated animal", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "animalpg",
		"func":    "AddAnimalsType",
	}).Info("added animal type")

	return BuildDomainAnimal(postgresAnimal), nil
}

func (s *PostgresAnimalStorage) UpdateAnimalsType(animalId, oldTypeId, newTypeId int64) (*animal.Animal, error) {
	postgresAnimal, err := s.getAnimalById(animalId)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(postgresAnimal.AnimalTypes); i++ {
		if postgresAnimal.AnimalTypes[i] == oldTypeId {
			postgresAnimal.AnimalTypes[i] = newTypeId
			break
		}
	}

	err = s.db.Table("animals").Save(&postgresAnimal).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "animalpg",
			"func":    "UpdateAnimalsType",
		}).Error(err)

		return nil, xerror.NewErrorWrapper(500, "cannot save updated animal", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "animalpg",
		"func":    "UpdateAnimalsType",
	}).Info("updated animal type")

	return BuildDomainAnimal(postgresAnimal), nil
}

func (s *PostgresAnimalStorage) DeleteAnimalsType(animalId, animalTypeId int64) (*animal.Animal, error) {
	postgresAnimal, err := s.getAnimalById(animalId)
	if err != nil {
		return nil, err
	}

	for i, animalType := range postgresAnimal.AnimalTypes {
		if animalType == animalTypeId {
			postgresAnimal.AnimalTypes = append(postgresAnimal.AnimalTypes[:i], postgresAnimal.AnimalTypes[i+1:]...)
			break
		}
	}

	err = s.db.Table("animals").Save(&postgresAnimal).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "animalpg",
			"func":    "DeleteAnimalsType",
		}).Error(err)

		return nil, xerror.NewErrorWrapper(500, "cannot save updated animal", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "animalpg",
		"func":    "DeleteAnimalsType",
	}).Info("deleted animal type")

	return BuildDomainAnimal(postgresAnimal), nil
}

func (s *PostgresAnimalStorage) GetAnimalsVisitedLocations(
	animalId int64,
	startDateTime *time.Time,
	endDateTime *time.Time,
	from int32,
	size int32) ([]*vstlocation.VisitedLocation, error) {

	var postgresVisitedLocations []*vstlocationpg.VisitedLocation

	var query strings.Builder

	query.WriteString("1=1")

	query.WriteString(fmt.Sprintf(" AND animal_id = '%d'", animalId))

	if startDateTime != nil {
		query.WriteString(fmt.Sprint(" AND date_time_of_visit_location_point >= ", "'"+(*startDateTime).Format("2006-01-02T15:04:05-07:00")) + "'")
	}

	if endDateTime != nil {
		query.WriteString(fmt.Sprint(" AND date_time_of_visit_location_point <= ", "'"+(*endDateTime).Format("2006-01-02T15:04:05-07:00")) + "'")
	}

	err := s.db.
		Table("visited_locations").
		Where(query.String()).
		Order("date_time_of_visit_location_point ASC").
		Offset(int(from)).
		Limit(int(size)).
		Find(&postgresVisitedLocations).Error

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "animalpg",
			"func":    "GetAnimalsVisitedLocations",
		}).Error(err)

		return nil, xerror.NewErrorWrapper(404, "cannot find filtered animal visited locations", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "animalpg",
		"func":    "GetAnimalsVisitedLocations",
	}).Info("received visited locations")

	return vstlocationpg.BuildDomainVisitedLocations(postgresVisitedLocations), nil
}

func (s *PostgresAnimalStorage) UpdateAnimalsVisitedLocation(visitedLocationId, locationId int64) (*vstlocation.VisitedLocation, error) {
	postgresVisitedLocation, err := s.vstLocationStorage.GetVisitedLocationById(visitedLocationId)
	if err != nil {
		return nil, err
	}

	postgresVisitedLocation.LocationPointId = locationId

	err = s.db.Table("visited_locations").Save(&postgresVisitedLocation).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "animalpg",
			"func":    "UpdateAnimalsVisitedLocation",
		}).Error(err)

		return nil, xerror.NewErrorWrapper(500, "cannot update visited location", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "animalpg",
		"func":    "UpdateAnimalsVisitedLocation",
	}).Info("updated visited locations")

	return vstlocationpg.BuildDomainVisitedLocation(postgresVisitedLocation), nil
}

func (s *PostgresAnimalStorage) DeleteAnimalsVisitedLocation(animalId int64, index int) error {
	postgresAnimal, err := s.getAnimalById(animalId)
	if err != nil {
		return err
	}

	if index == 0 && index+1 < len(postgresAnimal.VisitedLocations) &&
		postgresAnimal.VisitedLocations[index+1].LocationPointId == postgresAnimal.ChippingLocationId {

		err = s.db.
			Model(&postgresAnimal).
			Association("VisitedLocations").
			Delete(postgresAnimal.VisitedLocations[index], postgresAnimal.VisitedLocations[index+1])

		if err != nil {
			logrus.WithFields(logrus.Fields{
				"layer":   "infra",
				"package": "animalpg",
				"func":    "DeleteAnimalsVisitedLocation",
			}).Error(err)

			return xerror.NewErrorWrapper(500, "cannot delete visited locations", err)
		}

		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "animalpg",
			"func":    "DeleteAnimalsVisitedLocation",
		}).Info("deleted 2 visited locations")

		return nil
	}

	err = s.db.
		Model(&postgresAnimal).
		Association("VisitedLocations").
		Delete(postgresAnimal.VisitedLocations[index])

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "animalpg",
			"func":    "DeleteAnimalsVisitedLocation",
		}).Error(err)

		return xerror.NewErrorWrapper(500, "cannot delete visited location", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "animalpg",
		"func":    "DeleteAnimalsVisitedLocation",
	}).Info("deleted 1 visited location")

	return nil
}

func (s *PostgresAnimalStorage) getAnimalById(id int64) (*Animal, error) {
	var postgresAnimal Animal

	err := s.db.Where("id = ?", id).Preload("VisitedLocations").First(&postgresAnimal).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "animalpg",
			"func":    "getAnimalById",
		}).Error(err)

		return nil, xerror.NewErrorWrapper(404, "cannot find animal with given id", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "animalpg",
		"func":    "getAnimalById",
	}).Info("received postgres animal by id")

	return &postgresAnimal, nil
}
