package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/eknir/myGo/trace"
	"github.com/stretchr/objx"

	"github.com/joho/godotenv"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/signature"
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	t.templ.Execute(w, data)
}

func init() {
	// load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}
func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse() // parse the flags

	// get github_secret and github_key
	git_s, exists_s := os.LookupEnv("GITHUB_SECRET")

	git_k, exists_k := os.LookupEnv("GITHUB_KEY")

	if !exists_k || !exists_s {
		fmt.Printf("Error: github oauth2 credentials not given in .env file key: %v, secret: %v", exists_k, exists_s)
		os.Exit(1)
	}
	grok, exists_grok := os.LookupEnv("GROK")
	if !exists_grok {
		log.Fatal("Error: set GROK in environment variable")
	}
	// setup gomniauth
	gomniauth.SetSecurityKey(signature.RandomKey(64))
	gomniauth.WithProviders(
		github.New(git_k, git_s, fmt.Sprint(grok, "/auth/callback/github")),
	)
	r := newRoom(UseGravatarAvatar)
	r.tracer = trace.New(os.Stdout)
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.Handle("/upload", &templateHandler{filename: "upload.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/uploader", uploadHandler)
	http.Handle("/room", r)
	// get the room going
	go r.run()
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
