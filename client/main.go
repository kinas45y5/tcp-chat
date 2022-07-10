package main

import (
  "net";
  "log";
  "os";
  "bufio";
  "strings";
  "fmt";
  "io"
)

// error function
func logFatal(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func main() {
  connection, err := net.Dial("tcp", "localhost:8080") // create connection to server   
  logFatal(err) // check for errors

  defer connection.Close() // close at the end

  fmt.Println("Enter your username")

  reader := bufio.NewReader(os.Stdin)      // reading the username
  username, err := reader.ReadString('\n')

  logFatal(err)

  username = strings.Trim(username, " \r\n") // store username
 
  welcomeMsg := fmt.Sprintf("Welcome %s to the chat.\n", username)

  fmt.Println(welcomeMsg)  

  go read(connection)
  write(connection, username)
  
}

func read(connection net.Conn) {
  for {
    reader := bufio.NewReader(connection)
    message, err := reader.ReadString('\n')

    if err == io.EOF {
      connection.Close()
      fmt.Println("Connection closed.")
      os.Exit(0)
    }

    fmt.Println("----------------------------------------------------")
    fmt.Println(message)
    fmt.Println("----------------------------------------------------")
  }
}

func write(connection net.Conn, username string) {
  for {
    reader := bufio.NewReader(os.Stdin)
    message, err := reader.ReadString('\n')
  
    if err != nil {
      break
    }

    // username:- message

    message = fmt.Sprintf("%s:- %s\n", username, strings.Trim(message, " \r\n"))

    connection.Write([]byte(message))
  }
}
