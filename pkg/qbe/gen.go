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
