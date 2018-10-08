package preprocess

func PreProcessLog() {
	err := ProcessCQ9Log()
	if err != nil {
		return
	}
}
