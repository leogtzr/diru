package main

// URLDao ...
type URLDao interface {
	save(url URL) (int, error)
	update(id int, oldURL, newURL URL) (int, error)
	findByID(id int) (URL, error)
	findAllByUser() ([]URLStat, error)
}

// UserDAO ....
type UserDAO interface {
	findAll() ([]interface{}, error)
}

// StatsDAO ...
type StatsDAO interface {
	save(shortURL string, headers *map[string][]string) (int, error)
	findByShortID(id int) ([]interface{}, error)
	findAll() (map[int][]StatsInMemory, error)
}

func factoryStatsDao() *StatsDAO {
	var dao StatsDAO = StatsDAOMemoryImpl{
		db: map[int][]StatsInMemory{},
	}

	return &dao
}

func factoryURLDao() *URLDao {
	var dao URLDao = InMemoryURLDAOImpl{
		DB: &memoryDB{
			db: map[int]string{},
		},
	}

	return &dao
}
