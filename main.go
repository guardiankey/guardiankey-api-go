package main

import (
	"fmt"
	"guardiankey-api-go/guardiankey"
)

func main() {

	conf := map[string]string{
		"organization_id": "",
		"authgroup_id":    "",
		"key":             "",
		"iv":              "",
		"service":         "go-test",
		"agentId":         "",
	}

	gk := guardiankey.NewGuardianKey(conf)

	clientIP := "127.0.0.1"
	userAgent := "Mozilla/5.0"
	username := "testuser"
	useremail := "testuser@example.com"
	loginFailed := false

	response, err := gk.CheckAccess(clientIP, userAgent, username, useremail, loginFailed)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response:", response)
}
