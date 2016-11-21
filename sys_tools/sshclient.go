package main

/*
go get golang.org/x/crypto/ssh
*/

import (
	"bytes"
	"golang.org/x/crypto/ssh"
	"fmt"
	"os"
)

/*定义接收的参数:操作类型:exec_shell->linux连接IP,用户名，密码，命令。*/
type Parameters struct {
	conn_type, conn_ip, conn_user, conn_passwd, cmd_shell string
}

var (
	help = `
	Doc:
	conn_type:	Ex. exec_shell
	conn_ip: 	Ex. 1.1.1.1:22(www.scpman.com:22) 
	conn_user:  Ex. root
	conn_passwd:Ex.1234
	cmd_shell:  Ex. whoami
	Ex: ./cmd  "exec_shell" "www.scpman.com:22" "root" "123456" "ls /root/;whoami;"
	`
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

}

func main() {

	switch len(os.Args) {
	case 6: //exec_shell
		fmt.Println(Exec_shell(&Parameters{conn_type: os.Args[1], conn_ip: os.Args[2], conn_user: os.Args[3], conn_passwd: os.Args[4], cmd_shell: os.Args[5]}))
	default:
		fmt.Println(help)
	}

}

func Exec_shell(p *Parameters) string {
	fmt.Println(p.conn_type, p.conn_user, p.conn_passwd, p.conn_ip, p.cmd_shell)

	config := &ssh.ClientConfig{
		User: p.conn_user,
		Auth: []ssh.AuthMethod{
			ssh.Password(p.conn_passwd),
		},
	}
	client, err := ssh.Dial("tcp", p.conn_ip, config)
	checkError(err)

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	checkError(err)
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	err = session.Run(p.cmd_shell)
	checkError(err)
	return b.String()
}
