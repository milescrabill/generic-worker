package main

func helloGoodbye() []string {
	return []string{
		"echo hello world!",
		"echo goodbye world!",
	}
}

func checkSHASums() []string {
	return []string{
		"powershell -file preloaded\\check-shasums.ps1",
	}
}
