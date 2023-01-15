package main

import (
	"reflect"
	"runtime"

	"github.com/robfig/cron/v3"
)

type spec struct {
	schedule string
	cmd      string
	id       cron.EntryID
}

type myCron struct {
	cron      *cron.Cron
	cronSpecs []spec
}

func (m *myCron) addFunc(sched string, cmd func()) error {
	id, err := m.cron.AddFunc(sched, cmd)
	if err != nil {
		return err
	}
	cmdName := runtime.FuncForPC(reflect.ValueOf(cmd).Pointer()).Name()
	m.cronSpecs = append(m.cronSpecs, spec{schedule: sched, cmd: cmdName, id: id})
	return nil
}

func (m *myCron) addJob(sched string, j job) {
	m.cron.AddJob(sched, j)
}

func (m *myCron) start() {
	m.cron.Start()
}
