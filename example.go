package main

import (
	"fmt"
	"time"

	"github.com/goibibo/worktree"
)

func main() {
	fmt.Println("worktree usage Sync Run")
	wt := worktree.CommandTree{}
	n := 3
	t1 := time.Now()
	for i := n; i >= 1; i-- {
		wt.AddMapper(mapperSample, i)
	}
	wt.AddReducer(reducerSample)
	wt.Run(nil)
	fmt.Println("time taken", time.Now().Sub(t1))

	fmt.Println("\n\n\n worktree usage ASync Run")
	t1 = time.Now()
	wt1 := worktree.CommandTree{}
	for i := n; i >= 1; i-- {
		wt1.AddMapper(mapperSample, i)
	}
	wt1.AddReducer(AsyncReducerSample)
	wt1.RunMergeAsync(nil)
	fmt.Println("time taken", time.Now().Sub(t1))

}

func mapperSample(inp interface{}) interface{} {
	val := inp.(int)
	time.Sleep(time.Second * 2 * time.Duration(val))
	fmt.Println("added", val, time.Now())
	return "bla_bla" // its string, so reducers should typecaste with string
}

func reducerSample(inps []interface{}) interface{} {
	fmt.Printf("reducer input %+v\n", inps)
	//length of inps list will always be count of go-routine called.
	for _,inp := range inps {
		x := inp.(string)
		fmt.Println("reduc", x, time.Now())
	}

	return nil
}

func AsyncReducerSample(inps []interface{}) interface{} {
	fmt.Printf("Aysnc reducer input %+v\n", inps)
	//length of inps list will always be 2, where index 1 will have remaining count of goroutine yet to call reducer
	x := inps[0].(string)
	y := inps[1].(int)
	fmt.Println("reduc", x,y, time.Now())
	return nil
}
