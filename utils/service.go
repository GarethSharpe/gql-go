package utils

import (
	"net/http"
	"context"

	"github.com/garethsharpe/gql/api/appsvc"
	"github.com/garethsharpe/gql/api/infra"
)

func WithAppSvc(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		appSvc := getAppSvcInstance()
		ctx := context.WithValue(r.Context(), APPSVC, appSvc)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func getAppSvcInstance() api.IAppSvc {
	caseRepo := repos.CaseRepo{}
	appSvcInstance := api.AppSvc{
		CaseRepo: caseRepo,
	}
	return &appSvcInstance
}