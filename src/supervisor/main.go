package main

import (
  "fmt"
  "os"
  "os/signal"
  "time"

  "cirello.io/supervisor"
  "context"
)

type Simpleservice int

func (s *Simpleservice) String() string {
  return fmt.Sprintf("Simpleservice  %d", int(*s))
}

func (s *Simpleservice) Serve(ctx context.Context) {
  for {
    select {
    case <-ctx.Done():
      return
    default:
      fmt.Println("Doing something...")
      time.Sleep(500 * time.Millisecond)
    }
  }
}

func main(){
  svc := Simpleservice(1)
  supervisor.Add(&svc)

  // Simply, if not special context is needed:
  // supervisor.Serve()
  // Or, using context.Context to propagate behavior:
  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt)
  ctx, cancel := context.WithCancel(context.Background())
  go func(){
    <-c
    fmt.Println("halting supervisor...")
    cancel()
  }()
  supervisor.ServeContext(ctx)
}
