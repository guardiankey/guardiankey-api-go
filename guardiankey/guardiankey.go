package guardiankey

import (
    "io/ioutil"
    "net/http"
    "strings"
)

// Função para enviar uma solicitação POST para um endpoint
func SendPostRequest(endpoint string, data string) (string, error) {
    // Prepare o corpo da solicitação
    payload := strings.NewReader(data)

    // Faça a solicitação POST
    resp, err := http.Post(endpoint, "application/json", payload)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    // Leia a resposta
    responseBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    return string(responseBody), nil
}
