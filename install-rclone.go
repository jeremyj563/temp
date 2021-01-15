package main

import(
    "fmt"
    "os/exec"
)

func main() {    
    c := exec.Command("cmd", "/c", "copy", "rclone.exe", "rclone.new.exe")

    if err := c.Run(); err != nil { 
        fmt.Println("Error: ", err)
    }   
}