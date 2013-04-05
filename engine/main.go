package main

import (
	"os"
	"os/exec"
	"io"
	"fmt"
	"strings"
)

func main() {
	cmd := exec.Command("php", "-f", "../package/php/test.php")
	
	stdOut, errOut := cmd.StdoutPipe()
	if errOut != nil {
		fmt.Println(errOut)
		return
	}
	
	stdIn, errIn := cmd.StdinPipe()
	if errIn != nil {
		fmt.Println(errIn)
		return
	}
	
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return
	}
	
	s := strings.NewReader("Reader Value\n")
	
	fmt.Println("In")
	io.Copy(stdIn, s)
	
	go func() {
		fmt.Println("Out")
	    io.Copy(os.Stdout, stdOut)
	}()
	
	fmt.Println("Wait")
	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println("Closed")
}
