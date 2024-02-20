A Golang SDK for interacting with the Sayari graph API.

## Using the SDK
Look in the 'example' directory to see how to use the SDK.
- Create a `.env` file in the root of this project containing your Sayari api credentials
    - This file should look like this (with values updated)
  ```json
    CLIENT_ID=YOUR_CLIENT_ID_HERE
    CLIENT_SECRET=YOUR_CLIENT_SECRET_HERE
    ```
- Run the example **from the same directory as the .env file**
  `go run examples/smoke-test/smoke-test.go`

## Documentation
Please see our [docs site](http://documentation.sayari.com) for more info and or to get in touch with us.
