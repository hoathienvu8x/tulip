package server

import (
    "os"
    "strings"
    "strconv"  
    "net/http"
    "path/filepath"
    "tulip/pkgs/post"
    "github.com/julienschmidt/httprouter"
)

var (
    port string
    templateDir string
    postDir string
    staticDir string
    base string
    rel bool
    all []*post.Post
    maxPost uint8
)

func init() {
    ok := true
    port, ok = os.LookupEnv("TULIP_PORT")
    if !ok {
        port = "9600"
    }
    port = strings.Trim(port, ":")
    templateDir, ok = os.LookupEnv("TULIP_TEMPDIR")
    if !ok {
        templateDir = "templates"
    }
    postDir, ok = os.LookupEnv("TULIP_POSTDIR")
    if !ok {
        postDir = "posts"
    }
    staticDir, ok = os.LookupEnv("TULIP_STATICDIR")
    if !ok {
        staticDir = "static"
    }
    var err error
    rel, err = strconv.ParseBool(os.Getenv("TULIP_RELATIVE"))
    if err != nil {
        rel = false
    }
    if rel {
        base, ok = os.LookupEnv("TULIP_BASE")
        if !ok {
            base = os.Getenv("PWD")
        }
        templateDir = filepath.Join(base, templateDir)
        postDir = filepath.Join(base, postDir)
        staticDir = filepath.Join(base, staticDir)
    }
    if n, err := strconv.ParseUint(os.Getenv("TULIP_MAXPOSTS"), 10, 8); err != nil {
        maxPost = 5
    } else {
        maxPost = uint8(n)
    }
    all = post.GetPosts(postDir)
}

func newRouter() *httprouter.Router {
    r := httprouter.New()
    r.GET("/", Index)
    r.GET("/page/:page", Page)
    r.GET("/post/:title", ReadPost)
    r.GET("/about", About)
    r.GET("/tag/:name/:page", ByTag)
    r.GET("/category/:name/:page", ByCat)
    r.ServeFiles("/static/*filepath", http.Dir(staticDir))
    return r
}

func New() *http.Server {
    r := newRouter()
    return &http.Server{
        Addr : ":"+port,
        Handler : r,
    }
}
