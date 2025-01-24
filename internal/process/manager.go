package process

import (
	"bufio"
	"fmt"
	"os/exec"
	"sync"
	"syscall"
)

type ProcessInformation struct {
	Cmd    *exec.Cmd
	Name   string
	PID    int
	Status string
}

type ProcessManager struct {
	mu        sync.Mutex
	processes map[string]*ProcessInformation
}

func NewProcessManager() *ProcessManager {
	return &ProcessManager{
		processes: make(map[string]*ProcessInformation),
	}
}

func (pm *ProcessManager) StartProcess(name string, autoRestart string, command string, args ...string) (*ProcessInformation, error) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if _, exists := pm.processes[name]; exists {
		return nil, fmt.Errorf("process with name %q already exists", name)
	}
	fmt.Printf("executing command: %s %v\n", command, args)
	cmd := exec.Command(command, args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	pi := &ProcessInformation{
		Cmd:    cmd,
		Name:   name,
		PID:    cmd.Process.Pid,
		Status: "running",
	}
	pm.processes[name] = pi
	fmt.Printf("process %s stored", name)
	fmt.Println("processes: ", pm.processes)
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

	go func(pi *ProcessInformation) {
		if waitErr := pi.Cmd.Wait(); waitErr != nil {
			fmt.Printf("process %s exited with error: %v\n", pi.Name, waitErr)
		} else {
			fmt.Printf("process %s exited successfully\n", pi.Name)
		}

		pm.mu.Lock()
		pi.Status = "exited"
		pm.mu.Unlock()

	}(pi)

	return pi, nil
}

func (pm *ProcessManager) StopProcess(pi *ProcessInformation, force bool) error {
	if pi.Cmd.Process == nil {
		return fmt.Errorf("process not running")
	}

	if force {
		err := pi.Cmd.Process.Kill()
		if err != nil {
			return err
		}

		pi.Status = "stopped"
		return nil
	}

	err := pi.Cmd.Process.Signal(syscall.SIGTERM)
	if err != nil {
		return err
	}

	pi.Status = "stopped"
	return nil
}

func (pm *ProcessManager) GetProcess(name string) (*ProcessInformation, error) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	return pm.processes[name], nil
}

func (pm *ProcessManager) ListProcesses(verbose bool) (map[string]*ProcessInformation, error) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	return pm.processes, nil
}
