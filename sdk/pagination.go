package sdk

import (
	"context"

	sayari "github.com/sayari-analytics/sayari-go/generated/go"
)

func (c *Connection) GetAllEntitySearchResults(ctx context.Context, input *sayari.SearchEntity) (data []*sayari.EntityDetails, err error) {
	for {
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

func (c *Connection) GetAllRecordSearchResults(ctx context.Context, input *sayari.SearchRecord) (data []*sayari.RecordDetails, err error) {
	for {
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

func (c *Connection) GetAllTraversalResults(ctx context.Context, entityID sayari.EntityId, input *sayari.Traversal) (data []*sayari.TraversalData, err error) {
	for {
		result, err := c.Traversal.Traversal(ctx, entityID, input)
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
