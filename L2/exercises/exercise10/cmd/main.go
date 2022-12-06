package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"pkg/arguments"
	t "pkg/tcpConnector"
)

func main() {
	arg := *arguments.MakeArguments()
	arg.FillArguments()
	connector := *t.MakeTcpConnector(arg)
	ctx := context.Background()
	var wg sync.WaitGroup
	wg.Add(1)
	connector.RunTelnet(ctx, &wg)
	exit := make(chan os.Signal)
	signal.Notify(exit,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGINT,
		syscall.SIGHUP,
	)

	go func() {
		<-exit
		fmt.Println("Программа завершена вручную")
		ctx.Done()
	}()
	wg.Wait()
}
