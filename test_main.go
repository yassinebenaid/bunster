package main

import "os/exec"

func main() {

	cmd_1_name := "cmd"
	cmd_1 := exec.Command(cmd_1_name)
	if err := cmd_1.Run(); err != nil {
		panic(err)
	}

}
