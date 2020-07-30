/*
#######
##                __               __   __      __
##       ___ ____/ /__ __ _  __ __/ /  / /     / /__  ___ _
##      / -_) __/  '_//  ' \/ // / _ \/ / _   / / _ \/ _ `/
##      \__/\__/_/\_\/_/_/_/\_,_/_//_/_/ (_) /_/\___/\_, /
##                                                  /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package log

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"time"
)

type file struct {
	name    string
	file    *os.File
	reLevel *regexp.Regexp
	sigEnd  chan os.Signal
}

func (f *file) openFile() (bool, error) {
	fmt.Printf("--> %s\n", f.name)

	warn := true

	for {
		file, err := os.Open(f.name)
		if err == nil {
			fmt.Println("...")

			f.file = file

			return false, nil
		}

		if os.IsNotExist(err) {
			if warn {
				fmt.Println("--> this file doesn't exist...(wait or ^C ?)")

				warn = false
			}

			select {
			case <-time.After(100 * time.Millisecond):
			case <-f.sigEnd:
				fmt.Println("END")
				return true, nil
			}
		} else {
			return true, err
		}
	}
}

func (f *file) readFile() error {
	reader := bufio.NewReader(f.file)

	for {
		line, err := reader.ReadString('\n')
		switch err {
		case nil:
			printLine(f.reLevel.FindString(line), line)
		case io.EOF:
			select {
			case <-time.After(10 * time.Millisecond):
			case <-f.sigEnd:
				fmt.Println("END")
				return nil
			}
		default:
			return err
		}
	}
}

// TailFile AFAIRE.
func TailFile(fileName string) error {
	f := &file{
		name:    fileName,
		reLevel: regexp.MustCompile(`[{]\w{3}[}]`),
		sigEnd:  make(chan os.Signal, 1),
	}

	defer close(f.sigEnd)
	signal.Notify(f.sigEnd, syscall.SIGINT, syscall.SIGTERM)

	end, err := f.openFile()
	if err != nil {
		return err
	} else if end {
		return nil
	}

	defer f.file.Close()

	if err := f.readFile(); err != nil {
		return err
	}

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
