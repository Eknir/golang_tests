package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/stretchr/gomniauth"
	gomniauthcommon "github.com/stretchr/gomniauth/common"
	"github.com/stretchr/objx"
)

type ChatUser interface {
	UniqueID() string
	AvatarURL() string
}
type chatUser struct {
	gomniauthcommon.User
	uniqueID string
}

type authHandler struct {
	next http.Handler
}

func (u chatUser) UniqueID() string {
	return u.uniqueID
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("auth"); err != nil || cookie.Value == "" {
		// not authenticated
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
	// success - call the next handler
	h.next.ServeHTTP(w, r)
}
func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}

//loginHandler handles the third-party login process.
// format: /auth/{action}/{provider}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	if len(segs) != 4 {
		http.Error(w, fmt.Sprint("No correct path given. Expected: /auth/{action}/{provider} "), http.StatusBadRequest)
		return
	}
	action := segs[2]
	provider := segs[3]
	switch action {
	case "login":
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			http.Error(
				w, fmt.Sprintf("Error when trying to get provider %s: %s", segs[3], err),
				http.StatusBadRequest)
			return
		}
		loginURL, err := provider.GetBeginAuthURL(nil, nil)
		if err != nil {
			http.Error(
				w, fmt.Sprintf("Error when trying to get GetBeginAuthURL for %s: %s", provider, err),
				http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", loginURL)
		w.WriteHeader(http.StatusTemporaryRedirect)
	case "callback":
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			http.Error(
				w, fmt.Sprintf("Error when trying to get provider %s: %s",
					provider, err), http.StatusBadRequest)
			return
		}
		creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
		if err != nil {
			http.Error(
				w, fmt.Sprintf("Error when trying to complete auth for %s: %s", provider, err),
				http.StatusInternalServerError)
			return
		}
		user, err := provider.GetUser(creds)
		if err != nil {
			log.Fatalln("Error when trying to get user from", provider, "-", err)
		}
		chatUser := &chatUser{User: user}
		m := md5.New()
		io.WriteString(m, strings.ToLower(user.Email()))
		chatUser.uniqueID = fmt.Sprintf("%x", m.Sum(nil))
		avatarURL, err := avatars.GetAvatarURL(chatUser)
		if err != nil {
			log.Fatalln("Error when trying to GetAvatarURL", "-", err)
		}

		authCookieValue := objx.New(map[string]interface{}{
			"userid":     chatUser.uniqueID,
			"name":       user.Name(),
			"avatar_url": avatarURL,
		}).MustBase64()
		http.SetCookie(w, &http.Cookie{
			Name:  "auth",
			Value: authCookieValue,
			Path:  "/",
		})
		w.Header().Set("Location", "/chat")
		w.WriteHeader(http.StatusTemporaryRedirect)
	case "guest":
		email := "MyEmailAddress@example.com"
		m := md5.New()
		io.WriteString(m, strings.ToLower(email))
		userId := fmt.Sprintf("%x", m.Sum(nil))
		authCookieValue := objx.New(map[string]interface{}{
			"userid":     userId,
			"name":       "Guest",
			"avatar_url": "https://placekitten.com/g/50/50",
			"email":      email,
		}).MustBase64()
		http.SetCookie(w, &http.Cookie{
			Name:  "auth",
			Value: authCookieValue,
			Path:  "/",
		})
		w.Header().Set("Location", "/chat")
		w.WriteHeader(http.StatusTemporaryRedirect)
	default:
		http.Error(w, fmt.Sprintf("Auth action %s not supported", action),
			http.StatusNotFound)
		return
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "auth",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	w.Header().Set("Location", "/chat")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
