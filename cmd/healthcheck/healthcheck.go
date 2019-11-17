package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://localhost:8081/debug/health")
	if err != nil {
		fmt.Printf("[FAIL] could get call health endpoint")
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("[FAIL] could not read response body: %v", err)
			os.Exit(1)
		}

		var objmap map[string]string
		err = json.Unmarshal(body, &objmap)
		if err != nil {
			fmt.Printf("[FAIL] could not parse response body: %v", err)
			os.Exit(1)
		}

		for k, v := range objmap {
			fmt.Printf("[FAIL] %s health check failed: %s", k, v)
		}
		os.Exit(1)
	}
	fmt.Printf("[PASS] All checks healthy")
}
