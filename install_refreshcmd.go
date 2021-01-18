package main

import (
	"fmt"
	"os/exec"
)

func main() {
	c := exec.Command("C:\\ProgramData\\chocolatey\bin\\RefreshEnv.cmd")

	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}
