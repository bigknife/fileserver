package handler

import (
    "encoding/json"
    "errors"
    "image"
    _ "image/jpeg"
    _ "image/png"
    "mylog"
    "net/http"
    "os"
    "strings"
)

type WH struct {
    W   int
    H   int
}

// GET /pic/width-height?name=/a/b/c.jpg
func GetPicWidthHeight(w http.ResponseWriter, r *http.Request, serverConfig map[string]string) {
    name := r.FormValue("name")
    mylog.Info("get pic width/height, {name = %s}", name)

    if name == "" {
        panic(errors.New("POST /dir interface need 'name' param"))
    }

    storePath, _ := serverConfig["storePath"]

    absPath := storePath + "/" + strings.Trim(name, "/")

    f, err := os.Open(absPath)
    if err != nil {
        panic(err)
    }

    defer f.Close()

    im, _, err := image.DecodeConfig(f)
    if err != nil {
        panic(err)
    }

    wh := WH{im.Width, im.Height}

    data, err := json.Marshal(wh)
    if err != nil {
        panic(err)
    }

    w.Header()["Content-Type"] = []string{"application/json; charset=UTF-8"}
    w.WriteHeader(200)
    w.Write(data)
    mylog.Info("%s width-height = %v", name, wh)
}
