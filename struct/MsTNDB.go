package MsTNDB

import (
    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)

type Top struct {
    gorm.Model
    Data string
}

type Under struct {
    gorm.Model
    Data string
}
