---
title: Sayari Go SDK
category: 653044fd9479c1000c221860
---
# Introduction

Welcome to the Sayari Graph SDK for Go. The goal of this project is to get you up and running as quickly as possible
so you can start benefiting from the power of Sayari Graph. In the new few sections you will learn how to setup and
use the Sayari Graph SDK. We also document some example use cases to show how easy it is to build the power of Sayari
Graph into your application.

# Setup
## Prerequisites
The only thing you need to start using this SDK are your Client_ID and Client_Secret provided to you by Sayari. (@Aleks to add info about getting these creds)

## Installation
To install this SDK, simply run `go get "github.com/sayari-analytics/sayari-go/..."`
Then simply import "github.com/sayari-analytics/sayari-go/sdk" into your go code to use the SDK.

# Quickstart
This section will walk you through a basic example of connecting to Sayari Graph, resolving and entity, and getting that
entity's detailed information.

## Connecting
To connect to Sayari Graph, simply create a client object by calling the SDK's 'Connect' method and passing in your
client ID and secret. **Note**: For security purposes, it is highly recommended that you don't hardcode your client
ID and secret in your code. Instead, simply export them as environment variables and use those.

```go
client, err := sdk.Connect(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
if err != nil {
    log.Fatalf("Error: %v", err)
}
```

## Resolving an entity
Now that we have a client, we can use the Resolution method to find an entity. To do this we create a resolution request
with the entity information we are using to search. Full documentation of this endpoint can be seen in the API docs.

A request to resolve an entity with the name "Victoria Beckham" is shown below:
```go
// Create the request body
resolutionRequest := sayari.Resolution{Name: []*string{sdk.String("Victoria Beckham")}}

// Make the request and handle the error
resolution, err := client.Resolution.Resolution(context.Background(), &resolutionRequest)
if err != nil {
    log.Fatalf("Error: %v", err)
}
```

## Getting entity information
The resoltuion results themselves do contain some information about the entities found, but to get all of the details
for that entity we need to call the entity endpoint.

A request to view the first resolved entity (best match) from the previous request would look like this:
```go
// Get the entity details for the best match
entityDetails, err := client.Entity.GetEntity(context.Background(), resolution.Data[0].EntityId, &sayari.GetEntity{})
if err != nil {
    log.Fatalf("Error: %v", err)
}
```

## Complete example
After the steps above you should be left with code looks like this. We can add one final line to print all of the fields
of the resolved entity to see what it looks like.
```go
package main

import (
	"context"
	"github.com/joho/godotenv"
	sayari "github.com/sayari-analytics/sayari-go/generated/go"
	"github.com/sayari-analytics/sayari-go/sdk"
	"log"
	"os"
)

func main() {
	// load ENV file if ENV vars are not set
	if os.Getenv("CLIENT_ID") == "" || os.Getenv("CLIENT_SECRET") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Failed to load .env file. Err: %v", err)
		}
	}

	// NOTE: To connect you most provide your client ID and client secret. To avoid accidentally checking these into git,
	// it is recommended to use ENV variables

	// Create a client to auth against the API
	client, err := sdk.Connect(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Create the request body
	resolutionRequest := sayari.Resolution{
		Name: []*string{sdk.String("Victoria Beckham")},
	}

	// Make the request and handle the error
	resolution, err := client.Resolution.Resolution(context.Background(), &resolutionRequest)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Get the entity details for the best match
	entityDetails, err := client.Entity.GetEntity(context.Background(), resolution.Data[0].EntityId, &sayari.GetEntity{})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Printf("%+v", entityDetails)
}

```

# Advanced
When interacting with the API directly, there are a few concepts that you need to handle manually. The SDK takes care of
these things for you when used properly, but it is important to understand them if you want to use the SDK properly.

## Authentication and token management
- Connecting
- Token rotation

## Pagination
- How to use pagination
- The helper functions the SDK exposes

## Rate limiting
- How rate limiting works in Sayari Graph and what the responses look like
- How the SDK handles this
- Consideration (shared client, etc)

# Tutorials
You should now have all the tools you need to start using the Sayari Graph Go SDK yourself. If you would like additional
inspiration, please consider the following use-case-specific tutorials.

## Screening

## Investigations

## Trade Analysis

# Endpoints
Again, refer people to the API docs

Add a subsection of invocation examples for each SDK function

Note: this was published from the CI...
