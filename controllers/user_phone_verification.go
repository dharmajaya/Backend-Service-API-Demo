package controllers

import (
	"fmt"
	"strings"
	"math/rand"
	"time"
	"net/http"
	"net/url"
	"encoding/json"
	"src/github.com/astaxie/beego"
)

type PhoneVerifyController struct {
	BaseController
}

func (this * PhoneVerifyController) Get(){
	beego.Debug("News Information ")
	this.Data["ProfileActive"] = true
}

func (this * PhoneVerifyController) Post() {
	// Set account keys & information Test Authentication Key and Test Auth key
	accountSid := "AC4f4bd297d95a870d2086afac751aab11"
	authToken := "11b9a688927f5db41727b14ebedc33e6"
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	// Create possible message bodies
	quotes := [7]string{"Testing the SMS Message through twilo."}
	this.Data["ProfileActive"] = true
	// Set up rand
	rand.Seed(time.Now().Unix())

	flash := beego.NewFlash()
	// Pack up the data for our message
	msgData := url.Values{}
	msgData.Set("To","NUMBER_TO") // destination phone number with country code
	msgData.Set("From","NUMBER_FROM") // From number from Twil account
	msgData.Set("Body",quotes[rand.Intn(len(quotes))])
	msgDataReader := *strings.NewReader(msgData.Encode())

	// Create HTTP request client
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make HTTP POST request and return message SID
	resp, _ := client.Do(req)
	if (resp.StatusCode >= 200 && resp.StatusCode < 300) {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if (err == nil) {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status);
	}
	flash.Store(&this.Controller)
	this.Redirect("/user/signin", 303)
}
