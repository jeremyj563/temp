package main

import(
    "fmt"
    "os/exec"
)

func main() {
    c := exec.Command("forfiles", "/P", "C:\\to_install", "/C", "cmd /c move @file .\\installed\\")

    if err := c.Run(); err != nil { 
        fmt.Println("Error: ", err)
    }   
}