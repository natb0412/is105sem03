package mycrypt

var ALF_SEM03 []rune = []rune("abcdefghijklmnopqrstuvwxyzæøå0123456789.,:; KSN")

func Krypter(melding []rune, ALF_SEM03 []rune, chiffer int) []rune {
	kryptertMelding := make([]rune, len(melding))
	for i := 0; i < len(melding); i++ {
		indeks := sokIAlfabetet(melding[i], ALF_SEM03)
		if indeks+chiffer >= len(ALF_SEM03) {
			kryptertMelding[i] = ALF_SEM03[indeks+chiffer-len(ALF_SEM03)]
		} else {
			kryptertMelding[i] = ALF_SEM03[indeks+chiffer]
		}
	}
	return kryptertMelding
}

func sokIAlfabetet(symbol rune, ALF_SEM03 []rune) int {
	for i := 0; i < len(ALF_SEM03); i++ {
		if symbol == ALF_SEM03[i] {
			return i

		}

	}
	return -1
}
