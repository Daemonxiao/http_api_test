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
git clone https://github.com/your/repository.git
cd repository

# Build the tool
go build -o api-test-tool
```
# Usage
Create a configuration file (test_cases.yaml) defining your API test cases.
Execute tests using the tool:
```bash
./api-test-tool -config test_cases.yaml
```
Configuration File Example (test_cases.yaml)
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
