package handler

import (
    "errors"
    "fmt"
    "io"
    "mime/multipart"
    "mylog"
    "net/http"
    "os"
    "strings"
)

type File struct {
    Dir      string
    Name     string
    FormFile multipart.File
}

func (f *File) save() {
    absPath := strings.TrimRight(f.Dir, "/")
    absPath = absPath + "/" + f.Name

    file, err := os.Open(absPath)
    if err != nil {
        file, err = os.Create(absPath)
    } else {
        panic(errors.New(absPath + " has existed"))
    }

    if err != nil {
        panic(err)
    }

    defer file.Close()
    io.Copy(file, f.FormFile)
    mylog.Info("上传文件 %s", absPath)
}

//上传文件
//Post /file body: file, dir=/a/b/c
func PostFile(w http.ResponseWriter, r *http.Request, serverConfig map[string]string) {
    file, header, err := r.FormFile("file")
    if err != nil {
        panic("POST /file must have a file field named file")
    }

    defer file.Close()
    dir := r.FormValue("dir")
    if dir == "" {
        panic("POST /file must have a dir field named dir to set the save path")
    }
    dir = strings.TrimRight(serverConfig["storePath"], "/") + "/" + strings.Trim(dir, "/")
    f := &File{dir, header.Filename, file}
    mylog.Info("%v", *f)
    f.save()

    fmt.Fprintf(w, "%s", "file uploaded")
}
