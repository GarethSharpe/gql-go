package resolvers

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/garethsharpe/gql/generated"
	"github.com/garethsharpe/gql/models"
)

type Resolver struct {
	accounts []models.Account
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateAccount(ctx context.Context, account models.InputAccount) (string, error) {
	id := fmt.Sprintf("T%d", rand.Int())
	newAccount := models.Account{
		Name: account.Name,
		Id:   &id,
	}
	r.accounts = append(r.accounts, newAccount)
	return id, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Account(ctx context.Context, id string) (models.Account, error) {
	var account models.Account
	for i := range r.accounts {
		if *r.accounts[i].Id == id {
			account = r.accounts[i]
			break
		}
	}
	return account, nil
}
func (r *queryResolver) Accounts(ctx context.Context) ([]models.Account, error) {
	return r.accounts, nil
}
