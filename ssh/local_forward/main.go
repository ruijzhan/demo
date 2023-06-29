package main

import (
	"io"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

type SSH struct {
	sshConfig *ssh.ClientConfig
}

func NewSSH(user, password string) *SSH {

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return &SSH{
		sshConfig: config,
	}
}

func (s *SSH) NewReverseTunnel(sshHost, sshPort, fwdAddr, fwdPort, listenIP, listenPort string) (*ReverseTunnel, error) {
	client, err := ssh.Dial("tcp", sshHost+":"+sshPort, s.sshConfig)
	if err != nil {
		return nil, err
	}

	rt := &ReverseTunnel{
		sshClient:  client,
		fwdAddr:    fwdAddr,
		fwdPort:    fwdPort,
		listenIP:   listenIP,
		listenPort: listenPort,
	}

	return rt, nil
}

type ReverseTunnel struct {
	sshClient *ssh.Client

	// 想要代理的远程主机地址
	fwdAddr string
	fwdPort string

	// 想要在访问远程地址的本机 IP 和端口
	listenIP   string
	listenPort string
}

func (r *ReverseTunnel) Run() error {

	listener, err := r.sshClient.Listen("tcp", r.listenIP+":"+r.listenPort)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go func() {
			defer conn.Close()

			fwdConn, err := net.Dial("tcp", r.fwdAddr+":"+r.fwdPort)
			if err != nil {
				return
			}
			defer fwdConn.Close()

			chDone := make(chan struct{})

			go func() {
				io.Copy(fwdConn, conn)
				chDone <- struct{}{}
			}()

			go func() {
				io.Copy(conn, fwdConn)
				chDone <- struct{}{}
			}()

			<-chDone
		}()
	}

}

func (r *ReverseTunnel) Close() {

}

func createSSHClientConfig(user, password string) (*ssh.ClientConfig, error) {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return config, nil
}

func createReverseTunnel(sshAddr, remoteHost, remotePort, localHost, localPort string, config *ssh.ClientConfig) error {
	client, err := ssh.Dial("tcp", sshAddr, config)
	if err != nil {
		return err
	}
	defer client.Close()

	listener, err := client.Listen("tcp", remoteHost+":"+remotePort)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go func() {
			defer conn.Close()
			localConn, err := net.Dial("tcp", localHost+":"+localPort)
			if err != nil {
				return
			}

			chDone := make(chan struct{})

			go func() {
				io.Copy(localConn, conn)
				chDone <- struct{}{}
			}()

			go func() {
				io.Copy(conn, localConn)
				chDone <- struct{}{}
			}()

			<-chDone
		}()
	}

}

func main() {
	cliConf, err := createSSHClientConfig("ruijzhan", "aca04rz.")
	if err != nil {
		log.Fatalln(err)
	}
	err = createReverseTunnel("192.168.9.1:55222", "0.0.0.0", "55222", "news.163.com", "80", cliConf)
	if err != nil {
		log.Fatalln(err)
	}
}
