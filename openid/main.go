package main

import (
	"net/http"
	"time"

	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
)

func main() {
	ctxBg := context.Background()
	provider, err := oidc.NewProvider(ctxBg, "https://login.microsoftonline.com//v2.0")
	if err != nil {
		// handle error
	}

	config := oauth2.Config{
		ClientID:     "",
		ClientSecret: "", Endpoint: provider.Endpoint(),
		RedirectURL: "https://localhost:8081/signin-oidc",

		// this redirect Url must be set in your Identity Provider Portal

		Scopes: []string{oidc.ScopeOpenID, "profile", "email"}, //

		// scopes which you need

	}

	oidcConfig := &oidc.Config{ClientID: "<CLIENT_ID>"}

	r := gin.Default()
	// … here is our /ping endpoint handler
	
	r.GET("/login", func(ctx *gin.Context) { state, err := randString(32)
	if err != nil {
	http.Error(ctx.Writer, "Internal error", http.StatusInternalServerError)
	return
	}
	nonce, err := randString(32) 
	if err != nil {
	http.Error(ctx.Writer, "Internal error", http.StatusInternalServerError)
	return
	}
	setCookie(ctx.Writer, ctx.Request, "state", state)
	
	setCookie(ctx.Writer, ctx.Request, "nonce", nonce)
	
	http.Redirect(ctx.Writer, ctx.Request, config.AuthCodeURL(state, oidc.Nonce(nonce)), http.StatusFound)
	
	})
	
	r.Run(":8082")
}

func setCookie(w http.ResponseWriter, r *http.Request, name, value string) {
	c := &http.Cookie{
		Name:   name,
		Value:  value,
		MaxAge: int(time.Hour.Seconds()), Secure: r.TLS != nil,
		HttpOnly: true,
	}
	http.SetCookie(w, c)
}
