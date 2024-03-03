package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/thomascrmbz/vercel-go/onboarding"
)

var mux = chi.NewRouter()

type Service struct{}

func (*Service) SetupLocation(context.Context, *onboarding.SetupLocationRequest) (*onboarding.SetupLocationResponse, error) {
	return &onboarding.SetupLocationResponse{}, nil
}

func (*Service) InstallerSetupLocation(context.Context, *onboarding.InstallerSetupLocationRequest) (*onboarding.InstallerSetupLocationResponse, error) {
	return &onboarding.InstallerSetupLocationResponse{}, nil
}
func (*Service) ClaimLocation(context.Context, *onboarding.ClaimLocationRequest) (*onboarding.ClaimLocationResponse, error) {
	return &onboarding.ClaimLocationResponse{}, nil
}

func (*Service) SendClaimMail(context.Context, *onboarding.SendClaimMailRequest) (*onboarding.SendClaimMailResponse, error) {
	return &onboarding.SendClaimMailResponse{}, nil
}

func (*Service) GetClaimCode(ctx context.Context, req *onboarding.GetClaimCodeRequest) (*onboarding.GetClaimCodeResponse, error) {
	fmt.Println(req.Code, req.Sn)

	return &onboarding.GetClaimCodeResponse{
		ClaimCode:   "123456",
		LocationId:  "123",
		Claimed:     true,
		ClaimedBy:   nil,
		GeneratedBy: "me",
	}, nil
}

func init() {
	svc := &Service{}
	ts := onboarding.NewOnboardingServer(svc)

	mux.Use(middleware.CleanPath)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.StripSlashes)

	mux.Mount(ts.PathPrefix(), ts)

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<!doctype html>
<html>
<head>
    <title>Example Domain</title>

    <meta charset="utf-8" />
    <meta http-equiv="Content-type" content="text/html; charset=utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <script>
      window.si = window.si || function () { (window.siq = window.siq || []).push(arguments); };
    </script>
    <script defer src="/_vercel/speed-insights/script.js"></script>
    <script defer src="/_vercel/insights/script.js"></script>
    <style type="text/css">
    body {
        background-color: #f0f0f2;
        margin: 0;
        padding: 0;
        font-family: -apple-system, system-ui, BlinkMacSystemFont, "Segoe UI", "Open Sans", "Helvetica Neue", Helvetica, Arial, sans-serif;

    }
    div {
        width: 600px;
        margin: 5em auto;
        padding: 2em;
        background-color: #fdfdff;
        border-radius: 0.5em;
        box-shadow: 2px 3px 7px 2px rgba(0,0,0,0.02);
    }
    a:link, a:visited {
        color: #38488f;
        text-decoration: none;
    }
    @media (max-width: 700px) {
        div {
            margin: 0 auto;
            width: auto;
        }
    }
    </style>
</head>

<body>
<div>
    <h1>👋 Hello World!</h1>
    <p>This domain is for use in illustrative examples in documents. You may use this
    domain in literature without prior coordination or asking for permission.</p>
    <p><a href="https://www.iana.org/domains/example">More information...</a></p>
</div>
</body>
</html>`))
	})
	mux.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	fmt.Println(mux.Routes())
}

func Handler(w http.ResponseWriter, r *http.Request) {
	mux.ServeHTTP(w, r)
}
