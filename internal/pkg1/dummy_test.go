package pkg1

import "time"

type tt int

//fopa -base=test1 -factory
type test1 struct {
	id        uint64
	aaa       string `fopa:"false"`
	bbb       tt     `fopa:"accept:int"`
	ccc       hoge   `fopa:"accept:string;expr:hoge{tt2({})}"`
	ddd       int
	createdAt time.Time
}
