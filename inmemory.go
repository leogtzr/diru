package main

import (
	"time"
)

type memoryDB struct {
	db            map[int]string
	autoIncrement int
}

// InMemoryURLDAOImpl ...
type InMemoryURLDAOImpl struct {
	DB *memoryDB
}

// StatsDAOMemoryImpl ...
type StatsDAOMemoryImpl struct {
	db map[int][]Stats
}

func (im InMemoryURLDAOImpl) save(url URL) (int, error) {
	mu.Lock()
	defer mu.Unlock()

	im.DB.autoIncrement++
	id := im.DB.autoIncrement
	im.DB.db[id] = url.URL

	return id, nil
}

func (im InMemoryURLDAOImpl) findAllByUser() ([]URLStat, error) {
	// shortID:int, url:string
	var urls []URLStat

	// dummy impl...
	for shortID, url := range im.DB.db {
		urls = append(urls, URLStat{
			ShortID: shortID,
			URL:     url,
		})
	}

	return urls, nil
}

func (im InMemoryURLDAOImpl) findByID(id int) (URL, error) {
	u, found := im.DB.db[id]
	if found {
		url := URL{
			URL: u,
		}

		return url, nil
	}

	return URL{}, errorURLNotFound(id)
}

func (im InMemoryURLDAOImpl) update(id int, newURL URL) (int, error) {
	mu.Lock()
	defer mu.Unlock()

	if _, ok := im.DB.db[id]; !ok {
		return id, errorURLNotFound(id)
	}

	newID := shortURLToID(newURL.URL, chars)
	url := im.DB.db[id]

	im.DB.db[newID] = url
	delete(im.DB.db, id)

	return newID, nil
}

func (dao StatsDAOMemoryImpl) save(shortURL string, headers *map[string][]string) (int, error) {
	urlShortID := shortURLToID(shortURL, chars)

	stat := Stats{
		CreatedAt: time.Now(),
		ShortID:   urlShortID,
		Headers:   *headers,
	}

	dao.db[urlShortID] = append(dao.db[urlShortID], stat)

	return 0, nil
}

func (dao StatsDAOMemoryImpl) findByShortID(shortID int) ([]Stats, error) {
	for urlShortID, stats := range dao.db {
		if urlShortID == shortID {
			return stats, nil
		}
	}

	return []Stats{}, nil
}

func (dao StatsDAOMemoryImpl) findAll() (map[int][]Stats, error) {
	return dao.db, nil
}
