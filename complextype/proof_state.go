package complextype

type ProofStateType string

const (
	CLEAN ProofStateType = "clean"
	DIRTY ProofStateType = "dirty"
)

type ProofState struct {
	Spelling ProofStateType
	Grammar  ProofStateType
}
