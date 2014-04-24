package gopsutil

import (
	"encoding/json"
	"fmt"
	"runtime"
	"testing"
)

func Test_Pids(t *testing.T) {
	ret, err := Pids()
	if err != nil {
		t.Errorf("error %v", err)
	}
	if len(ret) == 0 {
		t.Errorf("could not get pids %v", ret)
	}
}

func Test_Pid_exists(t *testing.T) {
	check_pid := 1
	if runtime.GOOS == "windows" {
		check_pid = 0
	}

	ret, err := Pid_exists(int32(check_pid))
	if err != nil {
		t.Errorf("error %v", err)
	}

	if ret == false {
		t.Errorf("could not get init process %v", ret)
	}
}

func Test_NewProcess(t *testing.T) {
	check_pid := 1
	if runtime.GOOS == "windows" {
		check_pid = 0
	}

	ret, err := NewProcess(int32(check_pid))
	if err != nil {
		t.Errorf("error %v", err)
	}

	d, _ := json.Marshal(ret)
	fmt.Println(string(d))
}