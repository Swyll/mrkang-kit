package plugin

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/Swyll/mrkang-kit/plugin/command"
)

func TestSched(t *testing.T) {
	cmd1 := command.NewCommandPlu("ls")
	cmd2 := command.NewCommandPlu("echo ddddd")
	cmd3 := command.NewCommandPlu("./example")
	cmd4 := command.NewCommandPlu("./example")
	cmd5 := command.NewCommandPlu("./example")

	ps := []Plugin{cmd1, cmd2, cmd3, cmd4, cmd5}

	schd, msgch := NewSche(ps...)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		schd.Run()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for msg := range msgch {
			fmt.Printf("%+v\n", msg)
		}

		fmt.Println("dddddd")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		//time.Sleep(time.Second * 1)
		//schd.Stop()

		time.Sleep(time.Second * 2)
		schd.ReStart()

		time.Sleep(time.Second * 2)
		schd.Finalized()
	}()

	wg.Wait()
}
