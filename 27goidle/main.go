package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func callIdle() {
	timeQuantumStr := os.Getenv("TIME_QUANTUM")
	if timeQuantumStr == "" {
		timeQuantumStr = "60"
	}
	timeQuantum, _ := strconv.Atoi(timeQuantumStr)

	checkTimeStr := os.Getenv("IDLE_TIME_CHECK")
	if checkTimeStr == "" {
		checkTimeStr = "120000"
	}
	checkTime, _ := strconv.Atoi(checkTimeStr)

	userID := os.Getenv("USERID")
	sessionID := os.Getenv("SESSIONID")
	authCookie := os.Getenv("AUTH_ACCESS_TOKEN")

	elapsedTime := 0

	for {
		fmt.Println()
		fmt.Printf("Time quantum value in seconds: %d\n", timeQuantum)
		fmt.Printf("Check time value in seconds: %d\n", checkTime/1000)
		fmt.Printf("Time passed in seconds: %d\n", elapsedTime)

		cmd := exec.Command("xprintidle")
		output, _ := cmd.Output()
		idleTimeStr := strings.TrimSpace(string(output))
		idleTime, _ := strconv.Atoi(idleTimeStr)

		fmt.Printf("Idle time in seconds: %d\n", idleTime/1000)

		if idleTime > checkTime {
			fmt.Println("Xidle time exceeded")

			if userID == "" || sessionID == "" || authCookie == "" {
				fmt.Println("No userId or SessionId or auth_token")
				time.Sleep(time.Duration(timeQuantum) * time.Second)
				break
			}

			stopURL := "http://app.dev.ikarusnest.org:8003/aws/stop-container/"
			fmt.Println("Calling", stopURL)

			client := &http.Client{}
			req, _ := http.NewRequest("DELETE", stopURL, nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("rid", "thirdpartyemailpassword")
			req.Header.Set("Cookie", "sAccessToken="+authCookie)

			resp, err := client.Do(req)
			if err == nil && resp.StatusCode == 200 {
				fmt.Println("Container stopped successfully")
				break
			} else {
				fmt.Println("API endpoint didn't hit")
			}
		} else {
			fmt.Println("Idle time less than check time")
		}

		time.Sleep(time.Duration(timeQuantum) * time.Second)
		elapsedTime += timeQuantum
	}
}

func main() {
	callIdle()
}
