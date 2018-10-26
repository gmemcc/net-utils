package net_utils

import (
    "testing"
    "github.com/gmemcc/net-utils/ssh"
    "github.com/golang/glog"
)

func TestConnect(t *testing.T) {
    _, err := ssh.Connect("root", "lavender", "10.0.1.1", 22)
    if err != nil {
        glog.Fatal(err)
    }
}
func TestShell(t *testing.T) {
    sess, _ := ssh.Connect("root", "lavender", "10.0.1.1", 22)
    ssh.Shell(sess, func(composer ssh.CmdComposer) {
        composer.Write("echo Hello")
    })
}
