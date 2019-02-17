// Copyright 2019 Andreas Grüninger.
// Use of this source code is governed by a zlib-style
// license that can be found in the LICENSE file.

package service

import (
	"errors"
	"os"
	"os/signal"
	"syscall"
)

var errNoUserServiceSMF = errors.New("User services are not supported on SMF.")
var errNotYetSupported = errors.New("Not yet supported on SMF.")

const version = "solaris-smf"

type solarisSystem struct{}

func (solarisSystem) String() string {
	return version
}
func (solarisSystem) Detect() bool {
	return true
}
func (solarisSystem) Interactive() bool {
	return interactive
}
func (solarisSystem) New(i Interface, c *Config) (Service, error) {
	s := &solarisSMFService{
		i:      i,
		Config: c,
	}
	return s, nil
}

func init() {
	ChooseSystem(solarisSystem{})
}

var interactive = false

func init() {
	var err error
	interactive, err = isInteractive()
	if err != nil {
		panic(err)
	}
}

func isInteractive() (bool, error) {
	// TODO: The PPID of Launchd is 1. The PPid of a service process should match launchd's PID.
	return os.Getppid() != 1, nil
}

type solarisSMFService struct {
	i Interface
	*Config
}

func (s *solarisSMFService) String() string {
	if len(s.DisplayName) > 0 {
		return s.DisplayName
	}
	return s.Name
}

func (s *solarisSMFService) Install() error {
	var err error
	err = errNotYetSupported
  return err
}

func (s *solarisSMFService) Uninstall() error {
	var err error
	err = errNotYetSupported
  return err
}

func (s *solarisSMFService) Start() error {
  return run("svcadm", "enable", s.Name)
}

func (s *solarisSMFService) Stop() error {
  return run("svcadm", "disable", s.Name)
}

func (s *solarisSMFService) Status() error {
  return checkStatus("svcs", []string{"-a", s.Name}, "online", "doesn't match any instances")
}

func (s *solarisSMFService) Restart() error {
  return run("svcadm", "restart", s.Name)
}

func (s *solarisSMFService) Run() error {
	var err error

	err = s.i.Start(s)
	if err != nil {
		return err
	}

	s.Option.funcSingle(optionRunWait, func() {
		var sigChan = make(chan os.Signal, 3)
		signal.Notify(sigChan, syscall.SIGTERM, os.Interrupt)
		<-sigChan
	})()

	return s.i.Stop(s)
}

func (s *solarisSMFService) Logger(errs chan<- error) (Logger, error) {
	if interactive {
		return ConsoleLogger, nil
	}
	return s.SystemLogger(errs)
}

func (s *solarisSMFService) SystemLogger(errs chan<- error) (Logger, error) {
	return newSysLogger(s.Name, errs)
}

