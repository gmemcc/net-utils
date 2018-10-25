package ssh

import (
    "time"
    "fmt"
    "golang.org/x/crypto/ssh"
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
