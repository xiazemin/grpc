package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {

	stop := false
	go func() {
		//exec.Command("whoami").Run()
		//cmd := exec.Command("/bin/sh", "-c", `/sbin/ifconfig en0 | grep -E 'inet ' |  awk '{print $2}'`)
		//cmd.Start()
		//cmd.Stdout = os.Stdout
		//cmd.Run()
		for !stop {
			cmd := exec.Command("bash", "-c",
				"a='#';for i in `seq 1 100`;do a=$a#;echo -en '['$a ']%\r';sleep 0.1;done")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}

	}()
	go func() {
		time.Sleep(1000000)
		fmt.Println("go routine")
	}()
	fmt.Println("main")
	time.Sleep(1000000)
	stop = true
	fmt.Println("done")
}
