package servers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/darxkies/k8s-tew/config"
	"github.com/darxkies/k8s-tew/utils"

	log "github.com/sirupsen/logrus"
)

type ServerWrapper struct {
	stop          bool
	name          string
	baseDirectory string
	command       []string
	logger        config.LoggerConfig
}

func NewServerWrapper(_config config.InternalConfig, name string, serverConfig config.ServerConfig) (Server, error) {
	var error error

	serverConfig.Command, error = _config.ApplyTemplate("command", serverConfig.Command)

	if error != nil {
		return nil, error
	}

	server := &ServerWrapper{name: name, baseDirectory: _config.BaseDirectory, command: []string{serverConfig.Command}, logger: serverConfig.Logger}

	server.logger.Filename, error = _config.ApplyTemplate("LoggingDirectory", server.logger.Filename)
	if error != nil {
		return nil, error
	}

	for key, value := range serverConfig.Arguments {
		if len(value) == 0 {
			server.command = append(server.command, fmt.Sprintf("--%s", key))

		} else {
			newValue, error := _config.ApplyTemplate(fmt.Sprintf("%s.%s", server.Name(), key), value)
			if error != nil {
				return nil, error
			}

			server.command = append(server.command, fmt.Sprintf("--%s=%s", key, newValue))
		}
	}

	return server, nil
}

func (server *ServerWrapper) Start() error {
	server.stop = false

	if server.logger.Enabled {
		logsDirectory := filepath.Dir(server.logger.Filename)

		if error := utils.CreateDirectoryIfMissing(logsDirectory); error != nil {
			return error
		}
	}

	go func() {
		for !server.stop {
			log.WithFields(log.Fields{"name": server.Name(), "command": strings.Join(server.command, " ")}).Info("starting server")

			command := exec.Command(server.command[0], server.command[1:]...)

			var logFile *os.File
			var error error

			if server.logger.Enabled {
				logFile, error = os.OpenFile(server.logger.Filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

				if error != nil {
					log.WithFields(log.Fields{"filename": logFile, "error": error}).Error("could no open file")

					continue
				}

				command.Stdout = logFile
				command.Stderr = logFile
			}

			defer func() {
				if logFile != nil {
					logFile.Close()
				}
			}()

			command.Run()

			time.Sleep(time.Second)

			if !server.stop {
				log.WithFields(log.Fields{"name": server.name}).Error("server terminated")
			}
		}
	}()

	return nil
}

func (server *ServerWrapper) Stop() {
	server.stop = true
}

func (server *ServerWrapper) Name() string {
	return server.name
}
