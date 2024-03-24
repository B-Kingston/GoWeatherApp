package main

import (	
  "gorm.io/gorm"
)

type Product struct {
  gorm.Model
  Username string `gorm:"primaryKey"` 
  Password string
  Address string
  WeatherDescription string
  Tempreature float64
}
