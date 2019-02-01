package resolvers

import (
	"context"

	"github.com/garethsharpe/gql/utils"
	"github.com/garethsharpe/gql/models"
	"github.com/garethsharpe/gql/api/appsvc"
)

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateCase(ctx context.Context, caseArg models.InputCase) (string, error) {
	accessToken := ctx.Value(utils.ACCESS_TOKEN).(string)
	appSvc := ctx.Value(utils.APPSVC).(api.IAppSvc)
	id, err := appSvc.CreateCase(accessToken, caseArg)
	return id, err
}