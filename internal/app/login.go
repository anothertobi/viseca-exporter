package app

import (
	"context"
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/anothertobi/viseca-exporter/pkg/viseca"
	"github.com/urfave/cli/v2"
)

func loginCLI(ctx context.Context, cCtx *cli.Context) (*viseca.Client, error) {
	username := cCtx.String("username")
	password := cCtx.String("password")

	visecaClient, err := login(ctx, username, password)
	if err != nil {
		return nil, err
	}

	return visecaClient, nil
}

func login(ctx context.Context, username string, password string) (*viseca.Client, error) {
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	httpClient := &http.Client{
		Jar: cookieJar,
		// Don't follow redirects since the auth flow is built around redirects
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	res, err := httpClient.Get("https://one.viseca.ch/login/login")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var formToken string
	var found bool
	doc.Find(`input[name="FORM_TOKEN"]`).EachWithBreak(func(i int, s *goquery.Selection) bool {
		formToken, found = s.Attr("value")
		return !found
	})

	form := url.Values{}
	form.Add("FORM_TOKEN", formToken)
	form.Add("USERNAME", username)
	form.Add("PASSWORD", password)

	res, err = httpClient.PostForm("https://one.viseca.ch/login/login", form)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 302 {
		return nil, errors.New("login failed (no redirect response)")
	}

	err = awaitAppConfirmation(ctx, httpClient)
	if err != nil {
		return nil, err
	}

	visecaClient := viseca.NewClient(httpClient)

	return visecaClient, nil
}

func awaitAppConfirmation(ctx context.Context, httpClient *http.Client) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://one.viseca.ch/login/app-confirmation", nil)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				return err
			}
		default:
			res, err := httpClient.Do(req)
			if err != nil {
				return err
			}
			defer res.Body.Close()

			if res.StatusCode == 302 {
				return nil
			}
		}
		time.Sleep(2 * time.Second)
	}
}
