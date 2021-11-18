package model

import (
	"strings"

	"github.com/google/uuid"
	"github.com/itohio/collective/backend/pkg/db"
)

func (u *UserInput) ToDB() (db.User, []string, error) {
	var id *uuid.UUID
	if u.WalletID != nil {
		uuid, err := uuid.Parse(*u.WalletID)
		if err != nil {
			return db.User{}, nil, err
		}
		id = &uuid
	}

	fields := strings.Split("Name FirstName LastName NickName Email AvatarURL WalletID", " ")
	if u.WalletID == nil {
		fields = fields[:len(fields)-1]
	}

	return db.User{
		Name:      u.Name,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		NickName:  u.NickName,
		Email:     u.Email,
		AvatarURL: u.AvatarURL,
		WalletID:  id,
	}, fields, nil
}

func (u *User) FromDB(d db.User) {
	u.ID = d.ID.String()
	u.Name = d.Name
	u.FirstName = d.FirstName
	u.LastName = d.LastName
	u.NickName = d.NickName
	u.Email = d.Email
	u.AvatarURL = d.AvatarURL
	u.Wallet = &Wallet{
		ID:         d.Wallet.ID.String(),
		EthAddress: d.Wallet.EthAddress,
	}
}
