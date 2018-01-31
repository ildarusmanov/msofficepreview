package main

import (
    "time"
    "context"
    "flag"
    "log"
    "os"
    "os/signal"
    "net/http"
)

var configfile = flag.String(
    "configfile",
    "./config.yml",
    "YAML config file",
)

func main() {
    log.Println("[*] Starting")

    flag.Parse()

    config, err := LoadConfigYAML(*configfile)

    if err != nil {
        log.Fatalf("[*] Can not load config file %s\n", configfile)
    }

	serviceLocator, err := BuildServiceLocator(config)

    if err != nil {
        log.Fatal("[*] Can not load services")
    }

    router := CreateRouter(serviceLocator)

    srv := &http.Server{
        Addr:    config.ServerAddr,
        Handler: router,
    }

    go func() {
        log.Printf("[x] Stop server due to %s\n", srv.ListenAndServe());
    }()

    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt)

    <- stop

    log.Println("[*] Shutdown Server ...")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        log.Fatal("[*] Server Shutdown:", err)
    }

    log.Println("[*] Server exiting")
}
