package arithmetic

type Adapter struct {
}

func (arith Adapter) Addition(a, b int32) (int32, error) {
	return a + b, nil
}
