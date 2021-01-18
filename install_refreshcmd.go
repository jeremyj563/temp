package main

import (
	"fmt"
	"os/exec"
)

func main() {
	c := exec.Command("cmd", "/c", "copy", "C:\\ProgramData\\chocolatey\bin\\RefreshEnv.cmd", "C:\\to_install\\refreshenv.installed.cmd")

	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}
