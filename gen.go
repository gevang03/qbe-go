package qbe

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

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
	switch target {
	case AMD64Apple:
		return "amd64_apple"
	case AMD64SysV:
		return "amd64_sysv"
	case ARM64:
		return "arm64"
	case ARM64Apple:
		return "arm64_apple"
	case RV64:
		return "rv64"
	default:
		panic(fmt.Sprintf("unexpected qbe.Target: %#v", target))
	}
}

// ToIL writes mod as a QBE IL file to w. Returns number of bytes written and if an error occured.
func (mod *Module) ToIL(w *bufio.Writer) (int, error) {
	written := 0
	for _, def := range mod.definitions {
		count, err := fmt.Fprint(w, def)
		written += count
		if err != nil {
			return written, err
		}
		if err = w.WriteByte('\n'); err != nil {
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

// ToAsm pipes mod IL to qbe and generates assembly for target named mod.name + ".s".
// Returns non nil on error.
func (mod *Module) ToAsm(target Target) error {
	cmd := exec.Command("qbe", "-", "-t", target.String(), "-o", mod.name+".s")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		_ = stdin.Close()
		return err
	}
	w := bufio.NewWriter(stdin)
	cleanup := func() {
		if stdin != nil {
			_ = stdin.Close()
		}
		_ = cmd.Process.Kill()
		_ = cmd.Wait()
	}
	if _, err = mod.ToIL(w); err != nil {
		cleanup()
		return err
	}
	if err = w.Flush(); err != nil {
		cleanup()
		return err
	}
	if err = stdin.Close(); err != nil {
		stdin = nil
		cleanup()
		return err
	}
	return cmd.Wait()
}

// ToObj pipes mod IL to qbe and cc and generates an object file for target named mod.name + ".o".
// Returns non nil on error.
func (mod *Module) ToObj(target Target, cflags []string) error {
	qbe := exec.Command("qbe", "-", "-t", target.String())
	cargs := []string{"-xassembler", "-", "-c", "-o", mod.name + ".o"}
	cargs = append(cargs, cflags...)
	cc := exec.Command("cc", cargs...)
	qbeStdin, err := qbe.StdinPipe()
	if err != nil {
		return err
	}
	qbeStdout, err := qbe.StdoutPipe()
	if err != nil {
		_ = qbeStdin.Close()
		return err
	}
	cc.Stdin = qbeStdout
	cleanup := func() {
		if qbeStdin != nil {
			_ = qbeStdin.Close()
		}
		_ = qbeStdout.Close()
		if qbe.Process != nil {
			_ = qbe.Process.Kill()
			_ = qbe.Wait()
		}
		if cc.Process != nil {
			_ = cc.Process.Kill()
			_ = cc.Wait()
		}
	}
	if err = qbe.Start(); err != nil {
		cleanup()
		return err
	}
	if err = cc.Start(); err != nil {
		cleanup()
		return err
	}
	w := bufio.NewWriter(qbeStdin)
	if _, err = mod.ToIL(w); err != nil {
		cleanup()
		return err
	}
	if err = w.Flush(); err != nil {
		cleanup()
		return err
	}
	if err = qbeStdin.Close(); err != nil {
		qbeStdin = nil
		cleanup()
		return err
	}
	qbeStdin = nil
	if err = qbe.Wait(); err != nil {
		qbe.Process = nil
		cleanup()
		return err
	}
	qbe.Process = nil
	return cc.Wait()
}
