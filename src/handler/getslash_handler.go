package handler

import (
    "fmt"
    "log"
    "net/http"
)

//handle "GET /" 请求
func IamAlive(w http.ResponseWriter, r *http.Request, serverConfig map[string]string) {
    log.Println("alive")
    fmt.Fprintf(w, "I'm Alive")

}
