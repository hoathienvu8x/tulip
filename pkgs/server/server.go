package server

import (
    "log"
    "time"
    "strconv"
    "net/http"
    "path/filepath"
    "tulip/pkgs/post"
)

type Server struct {
    URI string
    Port int
    RootDir string
    BaseDir string
    maxPost uint8
    Relative bool
}

var (
    templateDir string
    postDir string
    staticDir string
    base string
    rel bool
    all []*post.Post
    maxPost uint8
)

func setupServer(s Server) {
    templateDir = filepath.Join(s.BaseDir, "tempalte")
    postDir = filepath.Join(s.BaseDir, "_posts")
    staticDir = s.RootDir
    maxPost = s.maxPost
    initTemplate()
}

func Run(s Server) {
    setupServer(s)
    go func() {
        all = post.GetPosts(postDir)
    }()
    http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
	    w.Header().Set("Cache-Control", "no-cache, must-revalidate, max-age=0")
	    w.Header().Set("Expires", time.Unix(0, 0).Format(http.TimeFormat))
	    w.Header().Set("Pragma", "no-cache")
	    w.Header().Set("X-Accel-Expires", "0")
	    w.Header().Set("Vary", "Origin")
        
    })

    log.Println("Listen and served on port "+strconv.Itoa(s.Port))
    log.Fatal(http.ListenAndServe(":" + strconv.Itoa(s.Port), nil))
}
