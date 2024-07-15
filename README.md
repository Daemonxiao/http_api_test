# Overview
This tool facilitates HTTP API testing through straightforward configuration files, allowing users to define and manage test cases effortlessly. It supports 
various HTTP methods (GET, POST, PUT, DELETE) and includes robust assertion capabilities for validating responses.
# Features
- Configuration-Driven: Define API test cases using simple configuration files.
- HTTP Method Support: Perform tests with popular HTTP methods.
# Getting Started
## Installation
```bash
# Clone the repository
git clone https://github.com/Daemonxiao/http_api_test.git
cd http_api_test

# Build the tool
go build -o api-test-tool  ./cmd/
mv api-test-tool /usr/local/bin/
```
# Usage
Create a configuration file (test_cases.yaml) defining your API test cases.
Execute tests using the tool:
```bash
api-test-tool -f test_cases.yaml
```
Configuration File Example (test_cases.yaml):
```yaml
url_prefix: "http://127.0.0.1:8080/api/v1"
headers:
  "Authorization": "Bearer key:secret"
tests:
  - name: "update parameter"
    url: "/cluster/parameters/config"
    method: "POST"
    headers:
      "Content-Type": "application/json"
    body:
      Name: "k1"
      Value: "v1"
    expected_status: 200
    expected_response:
      Success: true
    retry:
      count: 5
      delay: 100
```

| Option                    | Description                                                                                             | Required | Example                        |
|---------------------------|---------------------------------------------------------------------------------------------------------|----------|--------------------------------|
| `url_prefix`              | URL prefix                                                                                              | no       | `http://127.0.0.1:8080/api/v1` |
| `headers`                 | HTTP head                                                                                               | no       | `k1: v1`<br/> `k2: v2`         |
| `tests.name`              | test name                                                                                               | yes      | `test1`                        |
| `tests.url`               | URL. The actual url will be a combination of a url prefix and a url                                     | yes      | `/cluster/reload`              |
| `tests.method`            | HTTP method. Case sensitive                                                                             | yes      | `GET`                          |
| `tests.headers`           | HTTP head. Priority over `headers`                                                                      | no       | `k1: v1`<br/> `k2: v2`         |
| `tests.body`              | HTTP body                                                                                               | no       | `k1: v1`<br/> `k2: v2`         |
| `tests.expected_status`   | HTTP code                                                                                               | no       | `200`                          |
| `tests.expected_response` | HTTP response. You don't need to give the entire response body. Only fill in part of the concern value. | no       | `k1: v1`<br/> `k2: v2`         |
| `tests.retry.count`       | Retry times. Default: 0                                                                                 | no       | `5`                            |
| `tests.retry.delay`       | Wait time after failure. Default: 0. Unit: sec                                                          | no       | `10`                           |

For more examples, see [test_case.yaml](./example/test_case.yml)

Output:
```bash
# Test passed
Running test: update parameter
Test passed!

# Test failed
Running test: check parameter
Test failed!
```

