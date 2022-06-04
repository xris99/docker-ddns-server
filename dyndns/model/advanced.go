package model

import (
	"github.com/jinzhu/gorm"
)

// CName is a dns cname entry.
type Advanced struct {
	gorm.Model
	Hostname string `gorm:"not null" form:"hostname" validate:"required,hostname"`
	Value   string   `gorm:"not null" validate:"required,hostname"`
	TargetID uint
	Ttl      int `form:"ttl" validate:"required,min=20,max=86400"`
}
