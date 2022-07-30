package models

import (
	"gorm.io/gorm"
)

type EmailAddress struct {
	gorm.Model
	MailboxID string  `json:"mailbox_id"`
	Mailbox   Mailbox `json:"mailbox"`
	Address   string  `json:"address"`
}

type Email struct {
	gorm.Model
	Headers     string `form:"headers"`
	DKIM        string `form:"dkim"`
	To          string `form:"to"`
	Text        string `form:"text"`
	HTML        string `form:"html"`
	From        string `form:"from"`
	SenderIp    string `form:"sender_ip"`
	SpamReport  string `form:"spam_report"`
	Envelope    string `form:"envelope"`
	Attachments string `form:"attachments"`
	Subject     string `form:"subject"`
	SpamScore   string `form:"spam_score"`
	SPF         string `form:"spf"`
}
