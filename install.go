package main

import(
    "fmt"
    "os/exec"
)

func main() {
    c := exec.Command("cmd", "/c", "move", "C:\\!to_install\\file.exe", "C:\\!to_install\\file.installed.exe")

    if err := c.Run(); err != nil { 
        fmt.Println("Error: ", err)
    }   
}