package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/webview/webview"
)

type Res struct {
	Message string `json:"message"`
}

func main() {

	html, err := os.ReadFile("./webview/index.html")
	if err != nil {
		fmt.Println(err)
	}
	// var count uint = 0
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("ssl tunnel mini tool")
	w.SetSize(200, 600, webview.HintNone)

	w.Bind("start", func(
		ip string,
		username string,
		password string,
		targetPort string,
		port string) Res {

		if connect(ip,
			username,
			password,
			targetPort,
			port) {
			return Res{Message: "connected!"}
		}
		return Res{Message: "failed!"}

	})

	w.SetHtml(string(html))
	w.Run()
}

func connect(ip string, username string, password string, targetPort string, port string) bool {
	cmd := exec.Command("bash", "-c", "ssh -N -L 127.0.0.1:"+port+":"+ip+":"+targetPort+" "+username+"@"+ip) // Change to your path
	fmt.Println(cmd)
	stdin, err := cmd.StdinPipe()

	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBuffer(nil)

	if err := cmd.Start(); err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(stdin)
	go func() {
		defer writer.Flush()
		defer stdin.Close()
		_, err = stdin.Write([]byte(password + "\n")) // Send \n to submit
		if err != nil {
			panic(err)
		}
		io.Copy(buf, stdout)
	}()

	if err := cmd.Wait(); err != nil {
		panic(err)
	}
	return true
}
