package db

import (
	"time"

	"github.com/google/uuid"
)

type Organization struct {
	BaseModelSoftDelete
	Name        string `gorm:"type:varchar(255)"`
	Description string
	Type        string `gorm:"type:varchar(255)"`
	Wallets     []Wallet
	Assets      []Asset
	Members     []Member
	Keys        []AccessKey
}

type AccessKey struct {
	BaseModelSoftDelete
	OrganizationID uuid.UUID
	Name           string `gorm:"type:varchar(255)"`
	Token          []byte `gorm:"type:binary(60)"`
	ExpireAt       time.Time
}

type Asset struct {
	BaseModelSoftDelete
	OrganizationID uuid.UUID
	AccessKeyID    uuid.UUID
	Name           string `gorm:"type:varchar(255)"`
	Description    string
	Type           string `gorm:"type:varchar(255)"`
	EventCaps      []EventCapability
	ImageURL       string
	Disabled       bool
}

type EventCapability struct {
	BaseModelSoftDelete
	AssetID     uuid.UUID
	Code        int64
	Name        string `gorm:"type:varchar(255)"`
	Description string
}

type Member struct {
	BaseModelSoftDelete
	OrganizationID uuid.UUID
	User           User
	UserID         uuid.UUID
	Disabled       bool
	Roles          string
}

type Session struct {
	BaseModelSoftDelete
	Active  bool
	User    User
	Asset   Asset
	UserID  uuid.UUID
	AssetID uuid.UUID
}

type User struct {
	BaseModelSoftDelete
	Name        string `gorm:"type:varchar(255)"`
	FirstName   string `gorm:"type:varchar(255)"`
	LastName    string `gorm:"type:varchar(255)"`
	NickName    string `gorm:"type:varchar(255)"`
	Email       string `gorm:"type:varchar(255)"`
	AvatarURL   string
	Description string
	UserID      string `sql:"index"`
	Members     []Member
	Wallet      *Wallet
	WalletID    *uuid.UUID

	LoginAt     *time.Time
	ConfirmedAt *time.Time
	VerifiedAt  *time.Time
	InvitedAt   *time.Time
}

type Wallet struct {
	BaseModelSoftDelete
	OrganizationID uuid.UUID
	UserID         uuid.UUID
	Organization   Organization
	EthAddress     string
}
