package ssh

import (
    "time"
    "fmt"
    "golang.org/x/crypto/ssh"
    "io"
    "github.com/golang/glog"
)

func Connect(user, password, host string, port int) (*ssh.Session, error) {
    var (
        auth         []ssh.AuthMethod
        addr         string
        clientConfig *ssh.ClientConfig
        client       *ssh.Client
        err          error
    )
    auth = make([]ssh.AuthMethod, 0)
    auth = append(auth, ssh.Password(password))

    clientConfig = &ssh.ClientConfig{
        User:            user,
        Auth:            auth,
        Timeout:         30 * time.Second,
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    addr = fmt.Sprintf("%s:%d", host, port)
    if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
        return nil, err
    }
    return client.NewSession()
}

type CmdComposer interface {
    Write(cmd string)
}
type cmdComposer struct {
    stdinBuf io.WriteCloser
}

func (composer *cmdComposer) Write(cmd string) {
    composer.stdinBuf.Write([]byte(cmd))
    composer.stdinBuf.Write([]byte("\n"))
}

func newCmdComposer(sess *ssh.Session) CmdComposer {
    stdinBuf, err := sess.StdinPipe()
    if err != nil {
        glog.Errorf("Failed to create stdin pipe: %v", err)
    }
    return &cmdComposer{stdinBuf}
}

func Shell(sess *ssh.Session, cmdComposeCallback func(composer CmdComposer)) {
    composer := newCmdComposer(sess)
    sess.Shell()
    cmdComposeCallback(composer)
    composer.Write("exit")
    sess.Wait()
    sess.Close()
}
