package sdk

import (
	"context"

	sayari "github.com/sayari-analytics/sayari-go/generated/go"
)

func (c *Connection) GetAllEntitySearchResults(ctx context.Context, params *sayari.SearchEntity) (data []*sayari.EntityDetails, err error) {
	for {
		result, err := c.Search.SearchEntity(ctx, params)
		if err != nil {
			return nil, err
		}
		data = append(data, result.Data...)
		if !result.Next {
			break
		}
		params.Offset = Int(result.Offset + result.Limit)
	}
	return data, nil
}

func (c *Connection) GetAllRecordSearchResults(ctx context.Context, params *sayari.SearchRecord) (data []*sayari.RecordDetails, err error) {
	for {
		result, err := c.Search.SearchRecord(ctx, params)
		if err != nil {
			return nil, err
		}
		data = append(data, result.Data...)
		if !result.Next {
			break
		}
		params.Offset = Int(result.Offset + result.Limit)
	}
	return data, nil
}

func (c *Connection) GetAllTraversalResults(ctx context.Context, entityID sayari.EntityId, params *sayari.Traversal) (data []*sayari.TraversalData, err error) {
	for {
		result, err := c.Traversal.Traversal(ctx, entityID, params)
		if err != nil {
			return nil, err
		}
		data = append(data, result.Data...)
		if !result.Next {
			break
		}
		params.Offset = Int(result.Offset + result.Limit)
	}
	return data, nil
}
