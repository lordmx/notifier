package main

import(
	"log"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

type Notifier struct {
	dirs []string
	mask string	
	watcher *fsnotify.Watcher
}

func NewNotifier(config *Config, callback func(config *Config) error) *Notifier, err {
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		return nil, err
	}

	defer watcher.Close()

	notifier := &Notifier{
		dirs: config.Paths,
		mask: config.Mask,
		watcher: watcher,
	}

	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					matched, err := filepath.Match(strings.ToLower(config.Mask), strings.ToLower(event.Name))

					if err != nil {
						log.Println("error:", err)
					} else if (matched) {
						log.Println("modified file:", event.Name)
						err := callback(config)

						if err != nil {
							log.Println("error:", err)
						}
					}
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	for _, dir := range config.Paths {
		err = watcher.Add(dir)

		if err != nil {
			return nil, err
		}

	}

	<-done

	return notifier, nil
}