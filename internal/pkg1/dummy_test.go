package pkg1

type tt int

//fopa -base=test1 -factory
type test1 struct {
	aaa string `fopa:"false"`
	bbb tt     `fopa:"accept:int"`
	ccc hoge   `fopa:"accept:string;expr:hoge{tt2({})}"`
	ddd int
}
