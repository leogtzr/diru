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

// Stats ...
type Stats struct {
	CreatedAt time.Time
	ShortID   int
	UserID    int
	Headers   map[string][]string
}
