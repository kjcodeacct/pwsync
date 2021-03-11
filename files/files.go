package files

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

const (
	CSVExtension = ".csv"
)

func ListenForType(dir string, fileExtension string) <-chan string {
	filenameCh := make(chan string)

	go func() {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer watcher.Close()

		done := make(chan bool)
		go func() {
			for {
				select {
				case event, ok := <-watcher.Events:
					if !ok {
						return
					}

					if event.Op == fsnotify.Create {
						newFileExtension := filepath.Ext(event.Name)
						if newFileExtension == CSVExtension {
							filenameCh <- event.Name
						}
					}

				case err, ok := <-watcher.Errors:
					if !ok {
						return
					}

					fmt.Println(err)
				}
			}
		}()

		err = watcher.Add(dir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		<-done
	}()

	return filenameCh
}

func Cleanup(filepath string, passes int) error {
	for i := 0; i < passes; i++ {
		f, err := os.OpenFile(filepath, os.O_WRONLY, 0)
		if err != nil {
			return err
		}

		defer f.Close()

		info, err := f.Stat()
		if err != nil {
			return err
		}

		buff := make([]byte, info.Size())

		_, err = f.WriteAt(buff, 0)
		if err != nil {
			return err
		}
	}

	err := os.Remove(filepath)
	if err != nil {
		return err
	}

	return nil
}
