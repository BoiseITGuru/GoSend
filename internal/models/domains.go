package models

import (
	"gorm.io/gorm"
)

type Domain struct {
	gorm.Model
	Name string `json:"name"`
}

type SubDomain struct {
	gorm.Model
	Name     string `json:"name"`
	DomainID uint   `json:"domain_id"`
	Domain   Domain `json:"domain"`
}

func (d *Domain) RegisterInboundHook() {

}

func (d *Domain) SetupMX() {

}

func (sd *SubDomain) RegisterInboundHook() {

}

func (sd *SubDomain) SetupMX() {

}
