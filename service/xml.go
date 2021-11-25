package service

import (
	"data-pipeline/config"
	"data-pipeline/util"
	"fmt"
	"io/ioutil"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func XMLRetriever(date string) ([]byte, error) {

	var auths []ssh.AuthMethod
	auths = append(auths, ssh.Password(config.SFTPPass))

	// Initialize client configuration
	conf := ssh.ClientConfig{
		User:            config.SFTPUser,
		Auth:            auths,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	addr := fmt.Sprintf("%s:%d", config.SFTPHost, 22)

	// Connect to server
	conn, err := ssh.Dial("tcp", addr, &conf)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	// Create new SFTP client
	sc, err := sftp.NewClient(conn)
	if err != nil {
		return nil, err
	}
	defer sc.Close()

	// date := "2020-11-24"

	file, err := sc.Open(fmt.Sprintf("/Users/RM8005-jeffrey/Project/air-asia/GL_AK_%s.xml", date))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	xmlByte, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	//archive the file if the file exists
	if err := util.Archive(xmlByte, fmt.Sprintf("GL_AK_%s.xml", date), fmt.Sprintf("/Users/RM8005-jeffrey/Project/air-asia/archive_%s.zip", date)); err != nil {
		return nil, err
	}

	return xmlByte, nil
}
