package items

import (
	"errors"
	"itemsModule/clients/elasticsearch"
	"utils/rest_errors"
)

const (
	indexItems = "items"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}
