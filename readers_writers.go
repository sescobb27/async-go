package main

import (
    "fmt"
    "sync"
    "time"
)

// type representing shared resource.  contents could be anything
type resource struct {
    string
}

// struct wraps pointer to resource with embedded RWMutex, thus
// acquiring RWMutex methods
type access struct {
    sync.RWMutex
    res *resource
}

// reader reads three times at about 1 second intervals, taking
// 2 second to perform the read
func reader(name string, acc *access, wg *sync.WaitGroup) {
    for i := 0; i < 3; i++ {
        time.Sleep(1 * time.Second)
        fmt.Println("reader", name, "ready to read")
        acc.RLock()
        fmt.Println("reader", name, "reading")
        time.Sleep(2 * time.Second)
        msg := acc.res.string
        acc.RUnlock()
        fmt.Println("reader", name, "read:", msg)
    }
    wg.Done()
}

// writer writes twice
func writer(name string, acc *access, wg *sync.WaitGroup) {
    claim := []string{"once", "again"}
    for i := 0; i < 2; i++ {
        time.Sleep(1 * time.Second)
        msg := name + " was here " + claim[i]
        fmt.Println("writer", name, "ready to write")
        acc.Lock()
        fmt.Println("writer", name, "writing:", msg)
        time.Sleep(2 * time.Second)
        acc.res.string = msg
        acc.Unlock()
        fmt.Println("writer", name, "done")
    }
    wg.Done()
}

func main() {
    acc := &access{res: &resource{"zero"}}
    fmt.Println("Initial value:", acc.res.string)

    wg := new(sync.WaitGroup)
    wg.Add(5) // three readers and two writers

    go reader("A", acc, wg)
    go reader("B", acc, wg)
    go reader("C", acc, wg)

    go writer("X", acc, wg)
    go writer("Y", acc, wg)

    wg.Wait()
    fmt.Println("Final value:", acc.res.string)
}
