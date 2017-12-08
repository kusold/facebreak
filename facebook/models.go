package facebook

import "time"

// Response from Facebook's graphql api
type Response struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Photos *photos `json:"photos"`
}

type photos struct {
	Data []*photo `json:"data"`
}
type photo struct {
	Name        string    `json:"name"`
	ID          string    `json:"id"`
	Tags        *tags     `json:"tags"`
	Images      []*image  `json:"images"`
	CreatedTime time.Time `json:"created_at"`
}
type tags struct {
	Data []*tag `json:"data"`
}
type tag struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	CreatedTime time.Time `json:"created_at"`
	X           float64   `json:"x"`
	Y           float64   `json:"y"`
}
type image struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Source string `json:"source"`
}
