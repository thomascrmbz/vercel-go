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
	// mux.Use(func(h http.Handler) http.Handler {
	// 	return ts
	// })

	mux.Use(middleware.CleanPath)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.StripSlashes)

	// mux.Handle(ts.PathPrefix(), ts)
	// mux.Group(func(r chi.Router) {
	// 	r.Handle(ts.PathPrefix(), ts)
	// })
	// mux.Route(pattern string, fn func(r chi.Router))
	mux.Mount(ts.PathPrefix(), ts)

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	mux.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	fmt.Println(mux.Routes())
}

func Handler(w http.ResponseWriter, r *http.Request) {
	mux.ServeHTTP(w, r)
}
