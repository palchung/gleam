package main

import (
    "thefreepress/server"
    "thefreepress/util"
    "thefreepress/controller"
    "thefreepress/db"
    "fmt"
    "os"
)


func main() {
    config := util.LoadConfig("./")
    db.LoadReids(config.Redis)
    os.Setenv("JWTAccessSecret", config.JWTAccessSecret)
    os.Setenv("JWTRefreshSecret", config.JWTRefreshSecret)
    fmt.Println(config)

    server.
    Init().
    // Route(
    //     controller.NewUserController(),
    // ).
    UserRoute("v1",
        routing.Authentication(),
    ).
    Listen(config.PORT)
}





