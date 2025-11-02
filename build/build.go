package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	os.MkdirAll("bin", 0755)
	cmd := exec.Command("go", "build", "-o", "bin/chrise", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println("Building into bin/chrise")
	if err := cmd.Run(); err != nil {
		fmt.Println("Build failed", err)
		os.Exit(1)
	}
	fmt.Println("Done")
}
