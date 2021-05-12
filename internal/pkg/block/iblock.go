package block

type IBlock interface {
	HashBlock() (string, error)
}
