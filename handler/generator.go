package generator

import (
    "../struct"

    "net/http"

    "github.com/labstack/echo"

    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)

func Default(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        var top MsTNDB.Top
        db.Raw("SELECT * FROM tops ORDER BY RANDOM() LIMIT 1").Scan(&top)

        var under MsTNDB.Under
        db.Raw("SELECT * FROM unders ORDER BY RANDOM() LIMIT 1").Scan(&under)

        result := top.Data + under.Data

        return c.String(http.StatusOK, result)
    }
}

func JSON(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        var top MsTNDB.Top
        db.Raw("SELECT * FROM tops ORDER BY RANDOM() LIMIT 1").Scan(&top)

        var under MsTNDB.Under
        db.Raw("SELECT * FROM unders ORDER BY RANDOM() LIMIT 1").Scan(&under)

        result := map[string]string{"data": top.Data + under.Data}

        return c.JSON(http.StatusOK, result)
    }
}

func InsertTop(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) (err error) {
        t := new(MsTNDB.Top)
        if err = c.Bind(t); err != nil {
            return err
        }

        db.Create(&t)

        return c.String(http.StatusOK, "ok")
    }
}

func InsertUnder(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) (err error) {
        u := new(MsTNDB.Under)
        if err = c.Bind(u); err != nil {
            return err
        }

        db.Create(&u)

        return c.String(http.StatusOK, "ok")
    }
}
