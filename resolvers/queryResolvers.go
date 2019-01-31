package resolvers

import (
	"context"

	"github.com/garethsharpe/gql/models"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) Case(ctx context.Context, id string) (models.Case, error) {
	var returnCase models.Case
	for _, c := range r.cases {
		if c.Id == id {
			returnCase = c
			break
		}
	}
	return returnCase, nil
}

func (r *queryResolver) Cases(ctx context.Context) ([]models.Case, error) {
	return r.cases, nil
}