package generator

import (
    "../struct"

    "os"
    "fmt"
    "net/http"

    "github.com/labstack/echo"

    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)

func Default(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        var top db_tables.Top
        db.Raw("SELECT * FROM tops ORDER BY RANDOM() LIMIT 1").Scan(&top)

        var under db_tables.Under
        db.Raw("SELECT * FROM unders ORDER BY RANDOM() LIMIT 1").Scan(&under)

        result := top.Data + under.Data

        return c.String(http.StatusOK, result)
    }
}

func Create(db *gorm.DB, new_data string) echo.HandlerFunc {
    return func(c echo.Context) error {
        data := db_tables.Top{Data: "hoge"}
        db.Create(&data);

        return c.String(http.StatusOK, "ok")
    }
}
