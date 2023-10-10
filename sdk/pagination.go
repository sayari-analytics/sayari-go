package sdk

import (
	"context"
	sayari "github.com/sayari-analytics/sayari-go/generated/go"
	"github.com/sayari-analytics/sayari-go/generated/go/client"
)

func GetAllEntitySearchResults(ctx context.Context, c *client.Client, input *sayari.SearchEntity) (data []*sayari.EntityDetails, err error) {
	for true {
		result, err := c.Search.SearchEntity(ctx, input)
		if err != nil {
			return nil, err
		}
		data = append(data, result.Data...)
		if !result.Next {
			break
		}
		input.Offset = Int(result.Offset + result.Limit)
	}
	return data, nil
}

func GetAllRecordSearchResults(ctx context.Context, c *client.Client, input *sayari.SearchRecord) (data []*sayari.RecordDetails, err error) {
	for true {
		result, err := c.Search.SearchRecord(ctx, input)
		if err != nil {
			return nil, err
		}
		data = append(data, result.Data...)
		if !result.Next {
			break
		}
		input.Offset = Int(result.Offset + result.Limit)
	}
	return data, nil
}

func GetAllTraversalResults(ctx context.Context, c *client.Client, entityId sayari.EntityId, input *sayari.Traversal) (data []*sayari.TraversalData, err error) {
	for true {
		result, err := c.Traversal.Traversal(ctx, entityId, input)
		if err != nil {
			return nil, err
		}
		data = append(data, result.Data...)
		if !result.Next {
			break
		}
		input.Offset = Int(result.Offset + result.Limit)
	}
	return data, nil
}
