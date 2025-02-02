package figure

import (
	"fmt"
	"os/exec"
	"runtime"
)

func openBrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "darwin":
		err = exec.Command("open", url).Start()
	case "windows":
		err = exec.Command("cmd", "/c", "start", url).Start()
	default: // "linux", "freebsd", etc.
		err = exec.Command("xdg-open", url).Start()
	}

	if err != nil {
		return fmt.Errorf("error opening browser: %v", err)
	}

	return nil
}
