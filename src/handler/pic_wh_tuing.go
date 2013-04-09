package handler

import (
    "errors"
    "image"
    "image/jpeg"
    "image/png"
    "mylog"
    "net/http"
    "os"
    "resize"
    "strconv"
    "strings"
)

//调整图片分辨率，生成新的分辨率的图片文件
// PUT /pic/width-height , body : {name=/a.jpg&width=40&height=30}
// 生成文件： /a.jpg.40.30.jpg 如果该文件已经存在，则直接返回
func TuningPicWH(w http.ResponseWriter, r *http.Request, serverConfig map[string]string) {
    name, sWidth, sHeight := r.FormValue("name"), r.FormValue("width"), r.FormValue("height")
    mylog.Info("put /pic/width-height, {name = %s}", name)

    if name == "" {
        panic(errors.New("Put /pic/width-height interface need 'name' param"))
    }
    if sWidth == "" {
        panic(errors.New("Put /pic/width-height interface need 'width' param"))
    }
    if sHeight == "" {
        panic(errors.New("Put /pic/width-height interface need 'height' param"))
    }

    width, err := strconv.ParseInt(sWidth, 10, 0)
    if err != nil {
        panic(err)
    }
    height, err := strconv.ParseInt(sHeight, 10, 0)
    if err != nil {
        panic(err)
    }

    storePath, _ := serverConfig["storePath"]

    absPath := storePath + "/" + strings.Trim(name, "/")

    sufix := getSufix(name)

    tunedAbsPath := absPath + "." + sWidth + "." + sHeight + "." + sufix
    _, err = os.Open(tunedAbsPath)
    if err != nil {
        //文件不存在，进行调整，生成新的文件
        inFile, err := os.Open(absPath)
        if err != nil {
            panic(err)
        }
        defer inFile.Close()
        img, _, err := image.Decode(inFile)
        if err != nil {
            panic(err)
        }

        newImg := resize.Resample(img, image.Rect(0, 0, img.Bounds().Max.X, img.Bounds().Max.Y), int(width), int(height))
        outFile, err := os.Create(tunedAbsPath)

        if err != nil {
            panic(err)
        }
        defer outFile.Close()

        if sufix == "jpg" || sufix == "jpeg" || sufix == "JPG" || sufix == "JPEG" {
            err = jpeg.Encode(outFile, newImg, &jpeg.Options{100})
            if err != nil {
                panic(err)
            }
        } else if sufix == "png" || sufix == "PNG" {
            err = png.Encode(outFile, newImg)
            if err != nil {
                panic(err)
            }
        }

    }

}

func getSufix(str string) string {
    arr := strings.Split(str, ".")
    return arr[len(arr)-1.]
}
