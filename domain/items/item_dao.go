package items

import (
	"errors"
	"fmt"
	"itemsModule/clients/elasticsearch"
	"utils/rest_errors"
)

const (
	indexItems = "items"
	typeItem   = "_doc"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, typeItem, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Get() rest_errors.RestErr {
	itemId := i.Id
	result, err := elasticsearch.Client.Get(indexItems, typeItem, i.Id)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to get", errors.New("database error"))
	}
	if !result.Found {
		return rest_errors.NewNotFoundError("no item found with this id")
	}
	bytes, err := result.Source.MarshalJSON()
	i.Id = itemId
	fmt.Println(string(bytes))
	return nil
}
