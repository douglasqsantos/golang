package errors

func HandError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
