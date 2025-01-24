package process

import (
	"bufio"
	"fmt"
	"os/exec"
	"sync"
	"syscall"
	"time"
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

	pi := &ProcessInformation{
		Name:   name,
		Status: "starting",
	}

	pm.processes[name] = pi
	runOnce := func() error {
		cmd := exec.Command(command, args...)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return err
		}
		stderr, err := cmd.StderrPipe()
		if err != nil {
			return err
		}

		if err := cmd.Start(); err != nil {
			return err
		}

		pm.mu.Lock()
		pi.PID = cmd.Process.Pid
		pi.Status = "running"
		pm.mu.Unlock()

		go func() {
			scanner := bufio.NewScanner(stdout)
			for scanner.Scan() {
				fmt.Printf("[STDOUT] (%s) %s\n", name, scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				fmt.Printf("error reading stdout: %v\n", err)
			}
		}()
		go func() {
			scanner := bufio.NewScanner(stderr)
			for scanner.Scan() {
				fmt.Printf("[STDERR] (%s) %s\n", name, scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				fmt.Printf("error reading stderr: %v\n", err)
			}
		}()

		waitErr := cmd.Wait()

		pm.mu.Lock()
		pi.Status = "exited"
		pm.mu.Unlock()

		if waitErr != nil {
			fmt.Printf("process %s exited with error: %v\n", name, waitErr)
		} else {
			fmt.Printf("process %s exited successfully\n", name)
		}
		return waitErr
	}

	switch autoRestart {
	case "always":
		go func() {
			for {
				_ = runOnce()
				time.Sleep(1 * time.Second) // optional delay
			}
		}()

	case "on-failure":
		go func() {
			for {
				waitErr := runOnce()
				if waitErr == nil {
					break
				}
				time.Sleep(1 * time.Second) // optional delay
			}
		}()

	case "never", "":
		go func() {
			_ = runOnce()
		}()

	default:
		fmt.Printf("unrecognized auto-restart policy: %q (defaulting to never)\n", autoRestart)
		go func() {
			_ = runOnce()
		}()
	}

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
