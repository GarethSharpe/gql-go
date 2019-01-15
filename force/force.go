package force

import (
	"os"
	f "github.com/nimajalali/go-force/force"
	"github.com/subosito/gotenv"
)

func getConnection(accessToken string) (f.ForceApi, error) {
	gotenv.Load()
	connection, err := f.CreateWithAccessToken(
		os.Getenv("VERSION"),
		os.Getenv("CLIENT_ID"),
		accessToken,
		os.Getenv("INSTANCE_URL"),
	)
	return *connection, err
}