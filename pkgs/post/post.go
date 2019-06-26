package post

import (
    "os"
    "log"
    "sort"
    "bytes"
    "strings"
    "regexp"
    "path/filepath"
    "html/template"
    "io/ioutil"
    "tulip/pkgs/slug"
    "github.com/PuerkitoBio/goquery"
    "github.com/russross/blackfriday"
    "github.com/sourcegraph/syntaxhighlight"
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
    /*nbody, err := replaceCodeParts(body)
    if err == nil {
        body = []byte(nbody)
    }*/
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

func replaceCodeParts(mdFile []byte) (string, error) {
    doc, err := goquery.NewDocumentFromReader(bytes.NewReader(mdFile))
    if err != nil {
        return "", err
    }
    doc.Find("code[class*=\"language-\"]").Each(func(i int, s *goquery.Selection) {
        oldCode := s.Text()
        formatted, err := syntaxhighlight.AsHTML([]byte(oldCode))
        if err != nil {
            log.Fatal(err)
        }
        s.SetHtml(string(formatted))
    })
    new, err := doc.Html()
    if err != nil {
        return "", err
    }
    new = strings.Replace(new, "<html><head></head><body>", "", 1)
    new = strings.Replace(new, "</body></html>", "", 1)
    return new, nil
}