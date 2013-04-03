package handler

import (
    "html/template"
    "net/http"
)

func UploadPage(w http.ResponseWriter, r *http.Request, serverConfig map[string]string) {
    htmlPath, ok := serverConfig["uploadHtml"]
    if !ok {
        panic("uploadHtml is not found, upload page service unavailable!")
    }

    t, _ := template.ParseFiles(htmlPath)

    t.Execute(w, nil)
}
