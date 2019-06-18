package slug

import (
    "html"
    "strings"
    "regexp"
)

func Make(title string) string {
    title = html.UnescapeString(title)
    slug := strings.TrimSpace(title)
    slug = regexp.MustCompile("à|á|ạ|ả|ã|â|ầ|ấ|ậ|ẩ|ẫ|ă|ằ|ắ|ặ|ẳ|ẵ|À|Á|Ạ|Ả|Ã|Â|Ầ|Ấ|Ậ|Ẩ|Ẫ|Ă|Ằ|Ắ|Ặ|Ẳ|Ẵ").ReplaceAllString(slug, "a")
    slug = regexp.MustCompile("è|é|ẹ|ẻ|ẽ|ê|ề|ế|ệ|ể|ễ|È|É|Ẹ|Ẻ|Ẽ|Ê|Ề|Ế|Ệ|Ể|Ễ").ReplaceAllString(slug, "e")
    slug = regexp.MustCompile("ì|í|ị|ỉ|ĩ|Ì|Í|Ị|Ỉ|Ĩ").ReplaceAllString(slug, "i")
    slug = regexp.MustCompile("ò|ó|ọ|ỏ|õ|ô|ồ|ố|ộ|ổ|ỗ|ơ|ờ|ớ|ợ|ở|ỡ|Ò|Ó|Ọ|Ỏ|Õ|Ô|Ồ|Ố|Ộ|Ổ|Ỗ|Ơ|Ờ|Ớ|Ợ|Ở|Ỡ").ReplaceAllString(slug, "o")
    slug = regexp.MustCompile("ù|ú|ụ|ủ|ũ|ư|ừ|ứ|ự|ử|ữ|Ù|Ú|Ụ|Ủ|Ũ|Ư|Ừ|Ứ|Ự|Ử|Ữ").ReplaceAllString(slug, "u")
    slug = regexp.MustCompile("ỳ|ý|ỵ|ỷ|ỹ|Ỳ|Ý|Ỵ|Ỷ|Ỹ").ReplaceAllString(slug, "y")
    slug = regexp.MustCompile("đ|Đ").ReplaceAllString(slug, "d")
    slug = regexp.MustCompile(`\W+`).ReplaceAllString(slug, " ")
    slug = strings.TrimSpace(slug)
    slug = regexp.MustCompile(`\s`).ReplaceAllString(slug, "-")
    slug = strings.ToLower(slug)
    slug = regexp.MustCompile("[^a-z0-9-_]").ReplaceAllString(slug, "-")
    slug = regexp.MustCompile("-+").ReplaceAllString(slug, "-")
    slug = strings.Trim(slug, "-_")
    return slug
}
