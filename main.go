package main

import (
    "./handler"
    "./struct"

    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"

    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    e := echo.New()

    e.Use(middleware.Recover())
    e.Use(middleware.Logger())
    e.Use(middleware.CORS())

    db, err := gorm.Open("sqlite3", "DB/main.sqlite3")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    db.AutoMigrate(&MsTNDB.Top{})
    db.AutoMigrate(&MsTNDB.Under{})

    db_tmp, err := gorm.Open("sqlite3", "DB/temp.sqlite3")
    if err != nil {
        panic("failed to connect database")
    }
    defer db_tmp.Close()

    db_tmp.AutoMigrate(&MsTNDB.Top{})
    db_tmp.AutoMigrate(&MsTNDB.Under{})

    e.GET("api/default", generator.Default(db))
    e.GET("api/json", generator.JSON(db))
    e.POST("api/create/top", generator.InsertTop(db_tmp))
    e.POST("api/create/under", generator.InsertUnder(db_tmp))

    e.Start(":1417")
}
