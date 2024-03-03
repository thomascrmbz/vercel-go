package onboarding_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/thomascrmbz/vercel-go/onboarding"
)

func TestTwirpServer(t *testing.T) {
	client := onboarding.NewOnboardingProtobufClient("https://vercel-go-gamma.vercel.app/", http.DefaultClient)
	res, err := client.GetClaimCode(context.Background(), &onboarding.GetClaimCodeRequest{
		Code: "dit is de code",
		Sn:   "serie nummer",
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Log(res.String())
}
