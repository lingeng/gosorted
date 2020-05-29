package main

import (
	"fmt"
	"math/rand"
	"time"

	sorted "github.com/lingeng/golang-sorted"
)

type transaction struct {
	Timestamp int64
}

func (tx transaction) Less(a sorted.EleInterface) bool {
	another, ok := a.(transaction)
	if !ok {
		panic("wrong type")
	}
	return tx.Timestamp < another.Timestamp
}

func main() {
	var txs sorted.List
	index := -1
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		tx := transaction{time.Now().UnixNano() - r.Int63()}
		index, txs = txs.Insert(tx)
		fmt.Printf("insert: %d, %d\n", tx.Timestamp, index)
	}

	for _, i := range txs {
		fmt.Println(i)
	}

	l := sorted.List{
		transaction{time.Now().UnixNano() - r.Int63()},
		transaction{time.Now().UnixNano() - r.Int63()},
		transaction{time.Now().UnixNano() - r.Int63()},
		transaction{time.Now().UnixNano() - r.Int63()},
		transaction{time.Now().UnixNano() - r.Int63()},
	}
	fmt.Println("init:", l)
	l.Sort()
	fmt.Println("sorted:", l)

	fmt.Println("The End")
}
