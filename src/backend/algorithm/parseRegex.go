package algorithm

import (
	"fmt"
	"regexp"
	"strings"
	"backend/models"
)

func ParseInput(input string, listOfQuestion []models.ChatGPT, param string) []string {
	patternDate := `(?i)^(Hari\s\d{1,2}/\d{1,2}/\d{4}|\d{1,2}/\d{4})\s*\??$`
	regDate := regexp.MustCompile(patternDate)
	
	patternMath := `^(Hasil dari\s*|Hasil\s*)?-?\s*(\(\s*)*\s*-?\s*\d+(\.\d+)?(\s*\))*\s*(\s*[-+*/]\s*(\(\s*)*\s*-?\s*(\(\s*)*\d+(\.\d+)?(\s*\))*\s*)*\s*\??$`
	regCal := regexp.MustCompile(patternMath)
	
	patternAdd := `(?i)^tambahkan pertanyaan\s*(.+)\s*dengan jawaban\s*(.+)$`
	regAdd := regexp.MustCompile(patternAdd)
	
	patternDel := `(?i)^hapus pertanyaan\s*(.+)$`
	regDel := regexp.MustCompile(patternDel)	

	fmt.Println("DEBUG - ParseInput - input:", input)
	fmt.Println("DEBUG - ParseInput - param:", param)
	for _, q := range listOfQuestion {
		fmt.Printf("DEBUG - Pertanyaan di DB: %s | Jawaban: %s\n", q.Pertanyaan, q.Jawaban)
	}

	if regDate.MatchString(input) {
		getInput := getDate(input)
		date := Calendar(getInput)
		fmt.Println("DEBUG - masuk date")
		return []string{"kalender", "Hari " + date}
	} else if regCal.MatchString(input) {
		getInput := getCalculator(input)
		math := Calculator(getInput)
		fmt.Println("DEBUG - masuk math")
		return []string{"kalkulator", "Hasilnya adalah " + math}
	} else if regAdd.MatchString(input) {
		question, answer := extractQuestionAnswer(input)
		fmt.Println("DEBUG - masuk add")

		if question == "" || answer == "" {
			return []string{"error", "Format pertanyaan tidak valid"}
		}
		return []string{"tambah", question, answer}
	} else if regDel.MatchString(input) {
		question := getQuestionDeleteCommand(input)
		fmt.Println("DEBUG - masuk del")

		if question == "" {
			return []string{"error", "Format pertanyaan tidak valid"}
		}
		return []string{"hapus", question}
	} else {
		fmt.Println("DEBUG - masuk pertanyaan")
		question, index := searchQuestion(input, listOfQuestion, param)

		fmt.Println("DEBUG - hasil pencarian:", question, index)

		if len(index) == 0 || index[0] == -1 {
			return []string{"jawaban", "Saya tidak dapat menjawab pertanyaan Anda."}
		} else if len(question) == 1 {
			return []string{"jawaban", question[0]}
		} else {
			answer := "Apakah yang Anda maksud: " + strings.Join(question, " atau ") + "?"
			return []string{"rekomendasi", answer}
		}
	}
}
	

func extractQuestionAnswer(input string) (string, string) {
	pattern := regexp.MustCompile(`(?i)^tambahkan pertanyaan\s(.+)\sdengan jawaban\s(.+)$`)
	matches := pattern.FindStringSubmatch(input)

	if len(matches) < 3 {
		fmt.Println("Regex tidak cocok dengan input:", input)
		return "", ""
	}

	question := strings.TrimSpace(matches[1])
	answer := strings.TrimSpace(matches[2])

	return question, answer
}

func getQuestionDeleteCommand(input string) string {
	reg := regexp.MustCompile(`(?i)hapus pertanyaan (.*)`)
	matches := reg.FindStringSubmatch(input)

	if len(matches) < 2 {
		fmt.Println("Regex tidak cocok untuk hapus pertanyaan:", input)
		return ""
	}

	question := strings.TrimSpace(matches[1])
	return question
}



func getDate(input string) string {
	patt := "^(?i)(Hari\\s)?(\\d{1,2}/\\d{1,2}/\\d{4})\\??$"
	reg := regexp.MustCompile(patt)
	matches := reg.FindStringSubmatch(input)
	if len(matches) == 3 {
		return matches[2]
	}
	return ""
}

func getCalculator(input string) string {
	patt := "^(?i)Hasil dari (.+)|([0-9]+\\s*[-+*/%^]\\s*[0-9]+.*)$"
	reg := regexp.MustCompile(patt)
	matches := reg.FindStringSubmatch(input)
	if len(matches) == 3 {
		if len(matches[1]) > 0 {
			if matches[1][len(matches[1])-1] == '?' {
				return matches[1][:len(matches[1])-1]
			} else {
				return matches[1]
			}
		} else {
			if matches[2][len(matches[2])-1] == '?' {
				return matches[2][:len(matches[2])-1]
			} else {
				return matches[2]
			}
		}
	}
	return ""
}