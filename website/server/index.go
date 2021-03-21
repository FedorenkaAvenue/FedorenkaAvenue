package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func getDockerStatus(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("docker", "ps", "-a", "--format", "table {{ .ID }}\t{{ .Names }}\t{{.Status}}\t{{.Ports}}")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	res := []byte(string(stdout))
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Write(res)
}

func main() {
	http.HandleFunc("/dockerps", getDockerStatus)
	http.ListenAndServe(":1991", nil)
}
