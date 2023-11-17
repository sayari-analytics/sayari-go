package sdk

import (
	"context"
	"fmt"

	sayari "github.com/sayari-analytics/sayari-go/generated/go"
)

var maxResults = 10000
var ErrTooMuchDataRequested = fmt.Errorf("this request returns %v or more objects. please request individual pages of results, or narrow your request to return fewer objects", maxResults)

func (c *Connection) GetAllEntitySearchResults(ctx context.Context, params *sayari.SearchEntity) (resp *sayari.EntitySearchResults, err error) {
	var data []*sayari.EntityDetails
	for {
		resp, err = c.Search.SearchEntity(ctx, params)
		if err != nil {
			return
		}
		if resp.Size.Count >= maxResults {
			return nil, ErrTooMuchDataRequested
		}
		data = append(data, resp.Data...)
		if !resp.Next {
			break
		}
		params.Offset = sayari.Int(resp.Offset + resp.Limit)
	}
	resp.Limit = len(data)
	resp.Next = false
	resp.Size.Count = len(data)
	resp.Size.Qualifier = "eq"
	resp.Data = data
	return
}

func (c *Connection) GetAllRecordSearchResults(ctx context.Context, params *sayari.SearchRecord) (resp *sayari.RecordSearchResults, err error) {
	var data []*sayari.RecordDetails
	for {
		resp, err = c.Search.SearchRecord(ctx, params)
		if err != nil {
			return
		}
		if resp.Size.Count >= maxResults {
			return nil, ErrTooMuchDataRequested
		}
		data = append(data, resp.Data...)
		if !resp.Next {
			break
		}
		params.Offset = sayari.Int(resp.Offset + resp.Limit)
	}
	resp.Limit = len(data)
	resp.Next = false
	resp.Size.Count = len(data)
	resp.Size.Qualifier = "eq"
	resp.Data = data
	return
}

func (c *Connection) GetAllTraversalResults(ctx context.Context, entityID sayari.EntityId, params *sayari.Traversal) (resp *sayari.TraversalResponse, err error) {
	var data []*sayari.TraversalData
	for {
		resp, err = c.Traversal.Traversal(ctx, entityID, params)
		if err != nil {
			return
		}
		data = append(data, resp.Data...)
		if !resp.Next {
			break
		}
		params.Offset = sayari.Int(resp.Offset + resp.Limit)
	}
	resp.Limit = len(data)
	resp.Next = false
	resp.Data = data
	return
}
