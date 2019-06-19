package post

import (
    "os"
    "log"
    "sort"
    "bytes"
    "regexp"
    "path/filepath"
    "html/template"
    "io/ioutil"
    "tulip/pkgs/slug"
    "github.com/russross/blackfriday"
)

type Post struct {
    *Meta
    Image *string
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
    if len(arr) == 0 || len(arr) < 2 {
        return nil, errMissingFrontMaster
    }
    m, err := newMeta(string(arr[0]))
    if err != nil {
        return nil, err
    }
    body := blackfriday.MarkdownCommon(arr[1])
    htmlBody := template.HTML(body)
    p := &Post{
        m,
        findFirstImag(htmlBody),
        slug.Make(m.Title),
        htmlBody,
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
        if p != nil {
            all = append(all, p)
        }
    }
    sort.Slice(all, func(i, j int) bool {
        return all[i].Meta.Date.After(all[j].Meta.Date)
    })
    return all
}
// https://stackoverflow.com/a/36966144
func findFirstImag(html template.HTML) *string {
    if len(html) == 0 {
        return nil
    }
    re := regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)
    imgTags := re.FindAllStringSubmatch(string(html),1) // https://stackoverflow.com/a/44847651
    if len(imgTags) == 0 {
        return nil
    }
    img := imgTags[0][1]
    return &img
}