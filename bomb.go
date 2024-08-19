/*
The tool is written and distributed by "esfelurm"
any copying without mentioning the source will be prosecuted (lol  @_@).
Telegram and Github: @esfelurm
*/
package main

import (
	"bufio"
	"fmt"
	"net/smtp"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

// Colors
// ----------------------------------
const (
	rd    = "\033[00;31m"
	gn    = "\033[00;32m"
	lgn   = "\033[01;32m"
	yw    = "\033[01;33m"
	lrd   = "\033[01;31m"
	be    = "\033[00;34m"
	pe    = "\033[01;35m"
	cn    = "\033[00;36m"
	k     = "\033[90m"
	g     = "\033[38;5;130m"
	reset = "\033[0m"
)

// ----------------------------------

func clearScreen() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Multi Account

func sendEmail1(auth smtp.Auth, from string, to string, subject string, body string, server string, wg *sync.WaitGroup) {
	defer wg.Done()
	message := []byte("Subject: " + subject + "\r\n\r\n" + body)
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, message)
	if err != nil {
		fmt.Printf("%s", rd)
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("%s", gn)
	fmt.Println("Email sent to:", to)
}

// ----------------------------------
// Send To Emails
func sendEmail2(auth smtp.Auth, subject string, body string, sender string, recipient string, server string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := []byte("Subject: " + subject + "\r\n" + body)

	err := smtp.SendMail("smtp.gmail.com:587", auth, sender, []string{recipient}, msg)
	if err != nil {
		fmt.Printf("%s", rd)
		fmt.Println("Failed to send email to:", recipient, "Error:", err)
		return
	}
	fmt.Printf("%s", gn)
	fmt.Println("Email sent to:", recipient)
}

// ----------------------------------

func Multi() {

	fmt.Printf("\n%s[%s+%s]%s Enter the filename (with email:password format): %s", rd, gn, rd, gn, cn)
	var filename string
	fmt.Scan(&filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var accounts []struct {
		email    string
		password string
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			fmt.Println("Invalid format. Expected format: email:password")
			continue
		}
		email := strings.TrimSpace(parts[0])
		password := strings.TrimSpace(parts[1])
		accounts = append(accounts, struct {
			email    string
			password string
		}{email, password})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var emailCount int
	fmt.Printf("%s[%s+%s]%s Enter the number of emails to send: %s", rd, gn, rd, gn, cn)
	fmt.Scan(&emailCount)

	var wg sync.WaitGroup
	var server, subject, body string
	fmt.Printf("%s[%s+%s]%s Enter Email Target: %s", rd, gn, rd, gn, cn)
	var to string
	fmt.Scan(&to)
	fmt.Printf("%s[%s+%s]%s Server and Port [ex: smtp.gmail.com:587]: %s", rd, gn, rd, gn, cn)
	fmt.Scan(&server)
	fmt.Printf("%s[%s+%s]%s Subject: %s", rd, gn, rd, gn, cn)
	fmt.Scan(&subject)
	fmt.Printf("%s[%s+%s]%s Text: %s", rd, gn, rd, gn, cn)
	fmt.Scan(&body)

	for _, account := range accounts {
		for i := 0; i < emailCount; i++ {
			wg.Add(1)
			auth := smtp.PlainAuth("", account.email, account.password, "smtp.gmail.com")
			go sendEmail1(auth, account.email, to, subject, body, server, &wg)
		}
	}
	wg.Wait()

}

func MultiAcc() {
	var server, listEmail, subject, body, sender, password string
	fmt.Printf("\n%s[%s+%s]%s Email: %s", rd, gn, rd, gn, cn)
	fmt.Scan(&sender)
	fmt.Printf("%s[%s+%s]%s Password: %s", rd, gn, rd, gn, cn)
	fmt.Scan(&password)
	fmt.Printf("%s[%s+%s]%s Server and Port [ex: smtp.gmail.com:587]: %s", rd, gn, rd, gn, cn)
	fmt.Scan(&server)
	fmt.Printf("%s[%s+%s]%s Address File Emails: %s", rd, gn, rd, gn, cn)
	fmt.Scan(&listEmail)
	fmt.Printf("%s[%s+%s]%s Subject: %s", rd, gn, rd, gn, cn)
	fmt.Scan(&subject)
	fmt.Printf("%s[%s+%s]%s Text: %s", rd, gn, rd, gn, cn)
	fmt.Scan(&body)

	var formatted strings.Builder
	for i, r := range password {
		if i > 0 && i%4 == 0 {
			formatted.WriteRune(' ')
		}
		formatted.WriteRune(r)
	}
	password = formatted.String()

	file, err := os.Open(listEmail)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	auth := smtp.PlainAuth("", sender, password, "smtp.gmail.com")
	var wg sync.WaitGroup
	for scanner.Scan() {
		recipient := strings.TrimSpace(scanner.Text())
		if recipient != "" {
			wg.Add(1)
			go sendEmail2(auth, subject, body, sender, recipient, server, &wg)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	wg.Wait()
}
func main() {
	clearScreen()
	fmt.Printf("%s", rd)
	fmt.Print(`
                                     +            
                                    -%.           
                                    *@-           
                             .*+-  .%@*   :.      
                               *@@##=#%*#@*       
                               .**=*.-+*@%:       
                         :=*#%@#+=:    =*#@%#*=.  
                             .:-*%-    =%%+-:.    
                                #@#%.+#+*@*.      
                               ==--==%# -*%%:     
                          ..:::-=#%.%@=    :=-    
                       :*%%%####*=. #@.           
                      :%#-          =*            
                      %@.           ::            
                     :@#  Email Bomber / Tg&Git: @esfelurm                        
                  :--=@+:--.                      
                +%@@#-+-*@@@%:                    
                %.:=+*****##%+                    
             .-.%   +##%%@@@@+::                  
          :+%@@-%   *@@@@@@@@++@%#=.              
        -#@@@@@-#=  *@@@@@@@%-*%#*%@*:            
      :#@@@@@@@%+++=+****++++*@@@%-=#@*.          
     +@@@@@@@@@@@@@%%#***#%@@@@@@@@+.=%%-         
   .#@@@@@@@@@@@#=:.::---:..-*%@@@@@* .#%=        
   #@@@@@@@@@@%- -#%%###@%%#= :%@@@@@#  #@=       
  =@@@@@@@@@@#..#@#: .:..  :@# .%@@@@@- .@%.      
  %@@@@@@@@@@- *@+  +@@@-  #@@: #@@@@@*  *@+      
 .@@@@@@@@@@@  %%. -@@@@- .@@%. %@@@@@%  =@#      
 .@@@@@@@@@@@. %%  -@@@*  +@#. #@@@@@@%  =@#      
  %@@@@@@@@@@= =@+  -=:   .. -%@@@@@@@*  *@*      
  *@@@@@@@@@@%- -#%*++#%*+*#%#--#@@@@@- -@@:      
  .%@@@@@@@@@@@*- .-=+***+=-. -#@@@@@@#@@@*       
   -@@@@@@@@@@@@@%#+=-----=+#%@@@@@@@@@@@%.       
    =%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@#.        
     :%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@*.         
       =%@@@@@@@@@@@@@@@@@@@@@@@@@@@@#:           
         =#@@@@@@@@@@@@@@@@@@@@@@@%*:             
           .=#%@@@@@@@@@@@@@@@@%*-.               
	`)
	var choice int
	fmt.Printf("\n%s[%s1%s]%s Multi Sender\n\n%s[%s2%s]%s Send To Emails\n\n\n    %s[%s+%s]%s Select:", rd, gn, rd, gn, rd, gn, rd, gn, rd, gn, rd, gn)
	fmt.Scan(&choice)
	switch choice {

	case 1:
		clearScreen()
		fmt.Printf("%s", cn)
		fmt.Print(`
 _______  _______  _______ _________ _       
(  ____ \(       )(  ___  )\__   __/( \      
| (    \/| () () || (   ) |   ) (   | (      
| (__    | || || || (___) |   | |   | |      
|  __)   | |(_)| ||  ___  |   | |   | |      
| (      | |   | || (   ) |   | |   | |      
| (____/\| )   ( || )   ( |___) (___| (____/\
(_______/|/     \||/     \|\_______/(_______/
		`)
		Multi()
	case 2:
		clearScreen()
		fmt.Printf("%s", cn)
		fmt.Print(`
                         .__ .__   
  ____    _____  _____   |__||  |  
_/ __ \  /     \ \__  \  |  ||  |  
\  ___/ |  Y Y  \ / __ \_|  ||  |__
 \___  >|__|_|  /(____  /|__||____/
     \/       \/      \/           
		
		`)
		MultiAcc()

	default:
		fmt.Printf("%sError", rd)
	}

}
