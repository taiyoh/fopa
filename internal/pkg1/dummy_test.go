package pkg1

type tt int

//fopa -base=test1 -factory
type test1 struct {
	aaa string `fopa:"11;22:33"`
	bbb tt     `fopa:"44:55;66"`
	ccc hoge   `fopa:"77;88:99"`
}
