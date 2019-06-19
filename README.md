Simple Tulip blog markdown with golang
=====

That is simple blog markdown with golang. Not smart and clean code.

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
| PORT      | ":8080"     |
| BASEDIR   | $PWD        |
| POSTDIR   | "posts"     |
| TEMPLDIR  | "templates" |
| STATICDIR | "static"    |
| RELATIVE  | true        |
| MAXPOSTS  | 5           |
