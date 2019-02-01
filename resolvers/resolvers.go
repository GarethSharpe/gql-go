package resolvers

import (
	"github.com/garethsharpe/gql/generated"
)

type Resolver struct{ }

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Case() generated.CaseResolver {
	return &caseResolver{r}
}


