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

    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    db.AutoMigrate(&db_tables.Top{})
    db.AutoMigrate(&db_tables.Under{})

    e.GET("api/default", generator.Default(db))

    e.Start(":1417")
}
