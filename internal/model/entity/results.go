package model

import (
    "gorm.io/gorm"
)

type URL struct {
    gorm.Model 

    Name     string `json:"name" gorm:"primaryKey;unique"`
    Address  string `json:"address"`
    City     string `json:"city"`
    Country  string `json:"country"`
    Pincode  string `json:"pincode"`
    SatScore int    `json:"sat_score"`
    Passed   string `json:"passed"`
}
