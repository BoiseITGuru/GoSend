package models

import (
	"gorm.io/gorm"
)

type Mailbox struct {
	gorm.Model
	Name           string  `json:"name"`
	EmailAddresses []Email `json:"email_addresses"`
}
