// Package domain represent the structue of elements
package domain

type GeoPoint struct {
	Lat float64 `json:"lat" bson:"lat"`
	Lng float64 `json:"lng" bson:"lng"`
}

type Book struct {
	ID       string   `json:"id" bson:"_id,omitempty"`
	OwnerID  string   `json:"ownerId" bson:"ownerId"`
	Title    string   `json:"title" bson:"title"`
	Author   string   `json:"author" bson:"author"`
	Photos   []string `json:"photos" bson:"photos"`
	Price    float64  `json:"price" bson:"price"`
	Lend     bool     `json:"lend" bson:"lend"`
	Location GeoPoint `json:"location" bson:"location"`
}
