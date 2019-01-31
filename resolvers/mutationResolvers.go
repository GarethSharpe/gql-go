package resolvers

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/garethsharpe/gql/models"
)

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateCase(ctx context.Context, caseArg models.InputCase) (string, error) {
	id := fmt.Sprintf("T%d", rand.Int())
	c := models.Case{
		Name:   *caseArg.Name,
		Id:     id,
	}
	r.cases = append(r.cases, c)
	return id, nil
}