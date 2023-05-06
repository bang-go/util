package util

// VarAddr 获取变量地址
func VarAddr[T comparable](value T) *T {
	return &value
}

// VarVal VarAddress 获取变量值
func VarVal[T comparable](value *T) T {
	return *value
}
