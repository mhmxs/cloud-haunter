package types

const (
	Instances   = OpType("getInstances")
	CloudAccess = OpType("getAccess")
)

type OpType string

type Operation interface {
	Execute([]CloudType) []CloudItem
}

func (ot OpType) String() string {
	return string(ot)
}
