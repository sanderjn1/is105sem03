package mycrypt

var ALF_SEM03 []rune = []rune("abcdefghijklmnopqrstuvwxyzæøå0123456789ABCDEFGHIJKLMONPQRSTUWXYZÆØÅ.,:; ")

func Krypter(melding []rune, alphabet []rune, chiffer int) []rune {

	modCipher := chiffer % len(ALF_SEM03)

	shiftedAlphabet := shiftAlphabet(ALF_SEM03, modCipher)

	lookup := make(map[rune]rune)

	for i := 0; i < len(ALF_SEM03); i++ {
		lookup[ALF_SEM03[i]] = shiftedAlphabet[i]
	}

	kryptertMelding := make([]rune, len(melding))
	for i := 0; i < len(melding); i++ {
		kryptertMelding[i] = lookup[melding[i]]
	}
	return kryptertMelding
}

func shiftAlphabet(alphabet []rune, chiffer int) []rune {
	return append(alphabet[chiffer:], alphabet[:chiffer]...)
}

func sokIAlfabetet(symbol rune, alfabet []rune) int {
	for i := 0; i < len(alfabet); i++ {
		if symbol == alfabet[i] {
			return i
			break
		}
	}
	return -1
}
