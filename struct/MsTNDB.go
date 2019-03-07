package MsTNDB

import (
    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)

type Top struct {
    gorm.Model
    Data string `json:"data" form:"data" query:"data"`
}

type Under struct {
    gorm.Model
    Data string `json:"data" form:"data" query:"data"`
}
