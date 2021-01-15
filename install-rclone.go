package main

import(
    "fmt"
    "os/exec"
)

func main() {    
    c := exec.Command("cmd", "/c", "copy", "C:\\Users\\jxj1522\\source\\repos\\temp\\rclone.exe", "C:\\Users\\jxj1522\\source\\repos\\temp\\rclone.new.exe")

    if err := c.Run(); err != nil { 
        fmt.Println("Error: ", err)
    }   
}