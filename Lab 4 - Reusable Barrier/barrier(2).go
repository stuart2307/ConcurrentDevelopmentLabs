//Barrier.go Template Code
//Copyright (C) 2024 Dr. Joseph Kehoe

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Stuart Rossiter (C00284845@setu.ie)
// Issues:
//--------------------------------------------

package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/semaphore"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, wg *sync.WaitGroup, count *atomic.Int32, barrierOne chan struct{}, barrierTwo chan struct{}, total int) bool {
	time.Sleep(time.Second)
	for range 5 {
		fmt.Println("Part A", goNum)
		if count.Add(1) == int32(total) {
			for i := 1; i < total; i++ {
				barrierOne <- struct{}{}
			}
		} else {
			<-barrierOne
		}
		//we wait here until everyone has completed part A
		fmt.Println("PartB", goNum)
		if count.Add(-1) == 0 {
			for i := 1; i < total; i++ {
				barrierTwo <- struct{}{}
			}
		} else {
			<-barrierTwo
		}
	}
	wg.Done()
	return true
}

func main() {
	totalRoutines := 10
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	//we will need some of these
	ctx := context.TODO()
	var theLock sync.Mutex
	sem := semaphore.NewWeighted(int64(totalRoutines))
	theLock.Lock()
	var count atomic.Int32
	barrierOne := make(chan struct{})
	barrierTwo := make(chan struct{})
	sem.Acquire(ctx, 1)
	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &wg, &count, barrierOne, barrierTwo, totalRoutines)
	}
	sem.Release(1)
	theLock.Unlock()

	wg.Wait() //wait for everyone to finish before exiting
}
