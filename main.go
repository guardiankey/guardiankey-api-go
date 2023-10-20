package main

import (
    "fmt"
    "meu_projeto/guardiankey" // Importe o pacote "guardiankey"
)

func main() {
    // Defina o endpoint de destino e os dados a serem enviados
    endpoint := "https://exemplo.com/api/endpoint"
    postData := `{"key": "value"}`

    // Envie a solicitação POST usando o pacote "guardiankey"
    response, err := guardiankey.SendPostRequest(endpoint, postData)
    if err != nil {
        fmt.Println("Erro ao enviar a solicitação POST:", err)
        return
    }

    // Imprima a resposta
    fmt.Println("Resposta do servidor:", response)
}
