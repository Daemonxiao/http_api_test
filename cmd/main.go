package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"http_test_cli/pkg/compare"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	defaultExitCode      = 0
	defaultExitErrorCode = 1
)

type Config struct {
	URLPrefix string            `yaml:"url_prefix"`
	Headers   map[string]string `yaml:"headers"`
	Tests     []TestCase        `yaml:"tests"`
}

type TestCase struct {
	Name             string                 `yaml:"name"`
	URL              string                 `yaml:"url"`
	Method           string                 `yaml:"method"`
	Headers          map[string]string      `yaml:"headers"`
	Body             interface{}            `yaml:"body"`
	Retry            *Retry                 `yaml:"retry"`
	ExpectedResponse map[string]interface{} `yaml:"expected_response"`
	ExpectedStatus   int                    `yaml:"expected_status"`
}

type Retry struct {
	Count int `yaml:"count"`
	Delay int `yaml:"delay"`
}

func main() {
	filePath := flag.String("f", "", "file path")
	help := flag.String("h", "", "help info")
	flag.Parse()

	if *help != "" {
		printMsgAndExit("http_test_cli -f <config_file_path>", defaultExitCode)
	}

	if *filePath == "" {
		printErrorAndExit(fmt.Errorf("-f <config_file> is required"), defaultExitErrorCode)
	}

	yamlFile, err := os.ReadFile(*filePath)
	if err != nil {
		printErrorAndExit(err, defaultExitErrorCode)
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		printErrorAndExit(err, defaultExitErrorCode)
	}

	runTests(config)
}

func runTests(config Config) {
	// Execute each test case one by one
	for _, test := range config.Tests {
		printMsg("Running test: " + test.Name)
		// Make HTTP request
		test.URL = config.URLPrefix + test.URL
		tmpHeaders := config.Headers
		for key, value := range test.Headers {
			tmpHeaders[key] = value
		}
		test.Headers = tmpHeaders

		// init Retry
		if test.Retry == nil {
			test.Retry = &Retry{
				Count: 1,
			}
		}
		var err error
		for i := 0; i < test.Retry.Count; i++ {
			resp := doHttpReq(test)
			// Check response
			err = check(test, resp)
			if err == nil {
				break
			}
			printMsg("Test failed! Retrying after " + fmt.Sprintf("%d", test.Retry.Delay) + " seconds...")
			<-time.After(time.Duration(test.Retry.Delay) * time.Second)
		}
		if err != nil {
			printErrorAndExit(err, defaultExitErrorCode)
		}
		printMsg("Test passed!")
	}
}

func doHttpReq(test TestCase) []byte {
	// Prepare request
	var req *http.Request
	var err error
	if test.Method == "POST" {
		bodyBytes, err := json.Marshal(test.Body)
		if err != nil {
			printErrorAndExit(fmt.Errorf("error encoding request body: %v", err), defaultExitErrorCode)
		}
		req, err = http.NewRequest(test.Method, test.URL, bytes.NewBuffer(bodyBytes))
	} else {
		req, err = http.NewRequest(test.Method, test.URL, nil)
	}
	if err != nil {
		printErrorAndExit(fmt.Errorf("error creating request: %v", err), defaultExitErrorCode)
	}

	// Set headers
	for key, value := range test.Headers {
		req.Header.Set(key, value)
	}

	// Perform request
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		printErrorAndExit(fmt.Errorf("error sending request: %v", err), defaultExitErrorCode)
	}

	// Check response status
	if test.ExpectedStatus != 0 && resp.StatusCode != test.ExpectedStatus {
		printErrorAndExit(fmt.Errorf("expected status %d but got %d", test.ExpectedStatus, resp.StatusCode), defaultExitErrorCode)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		printErrorAndExit(fmt.Errorf("error reading response body: %s, error: %v", string(body), err), defaultExitErrorCode)
	}

	return body
}

func check(test TestCase, resp []byte) error {

	expectedResp := test.ExpectedResponse

	// Compare actual response with expected response
	var actualResp map[string]interface{}
	err := json.Unmarshal(resp, &actualResp)
	if err != nil {
		return fmt.Errorf("error parsing actual response JSON, response: %s, error: %v", string(resp), err)
	}

	if !compare.ContainsMap(actualResp, expectedResp) {
		expectedRespJson, _ := json.Marshal(expectedResp)
		actualRespJson, _ := json.Marshal(actualResp)
		return fmt.Errorf("Expected: %s\nActual: %s\n", expectedRespJson, actualRespJson)
	}
	return nil
}

func printErrorAndExit(err error, code int) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(code)
}

func printMsgAndExit(msg string, code int) {
	fmt.Fprintln(os.Stdout, msg)
	os.Exit(code)
}

func printMsg(msg string) {
	fmt.Fprintln(os.Stdout, msg)
}
