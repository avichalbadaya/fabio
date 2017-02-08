package logging

import (
	"testing"
	"os"
	"io/ioutil"
)

func TestAll(t *testing.T) {
	/*output := captureStdout(func() {Error("Test")})
	if !strings.Contains("Info",output){
		t.Fatalf("Not logging info")
	}*/
	Trace("Trace Logs")
	Debug("Debug logs")
	Info("Info logs")
	Warn("Warning logs")
	Error("Error logs")
}

func captureStdout(f func()) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
  f()
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
  return string(out)
}
