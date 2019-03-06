package generator

import (
    "os"
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
    "math/rand"
    "time"

    "github.com/labstack/echo"
)

func Default(c echo.Context) (err error) {
    head, tail := GetDataset()

    rand.Seed(time.Now().UnixNano())

    head_idx := rand.Intn(len(head))
    tail_idx := rand.Intn(len(tail))

    result := head[head_idx] + tail[tail_idx]

    return c.String(http.StatusOK, result)
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
