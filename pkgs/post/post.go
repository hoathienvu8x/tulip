package post

import (
    "os"
    "log"
    "bytes"
    "sort"
    "path/filepath"
    "html/template"
    "io/ioutil"
    "tulip/pkgs/slug"
    "github.com/russross/blackfriday"
)

type Post struct {
    *Meta
    Slug string
    Content template.HTML
}

func New(fn string) (*Post, error) {
    file, err := os.Open(fn)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    b, err := ioutil.ReadAll(file)
    if err != nil {
        return nil, err
    }
    if !bytes.HasPrefix(b, []byte("---\n")) {
        return nil, errMissingFrontMaster
    }
    b = bytes.TrimPrefix(b, []byte("---\n"))
    arr := bytes.SplitN(b, []byte("\n---\n"), 2)
    if len(arr) == 0 {
        return nil, errMissingFrontMaster
    }
    m, err := newMeta(string(arr[0]))
    if err != nil {
        return nil, err
    }
    body := blackfriday.MarkdownCommon(arr[1])
    p := &Post{
        m,
        slug.Make(m.Title),
        template.HTML(body),
    }
    return p, nil
}

func GetPosts(postDir string) []*Post {
    paths, err := filepath.Glob(filepath.Join(postDir, "*.md"))
    if err != nil {
        log.Print(err)
    }
    all := make([]*Post, 0, len(paths))
    for _, path := range paths {
        p, err := New(path)
        if err != nil {
            log.Print(err)
        }
        all = append(all, p)
    }
    sort.Slice(all, func(i, j int) bool {
        return all[i].Meta.Date.After(all[j].Meta.Date)
    })
    return all
}
