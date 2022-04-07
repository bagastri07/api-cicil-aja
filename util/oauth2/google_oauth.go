package oauth2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/bagastri07/api-cicil-aja/api/model"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OauthService struct {
	oauth2Config     *oauth2.Config
	oauthStateString string
}

func GetNewOuathService() *OauthService {
	return &OauthService{
		oauth2Config: &oauth2.Config{
			RedirectURL:  fmt.Sprintf("%sauth/oauth-callback", os.Getenv("BASE_REDIRECT_URL")),
			ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
			ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
			Endpoint:     google.Endpoint,
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.profile",
				"https://www.googleapis.com/auth/userinfo.email",
			},
		},
		oauthStateString: "random-state",
	}
}

func (srv *OauthService) GoogleLogin(c echo.Context) error {
	url := srv.oauth2Config.AuthCodeURL(srv.oauthStateString)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (srv *OauthService) GoogleCallback(c echo.Context) error {
	content, err := srv.GetUserInfo(c.FormValue("state"), c.FormValue("code"))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	resp := &model.DataResponse{
		Data: content,
	}

	fmt.Println(content)

	return c.JSON(http.StatusOK, resp)
}

func (srv *OauthService) GetUserInfo(state string, code string) (*googlePlusProfile, error) {
	if state != srv.oauthStateString {
		return nil, errors.New("invalid oauth state")
	}
	token, err := srv.oauth2Config.Exchange(context.TODO(), code)

	if err != nil {
		return nil, err
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	content := new(googlePlusProfile)

	json.NewDecoder(response.Body).Decode(&content)

	return content, nil

}
