package server

import (
    "os"
    "log"
    "time"
    "strconv"
    "net/http"
    "regexp"
    "path"
    "path/filepath"
    "tulip/pkgs/post"
)

type Server struct {
    URI string `json:"uri"`
    Port int `json:"port"`
    RootDir string `json:"root"`
    BaseDir string `json:"base"`
    maxPost uint8 `json:"maxpost"`
    Relative bool `json:"relative"`
}

type Params map[string]string

func (ps Params) ByName(name string) string {
    if val, ok := ps[name]; ok {
        return val
    }
    return ""
}

var (
    templateDir string
    postDir string
    staticDir string
    base string
    rel bool
    all []*post.Post
    maxPost uint8
    routes []Route
)

type Route struct {
    Pattern string
    Params []string
    Controller func(http.ResponseWriter, *http.Request, Params)
}

func init() {
    routes = []Route{
        Route{
            Pattern:"sitemap_index\\\\.xml$",
            Params:[]string{"sitemap"},
            Controller: Index,
        },
        Route{
            Pattern:"([^/]+?)-sitemap([0-9]+)?\\\\.xml$",
            Params:[]string{"sitemap","sitemap_n"},
            Controller: ByCat,
        },
        Route{
            Pattern:"category/(.+?)/(feed|rss|atom)/?$",
            Params:[]string{"name","feed"},
            Controller: ByCat,
        },
        Route{
            Pattern:"category/(.+?)/page/?([0-9]{1,})/?$",
            Params:[]string{"name","paged"},
            Controller: ByCat,
        },
        Route{
            Pattern:"category/(.+?)/?$",
            Params:[]string{"name"},
            Controller: ByCat,
        },
        Route{
            Pattern:"tag/([^/]+)/(feed|rss|atom)/?$",
            Params:[]string{"name","feed"},
            Controller: ByTag,
        },
        Route{
            Pattern:"tag/([^/]+)/page/?([0-9]{1,})/?$",
            Params:[]string{"name","paged"},
            Controller: ByTag,
        },
        Route{
            Pattern:"tag/([^/]+)/?$",
            Params:[]string{"name"},
            Controller: ByTag,
        },
        Route{
            Pattern:"(feed|rss|atom)/?$",
            Params:[]string{"feed"},
            Controller: Index,
        },
        Route{
            Pattern:"page/?([0-9]{1,})/?$",
            Params:[]string{"paged"},
            Controller: Index,
        },
        Route{
            Pattern:"([0-9]{4})/([0-9]{1,2})/([0-9]{1,2})/(feed|rss|atom)/?$",
            Params:[]string{"year","monthnum","day","feed"},
            Controller: Index,
        },
        Route{
            Pattern:"([0-9]{4})/([0-9]{1,2})/([0-9]{1,2})/page/?([0-9]{1,})/?$",
            Params:[]string{"year","monthnum","day","paged"},
            Controller: Index,
        },
        Route{
            Pattern:"([0-9]{4})/([0-9]{1,2})/([0-9]{1,2})/?$",
            Params:[]string{"year","monthnum","day"},
            Controller: Index,
        },
        Route{
            Pattern:"([0-9]{4})/([0-9]{1,2})/(feed|rss|atom)/?$",
            Params:[]string{"year","monthnum","feed"},
            Controller: Index,
        },
        Route{
            Pattern:"([0-9]{4})/([0-9]{1,2})/page/?([0-9]{1,})/?$",
            Params:[]string{"year","monthnum","paged"},
            Controller: Index,
        },
        Route{
            Pattern:"([0-9]{4})/([0-9]{1,2})/?$",
            Params:[]string{"year","monthnum"},
            Controller: Index,
        },
        Route{
            Pattern:"([0-9]{4})/(feed|rss|atom)/?$",
            Params:[]string{"year","feed"},
            Controller: Index,
        },
        Route{
            Pattern:"([0-9]{4})/page/?([0-9]{1,})/?$",
            Params:[]string{"year","paged"},
            Controller: Index,
        },
        Route{
            Pattern:"([0-9]{4})/?$",
            Params:[]string{"year"},
            Controller: Index,
        },
        Route{
            Pattern:"(.?.+?)/(feed|rss|atom)/?$",
            Params:[]string{"name","feed"},
            Controller: Page,
        },
        Route{
            Pattern:"(.?.+?)(?:/([0-9]+))?/?$",
            Params:[]string{"name","page"},
            Controller: Page,
        },
        Route{
            Pattern:"([^/]+)/(feed|rss|atom)/?$",
            Params:[]string{"name","feed"},
            Controller: ReadPost,
        },
        Route{
            Pattern:"([^/]+)/page/?([0-9]{1,})/?$",
            Params:[]string{"name","paged"},
            Controller: ReadPost,
        },
        Route{
            Pattern:"([^/]+)/amp(/(.*))?/?$",
            Params:[]string{"name","amp"},
            Controller: ReadPost,
        },
        Route{
            Pattern:"([^/]+)(?:/([0-9]+))?/?$",
            Params:[]string{"name","page"},
            Controller: ReadPost,
        },
    }
}

func setupServer(s Server) {
    templateDir = filepath.Join(s.BaseDir, "template")
    postDir = filepath.Join(s.BaseDir, "_posts")
    staticDir = s.RootDir
    maxPost = s.maxPost
    initTemplate()
    if _, err := os.Stat(postDir); os.IsNotExist(err) {
        err = os.MkdirAll(postDir, os.FileMode(0775))
        if err != nil {
            log.Fatal( err )
        }
    }
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
        
        if len(r.URL.Path) == 0 || r.URL.Path == "/" {
            Index(w, r, nil)
            return
        }
        
        for _, sw := range routes {
            matched := regexp.MustCompile(sw.Pattern).FindStringSubmatch(r.URL.Path)
            if len(matched) > 1 {
                params := make(Params)
                for i := 1; i < len(matched); i++ {
                    if (i - 1) < len(sw.Params) {
                        params[sw.Params[i - 1]] = matched[i]
                    }
                }
                log.Println(sw.Pattern)
                sw.Controller(w, r, params)
                return
            }
        }
        
        regx, _ := regexp.Compile(`\.(txt|js|css|3gp|gif|jpg|jpeg|png|ico|wmv|avi|asf|asx|mpg|mpeg|mp4|pls|mp3|mid|wav|swf|flv|exe|zip|tar|rar|gz|tgz|bz2|uha|7z|doc|docx|xls|xlsx|pdf|iso|eot|svg|ttf|woff)`)
        if regx.MatchString(r.URL.Path) {
            fs := http.Dir(staticDir)
            _, err := fs.Open(path.Clean(r.URL.Path))
            if os.IsNotExist(err) {
                w.Header().Set("Content-Type", "text/html; charset=UTF-8")
                
                return
            }
            fileServer := http.FileServer(fs)
            fileServer.ServeHTTP(w, r)
            return
        }
    })

    log.Println("Listen and served on port "+strconv.Itoa(s.Port))
    log.Fatal(http.ListenAndServe(":" + strconv.Itoa(s.Port), nil))
}
