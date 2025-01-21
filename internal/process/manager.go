package process

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"sync"
	"syscall"
)

type ProcessInfo struct {
	Cmd    *exec.Cmd
	Name   string
	Status string
}

type ProcessManager struct {
	mu        sync.Mutex
	processes map[string]*ProcessInfo
}

func NewProcessManager() *ProcessManager {
	return &ProcessManager{
		processes: make(map[string]*ProcessInfo),
	}
}

func (pm *ProcessManager) StartProcess(name string, command string, args ...string) (*ProcessInfo, error) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if _, exists := pm.processes[name]; exists {
		return nil, fmt.Errorf("process with name %q already exists", name)
	}
	// fmt.Printf("executing command: %s %v\n", command, args)
	cmd := exec.Command(command, args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	stdoutReader := bufio.NewScanner(stdout)
	go func() {
		for stdoutReader.Scan() {
			line := stdoutReader.Text()
			fmt.Printf("[STDOUT] %s\n", line)
		}
		if err := stdoutReader.Err(); err != nil {
			fmt.Printf("error reading stdout: %v\n", err)
		}
	}()

	stderrReader := bufio.NewScanner(stderr)
	go func() {
		for stderrReader.Scan() {
			line := stderrReader.Text()
			fmt.Printf("[STDERR] %s\n", line)
		}
		if err := stderrReader.Err(); err != nil {
			fmt.Printf("error reading stderr: %v\n", err)
		}
	}()

	pi := &ProcessInfo{
		Cmd:    cmd,
		Name:   name,
		Status: "running",
	}

	go func(pi *ProcessInfo) {
		_ = pi.Cmd.Wait()
		pm.mu.Lock()
		defer pm.mu.Unlock()
		pi.Status = "exited"
	}(pi)

	return pi, nil
}

func StopProcess(pi *ProcessInfo) error {
	if pi.Cmd.Process == nil {
		return fmt.Errorf("process not running")
	}

	err := pi.Cmd.Process.Signal(syscall.SIGTERM)
	if err != nil {
		return err
	}

	pi.Status = "stopped"
	return nil
}
