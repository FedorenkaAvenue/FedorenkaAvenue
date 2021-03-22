package main

import (
	"net/http"
	"os/exec"
)

func getDockerStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	cmd := exec.Command("docker", "ps", "-a", "--format", "table {{ .ID }}\t{{ .Names }}\t{{.Status}}\t{{.Ports}}")
	stdout, err := cmd.Output()

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)

		return
	}

	w.Write(stdout)
}

func main() {
	http.HandleFunc("/docker/ps", getDockerStatus)
	http.ListenAndServe(":1991", nil)
}
