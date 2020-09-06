package main

/*
#include <stdio.h>
#include <unistd.h>
#include <termios.h>
char getch(){
    char ch = 0;
    struct termios old = {0};
    fflush(stdout);
    if( tcgetattr(0, &old) < 0 ) perror("tcsetattr()");
    old.c_lflag &= ~ICANON;
    old.c_lflag &= ~ECHO;
    old.c_cc[VMIN] = 1;
    old.c_cc[VTIME] = 0;
    if( tcsetattr(0, TCSANOW, &old) < 0 ) perror("tcsetattr ICANON");
    if( read(0, &ch,1) < 0 ) perror("read()");
    old.c_lflag |= ICANON;
    old.c_lflag |= ECHO;
    if(tcsetattr(0, TCSADRAIN, &old) < 0) perror("tcsetattr ~ICANON");
    return ch;
}
*/
import "C"

import (
	rand "crypto/rand"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/big"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	d bool
	f string
)

func init() {
	flag.BoolVar(&d, "d", false, "Enable debug log")
	flag.StringVar(&f, "f", "", "Output to file")

	// set log to json format
	log.SetFormatter(&log.JSONFormatter{})

	flag.Parse()
	if d {
		// set log to debug level
		log.SetLevel(log.DebugLevel)
	} else {
		// set log to warn level
		log.SetLevel(log.WarnLevel)
	}
	if f != "" {
		file, err := os.OpenFile(f, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err == nil {
			// set log output to file
			log.SetOutput(file)
		} else {
			// set log output to stdout
			log.SetOutput(os.Stdout)
			log.Warn("Failed to log to file, using default stderr")
		}
	} else {
		// set log output to stdout
		log.SetOutput(os.Stdout)
	}

}

func main() {
	// create a chan
	c := make(chan os.Signal)
	// listen ctrl+c/ctrl+z/kill
	signal.Notify(c, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM:
				log.WithFields(log.Fields{
					"signal": s,
				}).Warn("Received signal!")
				exitFunc()
			default:
				log.WithFields(log.Fields{
					"signal": s,
				}).Warn("Received other signal!")
			}
		}
	}()

	log.Info("The process is starting...")
	var char int
	for {
		fmt.Println(printNums())
		fmt.Println("Press Q key to exit or press other key to continue")
		// getch() from c code
		char = int(C.getch())
		// Q=81 or q=113
		if char == 113 || char == 81 {
			log.Info("Q key input received!")
			exitFunc()
		}
	}
}

func exitFunc() {
	log.Info("The process is exiting...")
	log.Info("Cleaning up...")
	termEcho()
	os.Exit(0)
}

// turns terminal echo on during sigterm.
func termEcho() {
	// Common settings and variables for stty calls.
	attrs := syscall.ProcAttr{
		Dir:   "",
		Env:   []string{},
		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
		Sys:   nil}
	var ws syscall.WaitStatus

	cmd := "echo"
	// Enable echoing.
	pid, err := syscall.ForkExec(
		"/bin/stty",
		[]string{"stty", cmd},
		&attrs)
	if err != nil {
		panic(err)
	}
	// Wait for the stty process to complete.
	_, err = syscall.Wait4(pid, &ws, 0, nil)
	if err != nil {
		panic(err)
	}
}

func genRandIndex(min, max int64) int64 {
	if min == max {
		return 0 + min
	}
	if min > max {
		return min
	}

	log.WithFields(log.Fields{
		"min": min,
		"max": max,
	}).Debug("genRandIndex debug")
	// get value between 0 and max - min - 1
	n, err := rand.Int(rand.Reader, big.NewInt(max-min))
	if err != nil {
		panic(err)
	}

	// add n to min to support the passed in range
	return n.Int64() + min
}

func findStr(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func printNums() string {

	var index int
	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	output := []string{}

	for len(numbers) > 0 {
		index = int(genRandIndex(0, int64(len(numbers))))
		log.WithFields(log.Fields{
			"index":         index,
			"number":        numbers[index],
			"output_length": len(output),
		}).Debug("Random index generated")

		if findStr(output, numbers[index]) == false {
			output = append(output, numbers[index])
			copy(numbers[index:], numbers[index+1:]) // Shift numbers[index+1:] left one index.
			numbers[len(numbers)-1] = ""             // Erase last element (write zero value).
			numbers = numbers[:len(numbers)-1]       // Truncate slice.
		} else {
			continue
		}
	}

	return strings.Join(output, ",")
}
