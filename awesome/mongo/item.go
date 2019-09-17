package mongo

import (
	"awesomeProject/awesome"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Service struct {
	DB *mgo.Database
}

func (svc *Service) CreateItem(item *awesome.Item) error {
	c := svc.DB.C("item")
	return c.Insert(item)
}

func (svc *Service) GetItemByID(id string) (*awesome.Item, error) {
	c := svc.DB.C("item")
	item := &awesome.Item{}
	if err := c.Find(bson.M{"id": id}).One(item); err != nil {
		return nil, err
	}
	return item, nil
}
