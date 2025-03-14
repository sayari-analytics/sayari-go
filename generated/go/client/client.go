// This file was auto-generated by Fern from our API Definition.

package client

import (
	attributes "github.com/sayari-analytics/sayari-go/generated/go/attributes"
	auth "github.com/sayari-analytics/sayari-go/generated/go/auth"
	core "github.com/sayari-analytics/sayari-go/generated/go/core"
	entity "github.com/sayari-analytics/sayari-go/generated/go/entity"
	info "github.com/sayari-analytics/sayari-go/generated/go/info"
	internal "github.com/sayari-analytics/sayari-go/generated/go/internal"
	metadata "github.com/sayari-analytics/sayari-go/generated/go/metadata"
	negativenews "github.com/sayari-analytics/sayari-go/generated/go/negativenews"
	notifications "github.com/sayari-analytics/sayari-go/generated/go/notifications"
	option "github.com/sayari-analytics/sayari-go/generated/go/option"
	project "github.com/sayari-analytics/sayari-go/generated/go/project"
	record "github.com/sayari-analytics/sayari-go/generated/go/record"
	resolution "github.com/sayari-analytics/sayari-go/generated/go/resolution"
	resource "github.com/sayari-analytics/sayari-go/generated/go/resource"
	search "github.com/sayari-analytics/sayari-go/generated/go/search"
	source "github.com/sayari-analytics/sayari-go/generated/go/source"
	supplychain "github.com/sayari-analytics/sayari-go/generated/go/supplychain"
	trade "github.com/sayari-analytics/sayari-go/generated/go/trade"
	traversal "github.com/sayari-analytics/sayari-go/generated/go/traversal"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header

	Attributes    *attributes.Client
	Auth          *auth.Client
	Entity        *entity.Client
	Info          *info.Client
	Metadata      *metadata.Client
	NegativeNews  *negativenews.Client
	Notifications *notifications.Client
	Project       *project.Client
	Record        *record.Client
	Resolution    *resolution.Client
	Resource      *resource.Client
	Search        *search.Client
	Source        *source.Client
	SupplyChain   *supplychain.Client
	Trade         *trade.Client
	Traversal     *traversal.Client
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header:        options.ToHeader(),
		Attributes:    attributes.NewClient(opts...),
		Auth:          auth.NewClient(opts...),
		Entity:        entity.NewClient(opts...),
		Info:          info.NewClient(opts...),
		Metadata:      metadata.NewClient(opts...),
		NegativeNews:  negativenews.NewClient(opts...),
		Notifications: notifications.NewClient(opts...),
		Project:       project.NewClient(opts...),
		Record:        record.NewClient(opts...),
		Resolution:    resolution.NewClient(opts...),
		Resource:      resource.NewClient(opts...),
		Search:        search.NewClient(opts...),
		Source:        source.NewClient(opts...),
		SupplyChain:   supplychain.NewClient(opts...),
		Trade:         trade.NewClient(opts...),
		Traversal:     traversal.NewClient(opts...),
	}
}
