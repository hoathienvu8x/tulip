package post

import (
    "time"
    "errors"
    "strings"
)

var (
    errInvalidFrontMatter = errors.New("Invalid front master")
    errMissingFrontMaster = errors.New("Missing front master")

    dateFmt = map[int]string{
        10 : "2006-01-02",
        15 : "2006-01-02 15:04",
        19 : "2006-01-02 15:04 WIB",
        25 : time.RFC3339,
    }
)

func UnmarshalText(t []byte) (*time.Time, error) {
    d, err := time.Parse(dateFmt[len(string(t))],string(t))
    return &d, err
}

type Meta struct {
    Title string
    Author string
    Description string
    Date time.Time
    Tag []string
    Category []string
}

func newMeta(t string) (*Meta, error) {
    t = strings.TrimSpace(t)
    if len(t) == 0 {
        return nil, errors.New("Invalid meta string")
    }
    rows := strings.Split(t,"\n")
    if len(rows) == 0 {
        return nil, errors.New("Invalid meta string")
    }
    args := make(map[string]string)
    for _, row := range rows {
        row = strings.TrimSpace(row)
        if len(row) == 0 {
            continue
        }
        val := strings.Split(row, ":")
        if len(val) == 0 {
            continue
        }
        key := strings.TrimSpace(val[0])
        value := val[1:]
        args[key] = strings.Join(value[:],":")
    }
    if _, ok := args["title"]; !ok {
        return nil, errors.New("Title meta missing")
    }
    if _, ok := args["date"]; !ok {
        return nil, errors.New("Date meta missing")
    }
    md, err := UnmarshalText([]byte(args["date"]))
    if err != nil {
        return nil, err
    }
    if md == nil {
        return nil, errors.New("Date is not invalid")
    }
    m := new(Meta)
    m.Title = strings.TrimSpace(args["title"])
    m.Date = *md
    if author, ok := args["author"]; ok {
        m.Author = author
    }
    if desc, ok := args["excerpt"]; ok {
        m.Description = desc
    }
    if tags, ok := args["tags"]; ok {
        tags = strings.TrimSpace(tags)
        if len(tags) > 0 {
            m.Tag = strings.Split(tags, ",")
        }
    }
    if categories, ok := args["categories"]; ok {
        categories = strings.TrimSpace(categories)
        if len(categories) > 0 {
            m.Category = strings.Split(categories, ",")
        }
    }
    return m, nil
}
