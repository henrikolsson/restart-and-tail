package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Please provide a systemd unit name")
		fmt.Println("Usage: restart-and-tail <unit-name>")
		os.Exit(1)
	}

	unitName := os.Args[1]

	tailCmd := exec.Command("journalctl", "-f", "-u", unitName)
	tailCmd.Stdout = os.Stdout
	tailCmd.Stderr = os.Stderr

	err := tailCmd.Start()
	if err != nil {
		fmt.Printf("Error tailing logs for unit %s: %v\n", unitName, err)
		os.Exit(1)
	}

	restartCmd := exec.Command("systemctl", "restart", unitName)
	restartCmd.Stdout = os.Stdout
	restartCmd.Stderr = os.Stderr

	err = restartCmd.Start()
	if err != nil {
		fmt.Printf("Error restarting unit %s: %v\n", unitName, err)
		os.Exit(1)
	}
	err = restartCmd.Wait()
	if err != nil {
		fmt.Printf("Command finished with error: %v", err)
	}
	err = tailCmd.Wait()
	if err != nil {
		fmt.Printf("Command finished with error: %v", err)
	}
	fmt.Println("Command finished successfully")
}

	
	
	
	