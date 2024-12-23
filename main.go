package main

import (
    "log"
)

func main() {
    // Inicializa el servidor con Wire
    e, err := InitializeServer()
    if err != nil {
        log.Fatalf("Error initializing server: %v", err)
    }

    // Inicia el servidor
    e.Logger.Fatal(e.Start(":8080"))
}
