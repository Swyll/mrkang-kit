package command

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPlugin(t *testing.T) {
	cmd := NewCommandPlu("./swy")

	b := make([]byte, 10, 10)
	bf := bytes.NewBuffer(b)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := cmd.Run(bf)
		if err != nil {
			t.Errorf("%+v", err)
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Println("11111")
		time.Sleep(time.Second * 2)
		cmd.Cannel()
	}()

	wg.Wait()
	fmt.Println(bf.String())
}

func TestKill(t *testing.T) {
	err := kill(17598)
	if err != nil {
		t.Error(err)
		return
	}
}
