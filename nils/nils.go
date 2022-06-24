package nils

func IsNil(t *any) bool {
	return t == nil
}

func IsNotNil(t *any) bool {
	return t != nil
}
