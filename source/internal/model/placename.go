package model

// PlaceName represents a difficult-to-read place name in Hokkaido.
type PlaceName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Yomi string `json:"yomi"`
}
