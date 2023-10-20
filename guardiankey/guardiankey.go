package guardiankey

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type GuardianKey struct {
	OrganizationID string
	AuthgroupID    string
	Key            string
	IV             string
	Service        string
	AgentID        string
	APIURL         string
}

func NewGuardianKey(conf map[string]string) *GuardianKey {
	return &GuardianKey{
		OrganizationID: conf["organization_id"],
		AuthgroupID:    conf["authgroup_id"],
		Key:            conf["key"],
		IV:             conf["iv"],
		Service:        conf["service"],
		AgentID:        conf["agentId"],
		APIURL:         "https://api.guardiankey.io/v2/checkaccess",
	}
}

func (gk *GuardianKey) createEvent(clientIP, userAgent, username, useremail string, loginFailed int) map[string]interface{} {
	event := map[string]interface{}{
		"generatedTime":     int64(DateNowUnix()),
		"agentId":           gk.AgentID,
		"organizationId":    gk.OrganizationID,
		"authGroupId":       gk.AuthgroupID,
		"service":           gk.Service,
		"clientIP":          clientIP,
		"clientReverse":     "",
		"userName":          username,
		"authMethod":        "",
		"loginFailed":       loginFailed,
		"userAgent":         userAgent,
		"psychometricTyped": "",
		"psychometricImage": "",
		"event_type":        "Authentication",
		"userEmail":         useremail,
	}
	return event
}

func (gk *GuardianKey) CheckAccess(clientIP, userAgent, username, useremail string, loginFailed int) (map[string]interface{}, error) {
	event := gk.createEvent(clientIP, userAgent, username, useremail, loginFailed)
	eventStr, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	hash := sha256.Sum256(append(eventStr, []byte(gk.Key+gk.IV)...))
	hashStr := fmt.Sprintf("%x", hash)

	jsonmsg := map[string]interface{}{
		"id":      gk.AuthgroupID,
		"message": string(eventStr),
		"hash":    hashStr,
	}

	content, err := json.Marshal(jsonmsg)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", gk.APIURL, bytes.NewBuffer(content))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response map[string]interface{}
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func DateNowUnix() int64 {
	return int64(time.Now().Unix())
}
