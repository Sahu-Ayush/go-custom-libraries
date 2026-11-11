package systeminfo

import (
	"runtime"
	"testing"
)

func TestGetOS(t *testing.T) {
	got := GetOS()
	want := runtime.GOOS

	if got != want {
		t.Errorf("GetOS() = %v, want %v", got, want)
	}
}

func TestGetArchitecture(t *testing.T) {
	got := GetArchitecture()
	want := runtime.GOARCH

	if got != want {
		t.Errorf("GetArchitecture() = %v, want %v", got, want)
	}
}

func TestGetHostname(t *testing.T) {
	got, err := GetHostname()

	if err != nil {
		t.Fatalf("GetHostname() returned error: %v", err)
	}

	if got == "" {
		t.Error("GetHostname() returned an empty hostname")
	}
}

func TestGetCPUCount(t *testing.T) {
	got := GetCPUCount()

	if got <= 0 {
		t.Errorf("GetCPUCount() = %v, want a positive number", got)
	}
}

func TestGetGoVersion(t *testing.T) {
	got := GetGoVersion()
	want := runtime.Version()

	if got != want {
		t.Errorf("GetGoVersion() = %v, want %v", got, want)
	}
}
