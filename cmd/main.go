package main

import (
	"fmt"
	"golog/elog"
	"golog/vars"
	"math/rand"
	"time"

	"github.com/vbauerster/mpb/v5"
	"github.com/vbauerster/mpb/v5/decor"
)

func main() {
	//p := mpb.New(mpb.PopCompletedMode())

	total, numBars := 100, 1
	for i := 0; i < numBars; i++ {
		name := fmt.Sprintf("Bar#%d:", i)
		bar := vars.ProcessBar.AddBar(int64(total),
			mpb.BarNoPop(),
			mpb.PrependDecorators(
				decor.Name(name),
				decor.Percentage(decor.WCSyncSpace),
			),
			mpb.AppendDecorators(
				decor.OnComplete(
					decor.EwmaETA(decor.ET_STYLE_GO, 60), "done!",
				),
			),
		)
		// simulating some work
		go func() {
			rng := rand.New(rand.NewSource(time.Now().UnixNano()))
			max := 100 * time.Millisecond
			for i := 0; i < total; i++ {
				// start variable is solely for EWMA calculation
				// EWMA's unit of measure is an iteration's duration
				start := time.Now()
				elog.Info(fmt.Sprintf("hey i start~ %d", i))
				time.Sleep(time.Duration(rng.Intn(10)+1) * max / 10)
				elog.Warn(fmt.Sprintf("hey i start~ %d", i))
				elog.Error(fmt.Sprintf("hey i start~ %d", i))
				bar.Increment()
				// we need to call DecoratorEwmaUpdate to fulfill ewma decorator's contract
				bar.DecoratorEwmaUpdate(time.Since(start))
			}
			//elog.Info("bbbb")
		}()
	}

	//var wg sync.WaitGroup
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	//rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	//	//max := 3000 * time.Millisecond
	//	for i := 0; i < 10; i++ {
	//		filler := makeLogBar(fmt.Sprintf("some log: %d", i))
	//		p.Add(0, filler).SetTotal(0, true)
	//		//time.Sleep(time.Duration(rng.Intn(10)+1) * max / 10)
	//	}
	//}()
	//
	//wg.Wait()
	vars.ProcessBar.Wait()
}
