package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/itohio/collective/backend/pkg/api/graph/generated"
	"github.com/itohio/collective/backend/pkg/api/graph/model"
)

func (r *queryResolver) Organizations(ctx context.Context, id *string) ([]*model.Organization, error) {
	user := ctx.Value("user")

	name := ""
	description := ""
	if user != nil {
		name = "fake"
		description = "user"
	}

	return []*model.Organization{
		&model.Organization{
			Name:        name,
			Description: description,
		},
	}, nil
}

func (r *queryResolver) Members(ctx context.Context, orgID *string, id *string) ([]*model.Member, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context, id *string) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Assets(ctx context.Context, orgID *string, id *string) ([]*model.Asset, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Wallets(ctx context.Context, id *string) ([]*model.Wallet, error) {
	return []*model.Wallet{
		&model.Wallet{
			ID:         "234",
			EthAddress: "eth",
		},
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
