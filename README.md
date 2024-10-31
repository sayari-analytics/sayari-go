A Golang SDK for interacting with the Sayari graph API.

## Go Installation Guide

### macOS
1. **Using Homebrew**:
   ```bash
   brew install go
   ```

2. **Verify the Installation**:
   ```bash
   go version
   ```

### Windows
1. **Download the Installer**: Go to the [Go downloads page](https://golang.org/dl/) and download the `.msi` file for Windows.

2. **Run the Installer**: Double-click the downloaded `.msi` file and follow the installation steps.

3. **Verify the Installation**:
   Open Command Prompt or PowerShell and run:
   ```powershell
   go version
   ```

# Sayari API Example Scripts
## Using the SDK
Look in the 'example' directory to see how to use the SDK.  This repository contains example scripts for interacting with the Sayari API, using the `sayari-go` SDK. Each script demonstrates different API capabilities, such as entity resolution, load testing, pagination, and trade data searching.

## Prerequisites
- Go version 1.16+
- Set up environment variables (`CLIENT_ID` and `CLIENT_SECRET`) for authentication. You can create a `.env` file in the project root for this purpose.
- This file should look like this (with values updated)
  ```json
    CLIENT_ID=YOUR_CLIENT_ID_HERE
    CLIENT_SECRET=YOUR_CLIENT_SECRET_HERE
    ```

## Example Scripts
All Example Scripts are available under the folder `examples`

### 1. `hello-world.go`
This script demonstrates a basic connection to the Sayari API to resolve an entity by name. It:
- Loads client credentials from environment variables.
- Resolves a specified entity (e.g., "Victoria Beckham") and retrieves its entity details.

### 2. `load-testing.go`
This example performs load testing on the Sayari API by sending multiple search requests simultaneously:
- Uses a worker pool and rate limiting to control the frequency of requests.
- Conducts random entity searches based on generated strings, which simulates a high volume of traffic to test API response under load.

### 3. `pagination.go`
Demonstrates how to handle pagination when querying large datasets:
- Executes searches and iterates over paginated results for entities and records.
- Retrieves data in multiple pages to ensure all available results are fetched without overload.

### 4. `screening.go`
This script performs batch screening of entities based on data from a CSV file:
- Loads entities to screen from `entities_to_screen.csv`.
- Screens entities for potential risks and categorizes results as risky, non-risky, or unresolved.

### 5. `smoke-test.go`
A comprehensive smoke test script that performs several API operations:
- Lists and retrieves information on sources.
- Searches for entities and records, retrieves detailed entity data, and explores watchlist and beneficial ownership data.
- Conducts relationship traversal and shortest-path analysis for selected entities.
- Verifies various endpoints to ensure overall API functionality.

### 6. `trade-search.go`
This script explores trade-related data within the Sayari API, focusing on trade-specific searches:
- Searches for shipments, suppliers, and buyers based on a given query.
- Useful for retrieving trade and supplier information tied to specific keywords or entities.

## Running the Examples
Run each example using the following command:
```bash
go run examples/<example-file>.go
```
Replace <example-file> with the desired script name (e.g., hello-world.go).

## Notes
Ensure the .env file is configured correctly with your CLIENT_ID and CLIENT_SECRET.
Some examples may depend on specific data in the API; modify search terms as needed for testing.


## Documentation
Please see our [docs site](http://documentation.sayari.com) for more info and or to get in touch with us.
