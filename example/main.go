package main

import (
	"time"

	"github.com/makki0205/sms"
)

func main() {
	sms := sms.NewSMS("key", "secret", "ap-northeast-1")
	sms.SenderID = "hoge"
	err := sms.Send(time.Now().String(), "819012345678")
	if err != nil {
		panic(err)
	}
}
