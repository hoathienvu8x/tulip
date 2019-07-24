package server

import (
    "fmt"
    "log"
    "math"
    "path/filepath"
    "html/template"
    "tulip/pkgs/post"
)

var (
    t *template.Template
    funcMaps template.FuncMap
)

func initTemplate() {
    funcMaps = template.FuncMap{
        "inc" : func(i uint8) uint8 { return i + 1 },
        "dec" : func(i uint8) uint8 { return i - 1 },
    }
    t = template.New("tulip").Funcs(funcMaps)
    template.Must(t.ParseGlob(filepath.Join(templateDir, "*.ghtml")))
}

type PageData struct {
    PageNow uint8
    PageMax uint8
    Posts []*post.Post
}

func newPageData(n uint8) (*PageData, error) {
    pd := new(PageData)
    pd.PageMax = uint8(math.Ceil(float64(len(all))/float64(maxPost)))
    if n < 1 || n > pd.PageMax {
        return nil, fmt.Errorf("There is no page %d", n)
    }
    pd.PageNow = n
    s := maxPost * (n - 1)
    fmt.Println(all,s)
    pd.Posts = all[s:]
    f := maxPost * n
    if f < uint8(len(all)) {
        pd.Posts = all[s:f]
    }
    return pd, nil
}

type PostData struct {
    *post.Post
    Next *post.Post
    Prev *post.Post
}

func newPostData(slug string) (*PostData, error){
    pd := new(PostData)
    for i, p := range all {
        if p.Slug == slug {
            pd.Post = p
            if i > 0 {
                pd.Prev = all[i-1]
            }
            if i < len(all) - 1 {
                pd.Next = all[i+1]
            }
            return pd, nil
        }
    }
    return nil, fmt.Errorf("Can't find post with given slug: %s", slug)
}

type ByTagData struct {
    PageNow uint8
    PageMax uint8
    Tag string
    Posts []*post.Post
}

func newByTagData(n uint8, tag string) *ByTagData {
    pd := new(ByTagData)
    pd.Tag = tag
TOUTER:
    for _, p := range all {
        for _, t := range p.Tag {
            if tag == t {
                pd.Posts = append(pd.Posts, p)
                continue TOUTER
            }
        }
    }
    if len(pd.Posts) == 0 {
        pd.PageNow = 1
        pd.PageMax = 1
        return pd
    }
    pd.PageNow = n
    pd.PageMax = uint8(math.Ceil(float64(len(pd.Posts)) / float64(maxPost)))
    s := maxPost * (n - 1)
    pd.Posts = pd.Posts[s:]
    f := maxPost * n
    if f < uint8(len(pd.Posts)) {
        pd.Posts = pd.Posts[s:f]
    }
    return pd
}

type ByCatData struct {
    PageNow uint8
    PageMax uint8
    Category string
    Posts []*post.Post
}

func newByCatData(n uint8, cat string) *ByCatData {
    pd := new(ByCatData)
    pd.Category = cat
COUTER:
    for _, p := range all {
        for _, c := range p.Category {
            if cat == c {
                pd.Posts = append(pd.Posts, p)
                continue COUTER
            }
        }
    }
    if len(pd.Posts) == 0 {
        pd.PageNow = 1
        pd.PageMax = 1
        return pd
    }
    pd.PageNow = n
    pd.PageMax = uint8(math.Ceil(float64(len(pd.Posts)) / float64(maxPost)))
    s := maxPost * (n - 1)
    pd.Posts = pd.Posts[s:]
    f := maxPost * n
    if f < uint8(len(pd.Posts)) {
        pd.Posts = pd.Posts[s:f]
    }
    return pd
}

type AboutData struct {
    *post.Post
}

func newAboutData() *AboutData {
    pd := new(AboutData)
    p, err := post.New(filepath.Join(postDir, "about"))
    if err != nil {
        log.Println("No About page data")
    }
    pd.Post = p
    return pd
}
