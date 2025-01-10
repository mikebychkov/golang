package main

import ("fmt"; "os")

const (
    usageMsg = "Usage: [username] [password]"
    invalidUserMsg = "Access denied for \"%s\"\n"
    invalidPassMsg = "Invalid password for \"%s\"\n"
    accessGrantedMsg = "Access granted to \"%s\"\n"
)

var users = map[string]string {
    "jack": "1888",
    "mike": "1879",
}

func main() {

    if len(os.Args) != 3 {
        fmt.Println(usageMsg)
        return
    }

    u, p := os.Args[1], os.Args[2]

    pass, exists := users[u]

    if !exists {
        fmt.Printf(invalidUserMsg, u)
        return
    }

    if pass != p {
        fmt.Printf(invalidPassMsg, u)
        return
    }

    fmt.Printf(accessGrantedMsg, u)
}

