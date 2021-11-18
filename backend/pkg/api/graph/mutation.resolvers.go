package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/itohio/collective/backend/pkg/api/graph/generated"
	"github.com/itohio/collective/backend/pkg/api/graph/model"
	"github.com/itohio/collective/backend/pkg/auth"
	"github.com/itohio/collective/backend/pkg/db"
)

func (r *mutationResolver) CreateOrganization(ctx context.Context, input model.OrganizationInput) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateOrganization(ctx context.Context, organizationID string, input model.OrganizationInput) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteOrganization(ctx context.Context, organizationID string) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateMember(ctx context.Context, organizationID string, input model.MemberInput) (*model.Member, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMember(ctx context.Context, memberID string, input model.MemberInput) (*model.Member, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMember(ctx context.Context, organizationID *string, memberID string) (*model.Member, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateAsset(ctx context.Context, organizationID string, input model.AssetInput) (*model.Asset, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateAsset(ctx context.Context, assetID string, input model.AssetInput) (*model.Asset, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteAsset(ctx context.Context, assetID string) (*model.Asset, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	token, err := auth.GetAuthorization(ctx)
	if err != nil {
		return nil, err
	}
	userId, err := auth.GetSubject(token)
	if err != nil {
		return nil, err
	}

	var user db.User
	result := r.Orm.Where("user_id", userId).First(&user)
	if result.RowsAffected != 0 {
		return nil, fmt.Errorf("user already exists")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	user, _, err = input.ToDB()
	if err != nil {
		return nil, err
	}
	user.UserID = userId

	result = r.Orm.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	res := &model.User{}
	res.FromDB(user)

	return res, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, userID string, input model.UserInput) (*model.User, error) {
	token, err := auth.GetAuthorization(ctx)
	if err != nil {
		return nil, err
	}
	userId, err := auth.GetSubject(token)
	if err != nil {
		return nil, err
	}
	var user db.User
	result := r.Orm.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("user not found")
	}
	if userId != user.UserID {
		return nil, fmt.Errorf("user not found")
	}

	user, fields, err := input.ToDB()
	if err != nil {
		return nil, err
	}

	err = r.Orm.Model(&user).Select(fields).Updates(&user).Error
	if err != nil {
		return nil, err
	}

	res := model.User{}
	res.FromDB(user)
	return &res, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (*model.User, error) {
	token, err := auth.GetAuthorization(ctx)
	if err != nil {
		return nil, err
	}
	userId, err := auth.GetSubject(token)
	if err != nil {
		return nil, err
	}

	err = r.Orm.Where("user_id = ?", userId).Delete(&db.User{}).Error
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *mutationResolver) CreateWallet(ctx context.Context, organizationID *string, userID *string, input model.WalletInput) (*model.Wallet, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateWallet(ctx context.Context, walletID string, input model.WalletInput) (*model.Wallet, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteWallet(ctx context.Context, walletID string) (*model.Wallet, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Connect(ctx context.Context, input model.SessionInput) (*model.Session, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Disconnect(ctx context.Context, sessionID string) (*model.Session, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
