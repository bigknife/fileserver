package handler

import (
    "bufio"
    //"fmt"
    "io"
    "mylog"
    "net/http"
    "os"
    "strings"
)

//Get /file?name=/a/b/c/file
func GetFile(w http.ResponseWriter, r *http.Request, serverConfig map[string]string) {
    name := r.FormValue("name")
    if name == "" {
        return
    }

    absPath := strings.TrimRight(serverConfig["storePath"], "/") + "/" + strings.Trim(name, "/")
    mylog.Info("Get file, absPath = %s", absPath)

    if strings.HasSuffix(absPath, ".jpg") ||
        strings.HasSuffix(absPath, ".png") ||
        strings.HasSuffix(absPath, ".jpeg") ||
        strings.HasSuffix(absPath, ".JPG") ||
        strings.HasSuffix(absPath, ".JPEG") ||
        strings.HasSuffix(absPath, ".PNG") {
        w.Header().Add("Content-Type", "image/png")
    } else {
        w.Header().Add("Content-Type", "application/octet-stream")
        w.Header().Add("Content-Disposition", "attachment;filename="+name[strings.LastIndex(name, "/")+1:])
    }

    fi, err := os.Open(absPath)
    if err != nil {
        panic(err)
    }
    defer func() {
        if fi.Close() != nil {
            panic(err)
        }
    }()
    reader := bufio.NewReader(fi)
    io.Copy(w, reader)
}
