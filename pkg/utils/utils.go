package utils

import (
    "io"
    "os"
    "os/exec"
    "runtime"
)

// 文件是否存在
func FileExists(path string) bool {
    _, err := os.Stat(path)

    return err == nil || os.IsExist(err)
}

// 获取数据
func FileRead(path string) (string, error) {
    file, err := os.Open(path)
    if err != nil {
        return "", err
    }
    defer file.Close()

    data, err2 := io.ReadAll(file)
    if err2 != nil {
        return "", err2
    }

    return string(data), nil
}

//cmd 打开默认浏览器
func OpenUrl(uri string) error  {
    sys := runtime.GOOS
    if sys =="windows"{
        cmd := exec.Command("cmd", "/C", "start "+uri)
        return cmd.Run()
    }else if sys =="darwin"{
        return exec.Command("open", uri).Start()
    }
    return nil
}
