package main

import (
	"time"
)

// URL ...
type URL struct {
	URL string `form:"url"`
}

// URLChange ...
type URLChange struct {
	ShortURL string `form:"url"`
	NewURL   string `form:"new_url"`
}

// URLStat ...
type URLStat struct {
	ShortID int    `json:"id"`
	URL     string `json:"url"`
}

// URLStatFull is basically a URLStat but instead of the short ID, it has the short URL corresponding
// to the short ID value.
type URLStatFull struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

// UserInMemory ...
type UserInMemory struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	User      string
	Password  string
}

// StatsInMemory ...
type Stats struct {
	CreatedAt time.Time
	ShortID   int
	UserID    int
	Headers   map[string][]string
}

// StatsHeadersPostgresql ...
type StatsHeadersPostgresql struct {
	/*
			id serial PRIMARY KEY NOT NULL,
			name varchar(150) NOT NULL,
			value varchar(500) NOT NULL,
			stat_id int NOT NULL,
		    constraint fk_stats_headers
		        foreign key (stat_id)
		        REFERENCES stats (id)
	*/

}
