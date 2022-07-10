package main

import (
  "log";
  "net";
  "bufio"
)

func logFatal(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

var (
  openConnections = make(map[net.Conn]bool)
  newConnection   = make(chan net.Conn)
  deadConnection  = make(chan net.Conn)
)

func broadcastMessage(conn net.Conn) {
    for {
      reader := bufio.NewReader(conn)
      message, err := reader.ReadString('\n')
    

      if err != nil {
        break
      }

      // loop thru all open openConnections
      // & send msgs to connections except the one that sent the msg

      for item := range openConnections {
        if item != conn {
          item.Write([]byte(message))
        }   
      }
    }

    deadConnection <- conn
}

func main() {
  ln, err := net.Listen("tcp", ":8080")
  logFatal(err)

  defer ln.Close()

  go func() {
    for {
      conn, err := ln.Accept()
      logFatal(err)

      openConnections[conn] = true
      newConnection <- conn
    }
  }()

  for {
    select {
      case conn := <-newConnection:
      // invoke broadcasted message (broadcasts to other connections)
        go broadcastMessage(conn)

      case conn := <-deadConnection:
        // delete connection
        for item := range openConnections {
          if item == conn {
            break
          }
      }
      delete(openConnections, conn)
    }
  }
}

