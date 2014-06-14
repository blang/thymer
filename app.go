package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	duration := flag.Duration("duration", 25*time.Minute, "Pomodoro duration")
	interval := flag.Duration("interval", 1*time.Second, "Update interval")
	barLen := flag.Int("bar", 20, "Length of progressbar")
	flag.Parse()

	thymer := NewThymer(*duration, *interval)
	notifyCh := make(chan ThymerNotification)
	go func() {
		for n := range notifyCh {
			fmt.Printf("%s %d:%02d\n", progressBar(100-int(n.PercLeft), *barLen), int(math.Floor(n.TimeLeft.Minutes())), int(math.Floor(n.TimeLeft.Seconds()))%60)
		}
	}()
	thymer.Start(notifyCh)

	closedCh := make(chan bool)
	go func() {
		thymer.Wait()
		close(closedCh)
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-closedCh:
		fmt.Println("Stopped")
	case <-c:
		thymer.Stop()
		<-closedCh
		fmt.Println("Interrupted")
	}
}

func progressBar(percent int, length int) string {
	stars := percent * length / 100
	return fmt.Sprintf("[%s%s]", strings.Repeat("#", stars), strings.Repeat(" ", length-stars))
}
