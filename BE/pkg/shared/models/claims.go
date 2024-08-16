package models

import (
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserId      int      `json:"userId"`
	Name        string   `json:"name"`
	EmployeeId  string   `json:"employeeId"`
	Permissions []string `json:"permissions"`
	Locations   []string `json:"locations"`
	Equipments  []string `json:"equipments"`
	Zones       []string `json:"zones"`
	DocketTypes []string `json:"docketTypes"`
	jwt.StandardClaims
}
