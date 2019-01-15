package resolvers

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/garethsharpe/gql/models"
	"github.com/garethsharpe/gql/generated"
)

type Resolver struct{
	cases []models.Case
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateCase(ctx context.Context, caseArg models.InputCase) (string, error) {
	id := fmt.Sprintf("T%d", rand.Int())
	c := models.Case{
		Name:   caseArg.Name,
		ID:     &id,
	}
	r.cases = append(r.cases, c)
	return id, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Case(ctx context.Context, id string) (models.Case, error) {
	var returnCase models.Case
	for _, c := range r.cases {
		if *c.ID == id {
			returnCase = c
			break
		}
	}
	return returnCase, nil
}

func (r *queryResolver) Cases(ctx context.Context) ([]models.Case, error) {
	return r.cases, nil
}
