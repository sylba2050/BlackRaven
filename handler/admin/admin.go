package admin

import (
    "../../struct"

    "net/http"
    "os"

    "github.com/labstack/echo"

    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)

func InsertTop(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) (err error) {
        t := new(MsTNDB.TopWithAuth)
        if err = c.Bind(t); err != nil {
            return err
        }

        pass := os.Getenv("PASSWORD")
        if pass == t.Pass {

            d := new(MsTNDB.Top)
            d.Data = t.Data
            db.Create(&d)

            return c.String(http.StatusOK, "ok")

        } else {
            return c.String(http.StatusUnauthorized, "unnauthorized")
        }
    }
}

func InsertUnder(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) (err error) {
        u := new(MsTNDB.UnderWithAuth)
        if err = c.Bind(u); err != nil {
            return err
        }

        pass := os.Getenv("PASSWORD")
        if pass == u.Pass {

            d := new(MsTNDB.Under)
            d.Data = u.Data
            db.Create(&d)

            return c.String(http.StatusOK, "ok")

        } else {
            return c.String(http.StatusUnauthorized, "unnauthorized")
        }
    }
}
