package animaltypepg

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/vaberof/it-planet-task/internal/domain/animaltype"
	"github.com/vaberof/it-planet-task/internal/infra/storage/postgres/animalpg"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type PostgresAnimalTypeStorage struct {
	db *gorm.DB
}

func NewPostgresAnimalTypeStorage(db *gorm.DB) *PostgresAnimalTypeStorage {
	return &PostgresAnimalTypeStorage{db: db}
}

func (s *PostgresAnimalTypeStorage) CreateAnimalType(chipperId int32, animalType string) (*animaltype.AnimalType, error) {
	var postgresAnimalType AnimalType

	postgresAnimalType.ChipperId = chipperId
	postgresAnimalType.Type = animalType

	err := s.db.Table("animal_types").Create(&postgresAnimalType).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "animaltypepg",
			"func":    "CreateAnimalType",
		}).Error(err)

		return nil, xerror.NewErrorWrapper(500, "cannot create animal type", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "animaltypepg",
		"func":    "CreateAnimalType",
	}).Info("created animal type")

	return BuildDomainAnimalType(&postgresAnimalType), nil
}

func (s *PostgresAnimalTypeStorage) GetAnimalTypeById(id int64) (*animaltype.AnimalType, error) {
	postgresAnimalType, err := s.getAnimalTypeById(id)
	if err != nil {
		return nil, err
	}

	return BuildDomainAnimalType(postgresAnimalType), nil
}

func (s *PostgresAnimalTypeStorage) UpdateAnimalType(typeId int64, animalType string) (*animaltype.AnimalType, error) {
	postgresAnimalType, err := s.getAnimalTypeById(typeId)
	if err != nil {
		return nil, err
	}

	err = s.findAnimalType(animalType)
	if err == nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "animaltypepg",
			"func":    "UpdateAnimalType",
		}).Error("animal type with given type already exists")

		return nil, xerror.NewErrorWrapper(
			409,
			"animal type with given type already exists",
			errors.New("animal type with given type already exists"))
	}

	postgresAnimalType.Type = animalType

	err = s.db.Table("animal_types").Save(&postgresAnimalType).Error
	if err != nil {
		return nil, xerror.NewErrorWrapper(500, "cannot update account", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "animaltypepg",
		"func":    "UpdateAnimalType",
	}).Info("updated animal type")

	return BuildDomainAnimalType(postgresAnimalType), nil
}

func (s *PostgresAnimalTypeStorage) DeleteAnimalType(typeId int64) error {
	postgresAnimalType, err := s.getAnimalTypeById(typeId)
	if err != nil {
		return err
	}

	err = s.db.Table("animal_types").Delete(&postgresAnimalType).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "animaltypepg",
			"func":    "DeleteAnimalType",
		}).Error(err)

		return xerror.NewErrorWrapper(500, "cannot delete animal type", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "animaltypepg",
		"func":    "DeleteAnimalType",
	}).Info("deleted animal type")

	return nil
}

func (s *PostgresAnimalTypeStorage) FindAnimalType(animalType string) error {
	return s.findAnimalType(animalType)
}

func (s *PostgresAnimalTypeStorage) FindAnimalTypeById(typeId int64) error {
	var postgresAnimalType AnimalType

	err := s.db.Table("animal_types").Where("id = ?", typeId).First(&postgresAnimalType).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "animaltypepg",
			"func":    "FindAnimalTypeById",
		}).Error(err)

		return xerror.NewErrorWrapper(404, "cannot find animal type with given type", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "animaltypepg",
		"func":    "FindAnimalTypeById",
	}).Info("found animal type with given type id")

	return nil
}

func (s *PostgresAnimalTypeStorage) FindAnimalTypes(animalTypes []int64) error {
	var count int64
	var query strings.Builder

	convAnimalTypes := make([]string, len(animalTypes))

	for idx, animalType := range animalTypes {
		convAnimalTypes[idx] = strconv.FormatInt(animalType, 10)
	}

	query.WriteString(fmt.Sprintf("id IN (%v)", strings.Join(convAnimalTypes, ",")))

	err := s.db.Table("animal_types").Where(query.String()).Count(&count).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "animaltypepg",
			"func":    "FindAnimalTypes",
		}).Error(err)

		return xerror.NewErrorWrapper(500, "cannot get count of animal types", err)
	}

	if int(count) != len(animalTypes) {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "animaltypepg",
			"func":    "FindAnimalTypes",
		}).Error("cannot find given animal types")

		return xerror.NewErrorWrapper(
			404,
			"cannot find given animal types",
			errors.New("cannot find given animal types"))
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "animaltypepg",
		"func":    "FindAnimalTypes",
	}).Info("found all given animal types")

	return nil
}

func (s *PostgresAnimalTypeStorage) IsAnimalTypeAssociatedWithAnimal(animalTypeId int64) error {
	var postgresAnimal animalpg.Animal
	var query strings.Builder

	query.WriteString(fmt.Sprintf("%d = ANY(animal_types)", animalTypeId))

	err := s.db.Table("animals").Where(query.String()).First(&postgresAnimal).Error
	if err == nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "animaltypepg",
			"func":    "IsAnimalTypeAssociatedWithAnimal",
		}).Error("animal type associated with animal")

		return xerror.NewErrorWrapper(
			409,
			fmt.Sprintf("type associated with animal with %d id", postgresAnimal.Id),
			errors.New("animal type associated with animal"))
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "animaltypepg",
		"func":    "IsAnimalTypeAssociatedWithAnimal",
	}).Info("animal type is not associated with any of animals")

	return nil
}

func (s *PostgresAnimalTypeStorage) getAnimalTypeById(id int64) (*AnimalType, error) {
	var postgresAnimalType AnimalType

	err := s.db.Table("animal_types").Where("id = ?", id).First(&postgresAnimalType).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "animaltypepg",
			"func":    "getAnimalTypeById",
		}).Error(err)

		return nil, xerror.NewErrorWrapper(404, "cannot find animal type with given id", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "animaltypepg",
		"func":    "getAnimalTypeById",
	}).Info("received animal type by id")

	return &postgresAnimalType, nil
}

func (s *PostgresAnimalTypeStorage) findAnimalType(animalType string) error {
	var postgresAnimalType AnimalType

	err := s.db.Table("animal_types").Where("type = ?", animalType).First(&postgresAnimalType).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "animaltypepg",
			"func":    "findAnimalType",
		}).Error(err)

		return xerror.NewErrorWrapper(404, "cannot find animal type with given type", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "animaltypepg",
		"func":    "findAnimalType",
	}).Info("found animal type with type")

	return nil
}
