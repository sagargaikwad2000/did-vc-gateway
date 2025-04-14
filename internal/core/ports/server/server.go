package server

type IServer interface {
	MustStart()
	MustStop()
}
