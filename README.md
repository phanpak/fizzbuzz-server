# FizzBuzz Server
A simple HTTP server that responds to requests with a FizzBuzz sequence and tracks the most frequent requests.

## Requirements
* Go 1.16 or higher

## Installation
* Clone the repository:
```shell
git clone https://github.com/<username>/<repository>.git
```
Install the dependencies:
```shell
go mod download
```
Build the server:
```shell
go build -o fizzbuzz-server .
```

## Usage
* Start the server:
```shell
./fizzbuzz-server
```

* Send a FizzBuzz request:
```shell
curl --location --request POST 'http://localhost:8080/fizzbuzz' \
--header 'Content-Type: application/json' \
--data-raw '{
    "int1": 3,
    "int2": 5,
    "limit": 15,
    "str1": "fizz",
    "str2": "buzz"
}'
```

* Get the most frequent FizzBuzz request:
```shell
curl --location --request GET 'http://localhost:8080/stats'
```

## Endpoints
### `/fizzbuzz` 
Method: POST

Description: Generates a FizzBuzz sequence based on the provided parameters.

Request body:
```json
{
    "int1": 3,
    "int2": 5,
    "limit": 15,
    "str1": "fizz",
    "str2": "buzz"
}
```
Response:

```json
{
    "result": [
        "1",
        "2",
        "fizz",
        "4",
        "buzz",
        "fizz",
        "7",
        "8",
        "fizz",
        "buzz",
        "11",
        "fizz",
        "13",
        "14",
        "fizzbuzz"
    ]
}
```

### `/stats`
Method: GET

Description: Retrieves the most frequent FizzBuzz request.

Response:

```json
{
    "hits": 5,
    "parameters": {"int1":3,"int2":5,"limit":15,"str1":"fizz","str2":"buzz"}
}

```

## Improvements

* Dockerise the application. It's useful to ensure consistency across different deployments and makes it easier to manage an application.
* Continuous Integration and Deployment (CI/CD): Implement a CI/CD pipeline to automate the building, testing, and deployment of the application. CI/CD helps ensure consistency, reliability, and speed of deployment, and allows for continuous delivery of new features and bug fixes.
* Authentication: Adding authentication would ensure that only authorized users can access the API, providing security and limiting the risk of unauthorized access or misuse. This can be achieved with tools like bearer tokens, OAuth2 or JWT.
* Rate Limiting: This would restrict the number of requests a user can make within a certain time frame, preventing overload or abuse of the API.
* Timeouts: if a request takes too long, the server should abort and return an error to the client.
* API Documentation: A clear and well-written documentation would help users understand the API's capabilities, inputs, outputs, and limitations, making it easier to use and integrate. Some tools such as OpenAPI or Swagger are useful to let developers explore an API's capabilities and automatically generate clients for such API.
* Caching: Implementing caching for frequently accessed data can help reduce the number of requests and improve performance of the API.
* Logging and Monitoring: Recording and analyzing logs and metrics would provide insight into the usage, performance, and health of the API, allowing for quick identification and resolution of issues. This can be achieved by using a tool like Prometheus or Grafana.
* The current implementation of the hit counter in the application is a simple in-memory map, which means that the hit counts are stored in the application's memory. While this is a straightforward solution, it has some limitations, such as:
  - The data is not persistent: If the application is restarted or crashes, all the hit counts are lost.
  - The data is not shareable: If the application is scaled horizontally, each instance of the application will have its own separate hit counts, making it challenging to get an accurate view of the overall usage of the service.
* Unit testing: increase coverage of the units tests
* Integration Testing: Apart from unit tests, you should also perform integration tests to check the behavior of the application as a whole.
* Security matters: TLS, security headers.
* Add some development tools such as linters

## Shortcomings

* No graceful shutdown
* Request are grouped by raw json for the statistics. This means that the order matters and extra parameters make requests not group with each other.
* Error messages used with http.Error is not properly sanitized. There's a risk to leak internal/sensitive data to the client (e.g. stacktrace or other implementation details)
