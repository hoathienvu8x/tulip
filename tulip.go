package main

import (
    "os"
    "log"
    "time"
    "strings"
    "context"
    "os/signal"
    engine "tulip/pkgs/server"
)

func main() {
    stop := make(chan os.Signal)
    signal.Notify(stop, os.Interrupt)
    app := engine.New()
    go func() {
        log.Println("Listening to port :" + strings.Trim(app.Addr,":"))
        log.Fatal(app.ListenAndServe())
    }()
    <- stop
    ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
    log.Println("Shutting down the server...")
    app.Shutdown(ctx)
    log.Println("Server stopped.")
}
//http://www.guyrutenberg.com/2008/10/02/retrieving-googles-cache-for-a-whole-website/ + https://gist.github.com/minhajuddin/1504425