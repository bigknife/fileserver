package handler

import (
    "errors"
    "mylog"
    "net/http"
    "os"
    "strings"
)

//handle "POST /dir | body {name=/a/b/c&recursion=true}" 请求
func PostDir(w http.ResponseWriter, r *http.Request, serverConfig map[string]string) {
    name, recursion := r.FormValue("name"), r.FormValue("recursion")
    mylog.Info("post dir, {name = %s, recursion = %s}", name, recursion)

    if name == "" {
        panic(errors.New("POST /dir interface need 'name' param"))
    }

    if !(recursion == "true" || recursion == "false") {
        panic(errors.New("POST /dir interface need 'recursion' param, and the value must be 'true' or 'false'"))
    }

    storePath, _ := serverConfig["storePath"]

    absPath := storePath + "/" + strings.Trim(name, "/")

    mylog.Info("POST /dir make dir of absPath is %s", absPath)

    if recursion == "false" {
        err := os.Mkdir(absPath, 0777)
        check(err)
    } else if recursion == "true" {
        err := os.MkdirAll(absPath, 0777)
        check(err)
    }
}

func check(err error) {
    if err != nil {
        panic(err)
    }
}
