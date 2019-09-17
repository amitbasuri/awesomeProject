package awesome

import (
	"github.com/graph-gophers/graphql-go"
)

type Item struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description,omitempty"`
	Price       float64      `json:"price"`
	CreatedAt   graphql.Time `json:"created_at"`
}

type Service interface {
	CreateItem(item *Item) error
	GetItemByID(id string) (*Item, error)
}
