package qbe

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

// Target represents the architecture to generate assembly for with qbe.
type Target int

const (
	AMD64SysV Target = iota
	AMD64Apple
	ARM64
	ARM64Apple
	RV64
)

// String converts target to a string compatible with the `qbe -t` option.
func (target Target) String() string {
	return [...]string{
		AMD64SysV:  "amd64_sysv",
		AMD64Apple: "amd64_apple",
		ARM64:      "arm64",
		ARM64Apple: "arm64_apple",
		RV64:       "rv64",
	}[target]
}

// ToIL writes mod as QBE IL to w. Returns number of bytes written and if an error occured.
func (mod *Module) ToIL(w io.Writer) (int, error) {
	written := 0
	for _, typeDef := range mod.sortTypes() {
		count, err := fmt.Fprint(w, typeDef)
		written += count
		if err != nil {
			return written, err
		}
		count, err = io.WriteString(w, "\n\n")
		written += count
		if err != nil {
			return written, err
		}
	}
	for _, symbol := range mod.defOrder {
		count, err := fmt.Fprint(w, mod.definitions[symbol])
		written += count
		if err != nil {
			return written, err
		}
		count, err = io.WriteString(w, "\n\n")
		written += count
		if err != nil {
			return written, err
		}
	}
	return written, nil
}

// ToFile writes mod IL to a file named mod.name + ".ssa". Returns non nil on error.
func (mod *Module) ToFile() error {
	f, err := os.Create(mod.name + ".ssa")
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	_, err = mod.ToIL(w)
	if err != nil {
		return err
	}
	return w.Flush()
}

type processError struct {
	Name string
	*exec.ExitError
	Stderr string
}

func (e *processError) Error() string {
	return fmt.Sprintf("%v: %v\n%v", e.Name, e.ExitError, e.Stderr)
}

func waitCmd(stderr io.Reader, name string, cmd *exec.Cmd) error {
	stderrBytes, readErr := io.ReadAll(stderr)
	waitErr := cmd.Wait()
	if exitErr, ok := waitErr.(*exec.ExitError); ok && readErr == nil {
		return &processError{name, exitErr, string(stderrBytes)}
	}
	return waitErr
}

func killAndWait(cmd *exec.Cmd) {
	_ = cmd.Process.Kill()
	_ = cmd.Wait()
}

// ToAsm pipes mod IL to qbe and generates assembly for target named mod.name + ".s".
// Returns non nil on error.
func (mod *Module) ToAsm(target Target) error {
	cmd := exec.Command("qbe", "-", "-t", target.String(), "-o", mod.name+".s")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		_ = stdin.Close()
		return err
	}
	if err = cmd.Start(); err != nil {
		_ = stdin.Close()
		_ = stderr.Close()
		return err
	}
	w := bufio.NewWriter(stdin)
	if _, err = mod.ToIL(w); err != nil {
		killAndWait(cmd)
		return err
	}
	if err = w.Flush(); err != nil {
		killAndWait(cmd)
		return err
	}
	if err = stdin.Close(); err != nil {
		killAndWait(cmd)
		return err
	}
	return waitCmd(stderr, "qbe", cmd)
}

// ToObj pipes mod IL to qbe and cc and generates an object file for target named mod.name + ".o".
// Returns non nil on error.
func (mod *Module) ToObj(target Target, cflags []string) error {
	// Initialize qbe process
	qbe := exec.Command("qbe", "-", "-t", target.String())
	qbeStdin, err := qbe.StdinPipe()
	if err != nil {
		return err
	}
	qbeStdout, err := qbe.StdoutPipe()
	if err != nil {
		_ = qbeStdin.Close()
		return err
	}
	qbeStderr, err := qbe.StderrPipe()
	if err != nil {
		_ = qbeStdin.Close()
		_ = qbeStdout.Close()
		return err
	}

	// Initialize cc process
	cargs := []string{"-xassembler", "-", "-c", "-o", mod.name + ".o"}
	cargs = append(cargs, cflags...)
	cc := exec.Command("cc", cargs...)
	cc.Stdin = qbeStdout
	ccStderr, err := cc.StderrPipe()
	if err != nil {
		_ = qbeStdin.Close()
		_ = qbeStdout.Close()
		_ = qbeStderr.Close()
		return err
	}

	// Start processes
	if err = qbe.Start(); err != nil {
		_ = qbeStdin.Close()
		_ = qbeStdout.Close()
		_ = qbeStderr.Close()
		_ = ccStderr.Close()
		return err
	}
	if err = cc.Start(); err != nil {
		_ = ccStderr.Close()
		killAndWait(qbe)
		return err
	}

	// Write to pipe
	w := bufio.NewWriter(qbeStdin)
	if _, err = mod.ToIL(w); err != nil {
		killAndWait(qbe)
		killAndWait(cc)
		return err
	}
	if err = w.Flush(); err != nil {
		killAndWait(qbe)
		killAndWait(cc)
		return err
	}
	if err = qbeStdin.Close(); err != nil {
		killAndWait(qbe)
		killAndWait(cc)
		return err
	}

	// Wait processes
	if err = waitCmd(qbeStderr, "qbe", qbe); err != nil {
		killAndWait(cc)
		return err
	}
	return waitCmd(ccStderr, "cc", cc)
}
