package xssh

//
//import (
//	"fmt"
//	"testing"
//)
//
//func TestNewClient(t *testing.T) {
//
//	c, err := NewClient("127.0.0.1", "22", "root", "password")
//	if err != nil {
//		panic(err)
//	}
//	defer c.Close()
//	//cmd := `source installZtest`
//	cmd := `ztest version`
//	//cmd := `lscpu;fdisk -l`
//	output, err := c.Output(cmd)
//	fmt.Printf("%s", output)
//
//	cmd = `ztest -h`
//	output, err = c.Output(cmd)
//	fmt.Printf("%s", output)
//}
//
////func main() {
////
////	client, err := remotessh.NewClient( "172.20.1.34", "22", "root", "password")
////	if err != nil {
////		panic(err)
////	}
////	defer client.Close()
////	var remotedir = "/root/test/"
////	// upload dir
////	var local = "./test/"
////	client.Upload(local, remotedir)
////
////	// upload file
////	local = "/root/test"
////	client.Upload(local, remotedir)
////}
//
////func main() {
////
////	client, err := ssh.NewClient( "localhost", "22", "root", "ubuntu")
////	if err != nil {
////		panic(err)
////	}
////	defer client.Close()
////	var remotedir = "/root/test/"
////	// download dir
////	var local = "/home/ubuntu/go/src/github.com/pytool/ssh/test/download/"
////	client.Download(remotedir, local)
////
////	// upload file
////	var remotefile = "/root/test/file"
////
////	client.Download(remotefile, local)
////}
