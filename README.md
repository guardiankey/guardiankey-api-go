# guardiankey-api-go


1) git clone https://github.com/guardiankey/guardiankey-api-go.git
2) cd guardiankey-api-go
3) add your variables in main.go, line 10...

```
    conf := map[string]string{
        "organization_id": "",
        "authgroup_id":    "",
        "key":             "",
        "iv":              "",
        "service":         "go-test",
        "agentId":         "",
    }
```
4) go run main.go
