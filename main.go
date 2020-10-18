package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

func main() {
	sshKeyPath := flag.String("ssh-key-path", "", "Key path for ssh authentication")
	url := flag.String("url", "", "Url to clone from")
	dest := flag.String("destination", "/tmp/repository", "folder to clone to")
	flag.Parse()
	fmt.Printf("Cloning from: %v\n", *url)

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	opts := &git.CloneOptions{
		URL:      *url,
		Progress: os.Stdout,
	}

	if len(*sshKeyPath) != 0 {
		publicKeys, err := ssh.NewPublicKeysFromFile("git", *sshKeyPath, "")
		if err != nil {
			fmt.Println(err)
			return
		}

		opts.Auth = publicKeys
	}

	_, err := git.PlainClone(*dest, false, opts)
	fmt.Println(err)

}
