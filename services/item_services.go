package services

import (
	"itemsModule/domain/items"
	"net/http"
	"utils/rest_errors"
)

var ItemsService ItemsServiceInterface = &itemsService{}

type ItemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
}

type itemsService struct {
	//Create(items.Item) (*items.Item, rest_errors.RestErr)
	//Get(string) (*items.Item, rest_errors.RestErr)
}

func (s *itemsService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
	return nil, rest_errors.NewRestError("Implement me", http.StatusNotImplemented, "Not implemented", nil)
}

func (s *itemsService) Get(id string) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("Implement me", http.StatusNotImplemented, "Not implemented", nil)
}
