package newbee_traps

func NilInterface() (*byte, interface{}, interface{}) {
	var data *byte
	var in interface{}
	in2 := in
	in2 = data

	return data, in, in2
}
