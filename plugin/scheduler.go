package plugin

import (
	"bytes"
	"sync"
	"sync/atomic"
	"time"
)

type PluginScheduler struct {
	run    func(Plugin)
	funcMu sync.Mutex

	ps    []Plugin
	index int64
	msg   chan *message
}

type message struct {
	Stdout string
	Err    error
}

func NewSche(ps ...Plugin) (*PluginScheduler, <-chan *message) {
	schd := &PluginScheduler{
		ps:  ps,
		msg: make(chan *message, len(ps)),
	}
	schd.run = schd.play

	return schd, schd.msg
}

func (schd *PluginScheduler) Run() {
	defer close(schd.msg)

	for int(schd.index) < len(schd.ps) {
		index := schd.index
		if index == -1 {
			atomic.StoreInt64(&schd.index, 0)
			index = 0
		}

		schd.funcMu.Lock()
		run := schd.run
		schd.funcMu.Unlock()

		run(schd.ps[index])
	}
}

func (schd *PluginScheduler) Start() {
	schd.funcMu.Lock()
	defer schd.funcMu.Unlock()

	schd.run = schd.play
}

func (schd *PluginScheduler) Stop() {
	schd.funcMu.Lock()
	defer schd.funcMu.Unlock()

	schd.run = func(Plugin) {
		time.Sleep(time.Second * 1)
	}
}

func (schd *PluginScheduler) Finalized() {
	schd.Stop()
	index := atomic.LoadInt64(&schd.index)

	if index == -1 {
		index = 0
	}
	schd.ps[index].Cannel()

	atomic.StoreInt64(&schd.index, int64(len(schd.ps)))
}

func (schd *PluginScheduler) ReStart() {
	schd.Stop()

	schd.funcMu.Lock()
	schd.run = schd.play
	schd.funcMu.Unlock()

	schd.ps[schd.index].Cannel()

	schd.index = -1
}

func (schd *PluginScheduler) play(p Plugin) {
	var b bytes.Buffer

	err := p.Run(&b)
	schd.msg <- &message{
		Stdout: b.String(),
		Err:    err,
	}

	atomic.AddInt64(&schd.index, 1)
}
