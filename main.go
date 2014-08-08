package main

import (
	"os"
	. "syscall"

	"github.com/mrunalp/usernsexec/forklib"
)

func main() {
	creds := &Credential{
		Uid: 0,
		Gid: 0,
	}

	uidMappings := []forklib.IdMap{
		forklib.IdMap{
			ContainerId: 0,
			HostId: 1013,
			Size: 1,
		},
	}

	gidMappings := []forklib.IdMap{
		forklib.IdMap{
			ContainerId: 0,
			HostId: 1013,
			Size: 1,
		},
	}

	pid, err := forklib.ForkExecNew("/bin/sh", []string{"sh"}, &ProcAttr{
		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
		Sys: &SysProcAttr{
			Cloneflags: CLONE_NEWNS | CLONE_NEWUSER | CLONE_NEWPID | CLONE_NEWNET | CLONE_NEWIPC | CLONE_NEWUTS,
			Credential: creds,
		},
	}, uidMappings, gidMappings)
	if err != nil {
		panic(err)
	}

	var wstatus WaitStatus
	_, err1 := Wait4(pid, &wstatus, 0, nil)
	if err != nil {
		panic(err1)
	}
}
