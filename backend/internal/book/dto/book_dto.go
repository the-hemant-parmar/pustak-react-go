package dto

type CreateBookRequest struct {
	OwnerID string   `json:"ownerId"`
	Title   string   `json:"title"`
	Author  string   `json:"author"`
	Photos  []string `json:"photos"`
	Price   float64  `json:"price"`
	Lend    bool     `json:"lend"`
	Lat     float64  `json:"lat"`
	Lng     float64  `json:"lng"`
}

type UpdateBookRequest struct {
	Title  *string  `json:"title"`
	Author *string  `json:"author"`
	Price  *float64 `json:"price"`
	Lend   *bool    `json:"lend"`
}
