package pkg1

import (
	"time"

	"github.com/taiyoh/fopa/internal/pkg2"
	mypkg "github.com/taiyoh/fopa/internal/pkg3"
)

type tt int

//fopa -base=test1 -factory
type test1 struct {
	id        uint64
	aaa       string `fopa:"false"`
	bbb       tt     `fopa:"accept:int"`
	ccc       hoge   `fopa:"accept:string;expr:hoge{tt2({})}"`
	ddd       int
	eee       *int
	fff       mypkg.Fuga
	ggg       pkg2.Hoge `fopa:"accept:string"`
	createdAt time.Time
	deletedAt *time.Time
}
