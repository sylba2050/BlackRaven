package generator

import (
    "os"
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"

    "github.com/labstack/echo"
)

func Default(c echo.Context) (err error) {
    head, tail := GetDataset()
    fmt.Fprintln(os.Stderr, head)
    fmt.Fprintln(os.Stderr, head[0])
    fmt.Fprintln(os.Stderr, tail)
    return c.String(http.StatusOK, "Hello, World!")
}

func GetDataset() ([]string, []string){
    head, err := ioutil.ReadFile("dataset/head.txt")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return nil,nil
    }

    tail, err := ioutil.ReadFile("dataset/tail.txt")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return nil,nil
    }

    head_spl := strings.Split(string(head), "\n")
    tail_spl := strings.Split(string(tail), "\n")

    return head_spl, tail_spl
}
