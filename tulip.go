package main

import (
    "os"
    "log"
    "time"
    "context"
    "os/signal"
    engine "tulip/pkgs/server"
)

func main() {
    stop := make(chan os.Signal)
    signal.Notify(stop, os.Interrupt)
    app := engine.New()
    go func() {
        log.Println("Listening to port :" + app.Addr)
        log.Fatal(app.ListenAndServe())
    }()
    <- stop
    ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
    log.Println("Shutting down the server...")
    app.Shutdown(ctx)
    log.Println("Server stopped.")
}
