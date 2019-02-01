package resolvers

import (
	"context"

	"github.com/garethsharpe/gql/utils"
	"github.com/garethsharpe/gql/models"
	"github.com/garethsharpe/gql/api/appsvc"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) Case(ctx context.Context, id string) (models.Case, error) {
	accessToken := ctx.Value(utils.ACCESS_TOKEN).(string)
	appSvc := ctx.Value(utils.APPSVC).(api.IAppSvc)
	returnCase, err := appSvc.GetCase(accessToken, id)
	return returnCase, err
}

func (r *queryResolver) Cases(ctx context.Context) ([]models.Case, error) {
	accessToken := ctx.Value(utils.ACCESS_TOKEN).(string)
	appSvc := ctx.Value(utils.APPSVC).(api.IAppSvc)
	cases, err := appSvc.GetCases(accessToken)
	return cases, err
}