package utils 

func CheckError(description string, err error) {
	if err != nil {
		panic(err)
	}
}
