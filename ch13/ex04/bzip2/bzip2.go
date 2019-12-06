package bzip2

import (
	"io"
	"os/exec"
	"sync"

	"golang.org/x/xerrors"
)

type writer struct {
	mu  sync.Mutex
	cmd *exec.Cmd
	w   io.WriteCloser
	wg  sync.WaitGroup
}

const (
	BZIP = "/usr/bin/bzip2"
)

func NewWriter(out io.Writer) (io.WriteCloser, error) {
	var w writer
	w.cmd = exec.Command(BZIP)
	stdout, err := w.cmd.StdoutPipe()
	if err != nil {
		return nil, xerrors.Errorf("failed to create io.reader: %w", err)
	}
	stdin, err := w.cmd.StdinPipe()
	if err != nil {
		return nil, xerrors.Errorf("failed to create io.writer: %w", err)
	}
	w.w = stdin
	if err := w.cmd.Start(); err != nil {
		return nil, xerrors.Errorf("failed to exec %s: %w", BZIP, err)
	}
	w.wg.Add(1)
	go func() {
		defer w.wg.Done()
		io.Copy(out, stdout)
	}()
	return &w, nil
}

func (w *writer) Write(data []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	var total int
	for len(data) > 0 {
		n, err := w.w.Write(data)
		if err != nil {
			return total + n, err
		}
		total += n
		data = data[total:]

	}
	return total, nil
}

func (w *writer) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.w.Close()
	w.wg.Wait()
	if err := w.cmd.Wait(); err != nil {
		return err
	}
	return nil
}
