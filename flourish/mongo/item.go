package mongo

import (
	"awesomeProject/flourish"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var MongoSvc Service

type Service struct {
	DB *mgo.Database
}

func (svc *Service) CreateItem(item *flourish.Item) error {
	c := svc.DB.C("item")
	return c.Insert(item)
}

func (svc *Service) GetItemByID(id string) (*flourish.Item, error) {
	c := svc.DB.C("item")
	item := &flourish.Item{}
	if err := c.Find(bson.M{"id": id}).One(item); err != nil {
		return nil, err
	}
	fmt.Println("found ", item)
	return item, nil
}
