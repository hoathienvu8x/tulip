package server

import (
    "log"
    "net/http"
    "strconv"
    "github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    data, err := newPageData(1)
    if err != nil {
        http.NotFound(w, r)
    }
    err = t.ExecuteTemplate(w, "index", data)
    if err != nil {
        log.Fatal(err)
    }
}

func ReadPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    data, err := newPostData(ps.ByName("title"))
    if err != nil {
        http.NotFound(w, r)
    }
    err = t.ExecuteTemplate(w, "post", data)
    if err != nil {
        log.Fatal(err)
    }
}

func Page(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    p, err := strconv.ParseUint(ps.ByName("page"), 10,8)
    if err != nil {
        http.NotFound(w, r)
    }
    data, err := newPageData(uint8(p))
    if err != nil {
        http.NotFound(w, r)
    }
    err = t.ExecuteTemplate(w, "index", data)
    if err != nil {
        log.Fatal(err)
    }
}

func About(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    data := newAboutData()
    err := t.ExecuteTemplate(w, "about", data)
    if err != nil {
        log.Fatal(err)
    }
}

func ByTag(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    p, err := strconv.ParseUint(ps.ByName("page"), 10,8)
    if err != nil {
        p = 1
    }
    data := newByTagData(uint8(p), ps.ByName("name"))
    err = t.ExecuteTemplate(w, "tag", data)
    if err != nil {
        log.Fatal(err)
    }
}
