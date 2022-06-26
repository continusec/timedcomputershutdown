package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func generateAnnouncement(d time.Duration) string {
	ttl := int64(d.Round(time.Second).Seconds()) - 1
	switch {
	case ttl <= 10:
		return fmt.Sprintf("%d", ttl)
	case ttl <= 30 && ttl%5 == 0:
		return fmt.Sprintf("%d seconds", ttl)
	case ttl < 60 && ttl%10 == 0:
		return fmt.Sprintf("%d seconds", ttl)
	case ttl == 60:
		return fmt.Sprintf("%d minute warning", ttl/60)
	case ttl <= 5*60 && ttl%60 == 0:
		return fmt.Sprintf("%d minute warning", ttl/60)
	case ttl%(5*60) == 0:
		return fmt.Sprintf("%d minute warning", ttl/60)
	default:
		return ""
	}
}

func makeAnnouncement(s string) {
	if s == "" {
		return
	}
	log.Println(s)
	err := exec.Command("/usr/bin/say", "--", s).Run()
	if err != nil {
		log.Println(err)
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: timedcomputershutdown 20m (Golang duration string)")
	}
	duration, err := time.ParseDuration(os.Args[1])
	if err != nil {
		log.Fatal("Time must be a Golang duration string: ", err)
	}

	makeAnnouncement(fmt.Sprintf("Computer will turn off in %d minutes", int64(duration.Round(time.Minute).Minutes())))
	shutdownTime := time.Now().Add(duration)

	ticker := time.NewTicker(time.Second)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	for {
		select {
		case <-sigs:
			makeAnnouncement("Shutdown cancelled")
			os.Exit(1) // quit early
		case t := <-ticker.C:
			ttl := shutdownTime.Sub(t)
			if ttl <= 0 {
				makeAnnouncement("Shutting down")
				err := exec.Command("/sbin/shutdown", "-h", "now").Run()
				if err != nil {
					log.Println(err)
					makeAnnouncement("Error shutting down")
					os.Exit(1)
				}
				os.Exit(0)
			}
			go makeAnnouncement(generateAnnouncement(ttl))
		}
	}
}
