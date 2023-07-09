package main

import (
	"crypto/rand"
	"encoding/binary"
	"log"

	"github.com/spf13/viper"
)

type randGenSrc struct{}

func (s *randGenSrc) Seed(int64) {}

func (s *randGenSrc) Uint64() (value uint64) {
	_ = binary.Read(rand.Reader, binary.BigEndian, &value)

	return value
}

// URLDao ...
type URLDao interface {
	save(url URL) (int, error)
	update(id int, oldURL, newURL URL) (int, error)
	findByID(id int) (URL, error)
	findAllByUser() ([]URLStat, error)
}

// UserDAO ....
type UserDAO interface {
	addUser(username, password string) (interface{}, error)
	userExists(username string) (bool, error)
	findByUsername(username string) (interface{}, error)
	validateUserAndPassword(username, password string) (bool, error)
	findAll() ([]interface{}, error)
}

// StatsDAO ...
type StatsDAO interface {
	save(shortURL string, headers *map[string][]string) (int, error)
	findByShortID(id int) ([]interface{}, error)
	findAll() (map[int][]StatsInMemory, error)
}

func factoryStatsDao(config *viper.Viper) *StatsDAO {
	var dao StatsDAO

	engine := config.GetString("dbengine")

	switch engine {
	case "memory":
		dao = StatsDAOMemoryImpl{
			db: map[int][]StatsInMemory{},
		}
	default:
		log.Fatalf("error: wrong engine: %s", engine)

		return nil
	}

	return &dao
}

func factoryURLDao(config *viper.Viper) *URLDao {
	var dao URLDao

	engine := config.GetString("dbengine")

	switch engine {
	case "memory":
		dao = InMemoryURLDAOImpl{
			DB: &memoryDB{
				db: map[int]string{},
			},
		}
	default:
		log.Fatalf("error: wrong engine: %s", engine)

		return nil
	}

	return &dao
}

func factoryUserDAO(config *viper.Viper) *UserDAO {
	var userDAO UserDAO

	engine := config.GetString("dbengine")

	switch engine {
	case "memory":
		userDAO = InMemoryUserDAOImpl{
			db:       map[string]UserInMemory{},
			rndIDGen: randGenSrc{},
		}

	default:
		log.Fatalf("error: wrong engine: %s", engine)

		return nil
	}

	return &userDAO
}
