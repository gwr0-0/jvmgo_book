package rtda

type LocalVars []slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]slot, maxLocals)
	}
	return nil
}
