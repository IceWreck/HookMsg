package actions

import (
	"log"
	"os/exec"
)

type Script struct {
	Shell     string
	FileName  string
	ArgString string
}

func RunScript(s Script) {
	// Enable shells like python or bash
	switch s.Shell {
	case "python":
		s.Shell = "python3"
	case "bash":
		s.Shell = "/bin/bash"
	default:
		s.Shell = "/bin/bash"
	}
	location := s.FileName
	out, _ := exec.Command(s.Shell, location).CombinedOutput()
	log.Print(string(out))
}
