package main

import (
	"fmt"
	"os/exec"
)

func main() {
	c := exec.Command("copy", "C:\\ProgramData\\chocolatey\bin\\RefreshEnv.cmd", "C:\\ProgramData\\chocolatey\bin\\RefreshEnv.installed.cmd")

	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}
