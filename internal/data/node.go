package data

type NodesQ interface {
	New() NodesQ
	Select() ([]string, error)
}
