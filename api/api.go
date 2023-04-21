package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Emyrk/strava/api/queue"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"golang.org/x/oauth2"

	"github.com/Emyrk/strava/api/httpmw"
	"github.com/Emyrk/strava/api/webhooks"
	"github.com/Emyrk/strava/database"
)

type Options struct {
	OAuth     OAuthOptions
	DB        database.Store
	Logger    zerolog.Logger
	AccessURL *url.URL
}

type OAuthOptions struct {
	ClientID string
	Secret   string
}

type API struct {
	Opts    *Options
	Handler http.Handler

	OAuthConfig *oauth2.Config
	Events      *webhooks.ActivityEvents
	Manager     *queue.Manager
}

func New(opts Options) (*API, error) {
	api := &API{
		Opts: &opts,
		OAuthConfig: &oauth2.Config{
			ClientID:     opts.OAuth.ClientID,
			ClientSecret: opts.OAuth.Secret,
			Endpoint: oauth2.Endpoint{
				AuthURL:   "https://www.strava.com/oauth/authorize",
				TokenURL:  "https://www.strava.com/oauth/token",
				AuthStyle: 0,
			},
			RedirectURL: fmt.Sprintf("%s/oauth2/callback", opts.AccessURL.String()),
			// Must be comma joined
			Scopes: []string{strings.Join([]string{"read", "read_all", "profile:read_all", "activity:read"}, ",")},
		},
	}
	api.Events = webhooks.NewActivityEvents(opts.Logger, api.OAuthConfig, api.Opts.DB, opts.AccessURL)
	r := api.Routes()
	api.Events.Attach(r)
	api.Handler = r

	return api, nil
}

// StartWebhook needs to be called after the API is served.
func (api *API) StartWebhook(ctx context.Context) (<-chan *webhooks.WebhookEvent, error) {
	err := api.Events.Setup(ctx)
	if err != nil {
		return nil, err
	}
	return api.Events.EventQueue(), nil
}

func (api *API) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/oauth2", func(r chi.Router) {
		r.Use(httpmw.ExtractOauth2(api.OAuthConfig, nil))
		r.Get("/callback", api.stravaOAuth2)
	})

	return r
}