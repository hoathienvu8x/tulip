Simple Tulip blog markdown with golang
=====

That is simple blog markdown with golang. Not smart and clean code.

## Requirements ##

```shell
go get github.com/julienschmidt/httprouter
go get github.com/russross/blackfriday
```

## Installation ##

To install this repository using command

```shell
go get github.com/hoathienvu8x/tulip && cp -rf ~/go/src/github.com/hoathienvu8x/tulip ~/go/src/tulip
```

That command will get source store to go source and copy to go src folder

To run script follow command

```shell
cd ~/go/src/tulip
export TULIP_PORT=9600
go run tulip.go
```

Open webbrower type `http://localhost:9600`

## Configuration ##

You can configure this using environment variables. You
need to add prefix `TULIP_` before the name of options.

| Option        | Description                                   | Type     |
|---------------|-----------------------------------------------|----------|
| PORT          | Port number                                   | string   |
| BASEDIR       | Base directory for your site                  | string   |
| POSTDIR       | Path to directory of your blog's posts        | string   |
| TEMPLDIR      | Path to directory of your blog's templates    | string   |
| STATICDIR     | Path to directory of your blog's static files | string   |
| RELATIVE      | If true make *DIR relative to BASEDIR         | bool     |
| MAXPOSTS      | Number of max posts in one page               | uint8    |

| Option    | Default     |
|-----------|-------------|
| PORT      | ":9600"     |
| BASEDIR   | $PWD        |
| POSTDIR   | "posts"     |
| TEMPLDIR  | "templates" |
| STATICDIR | "static"    |
| RELATIVE  | true        |
| MAXPOSTS  | 5           |

## Write a post ##

To write a post you must create `*.md` file with struct and save to posts folder


```
---
title:The title
excerpt:The excerpt
date:2019-06-18 15:02:00
author:The author
tags:the tag with ","
categories:The category with ","
---

The content will be here
```

<<<<<<< HEAD
![La Tulip](https://github.com/hoathienvu8x/tulip/raw/master/tulip.jpg "La Tulip")
=======
[La Tulip](https://github.com/hoathienvu8x/tulip/raw/master/tulip.jpg)
>>>>>>> 1e1bb72b934569d0c98eef6b9a5d3368c6a87f2e
