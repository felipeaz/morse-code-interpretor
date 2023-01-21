package main

import (
	"fmt"
	"sort"
)

const (
	maxSignalsLength = 3
	unknownSignal    = "?"
	morseCodeA       = ".-"
	morseCodeB       = "-..."
	morseCodeC       = "-.-."
	morseCodeD       = "-.."
	morseCodeE       = "."
	morseCodeF       = "..-."
	morseCodeG       = "--."
	morseCodeH       = "...."
	morseCodeI       = ".."
	morseCodeJ       = ".---"
	morseCodeK       = "-.-"
	morseCodeL       = ".-.."
	morseCodeM       = "--"
	morseCodeN       = "-."
	morseCodeO       = "---"
	morseCodeP       = ".--."
	morseCodeQ       = "--.-"
	morseCodeR       = ".-."
	morseCodeS       = "..."
	morseCodeT       = "-"
	morseCodeU       = "..-"
	morseCodeV       = "...-"
	morseCodeW       = ".--"
	morseCodeX       = "-..-"
	morseCodeY       = "-.--"
	morseCodeZ       = "--.."
)

type morseCode struct {
	letter string
	weigth int
}

var morseDichotomicCodeMap = map[string]morseCode{
	morseCodeE: {letter: "E", weigth: 100}, morseCodeT: {letter: "T", weigth: 99}, morseCodeI: {letter: "I", weigth: 98},
	morseCodeA: {letter: "A", weigth: 97}, morseCodeN: {letter: "N", weigth: 96}, morseCodeM: {letter: "M", weigth: 95},
	morseCodeS: {letter: "S", weigth: 94}, morseCodeU: {letter: "U", weigth: 93}, morseCodeR: {letter: "R", weigth: 92},
	morseCodeW: {letter: "W", weigth: 91}, morseCodeD: {letter: "D", weigth: 90}, morseCodeK: {letter: "K", weigth: 89},
	morseCodeG: {letter: "G", weigth: 88}, morseCodeO: {letter: "O", weigth: 87}, morseCodeH: {letter: "H", weigth: 86},
	morseCodeV: {letter: "V", weigth: 85}, morseCodeF: {letter: "F", weigth: 84}, morseCodeL: {letter: "L", weigth: 83},
	morseCodeP: {letter: "P", weigth: 82}, morseCodeJ: {letter: "J", weigth: 81}, morseCodeB: {letter: "B", weigth: 80},
	morseCodeX: {letter: "X", weigth: 79}, morseCodeC: {letter: "C", weigth: 78}, morseCodeY: {letter: "Y", weigth: 77},
	morseCodeZ: {letter: "Z", weigth: 76}, morseCodeQ: {letter: "Q", weigth: 75},
}

func Possibilities(signals string) []string {
	if signal, ok := morseDichotomicCodeMap[signals]; ok {
		return []string{signal.letter}
	}

	inputLength := len(signals)
	knownSignalsMap, hasKnownSignals := getKnownSignalsPositions(signals)
	if !hasKnownSignals {
		return getSignalsPossibilitiesWithInputLength(inputLength)
	}

	return getSignalsPossibilitiesFromKnownSignalPosition(knownSignalsMap, inputLength)
}

func getKnownSignalsPositions(signals string) (map[int]string, bool) {
	knownSignalsMap := map[int]string{}

	signalsArray := []rune(signals)
	for pos, s := range signalsArray {
		signal := string(s)
		if signal != unknownSignal {
			knownSignalsMap[pos] = signal
		}
	}

	if len(knownSignalsMap) == 0 {
		return nil, false
	}

	return knownSignalsMap, true
}

func getSignalsPossibilitiesWithInputLength(inputLength int) []string {
	nodes := make([]morseCode, 0)
	for morseCode, threeNode := range morseDichotomicCodeMap {
		if len(morseCode) == inputLength {
			nodes = append(nodes, threeNode)
		}
	}

	return getOrderedPossibilities(nodes)
}

func getSignalsPossibilitiesFromKnownSignalPosition(knownSignalsMap map[int]string, inputLength int) []string {
	nodes := make([]morseCode, 0)
	for morseKey, threeNode := range morseDichotomicCodeMap {
		morseKeyArray := []rune(morseKey)
		if hasAllSignalsInPlace(morseKeyArray, knownSignalsMap) && satisfyInputLength(morseKey, inputLength) {
			nodes = append(nodes, threeNode)
		}
	}

	return getOrderedPossibilities(nodes)
}

func hasAllSignalsInPlace(signals []rune, knownSignalsMap map[int]string) bool {
	for pos, knownSignal := range knownSignalsMap {
		if pos > len(signals)-1 || string(signals[pos]) != knownSignal {
			return false
		}
	}

	return true
}

func satisfyInputLength(morseKey string, limitedLength int) bool {
	if limitedLength == maxSignalsLength {
		return true
	}

	return len(morseKey) == limitedLength
}

func getOrderedPossibilities(node []morseCode) []string {
	sort.Slice(node, func(i, j int) bool {
		return node[i].weigth > node[j].weigth
	})

	possibilities := make([]string, len(node))
	for i, v := range node {
		possibilities[i] = v.letter
	}

	return possibilities
}

func main() {
	fmt.Println(Possibilities(".?"))
	fmt.Println(Possibilities("?."))
	fmt.Println(Possibilities("?"))
}
