package algorithm

import (
	"fmt"
	"sort"
	"strings"
	"backend/models"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

type similarity struct {
	index int
	sim   float64
}

func searchQuestion(pattern string, questions []models.ChatGPT, param string) ([]string, []int) {
	pattern = strings.ToLower(pattern)


	if len(questions) == 0 {
		return []string{"Tidak ada pertanyaan yang cocok"}, []int{-1}
	}

	if param == "kmp" {
		for i, q := range questions {
			if KnuthMorrisPratt(strings.ToLower(q.Pertanyaan), pattern) != -1 {
				fmt.Println("DEBUG - KMP match ditemukan untuk:", q.Pertanyaan)
				return []string{q.Jawaban}, []int{i}
			}
		}
	} else if param == "bm" {
		for i, q := range questions {
			if BoyerMooreAlgorithm(strings.ToLower(q.Pertanyaan), pattern) != -1 {
				fmt.Println("DEBUG - BM match ditemukan untuk:", q.Pertanyaan)
				return []string{q.Jawaban}, []int{i}
			}
		}
	}

	return lv(pattern, questions)
}

func lv(pattern string, questions []models.ChatGPT) ([]string, []int) {
	if len(questions) == 0 {
		return []string{"Tidak ada pertanyaan yang cocok"}, []int{-1}
	}

	var matrixSim []similarity
	fmt.Println("DEBUG - Memulai perhitungan Levenshtein untuk:", pattern)

	for i, q := range questions {
		lowerPattern := strings.ToLower(pattern)
		lowerQuestion := strings.ToLower(q.Pertanyaan)

		dist := levenshtein.DistanceForStrings([]rune(lowerPattern), []rune(lowerQuestion), levenshtein.DefaultOptions)
		minLength := float64(len(lowerPattern))
		if len(lowerQuestion) < len(lowerPattern) {
			minLength = float64(len(lowerQuestion))
		}

		score := 1.0 - (float64(dist) / minLength)

		fmt.Printf("DEBUG - Pertanyaan: '%s' | Distance: %d | Similarity: %.2f\n", lowerQuestion, dist, score)

		matrixSim = append(matrixSim, similarity{i, score})
	}

	sort.Slice(matrixSim, func(i, j int) bool {
		return matrixSim[i].sim > matrixSim[j].sim
	})

	fmt.Println("DEBUG - Hasil Sorting Similarity:")
	for _, sim := range matrixSim {
		fmt.Printf("DEBUG - Index: %d | Similarity: %.2f | Pertanyaan: %s\n", sim.index, sim.sim, questions[sim.index].Pertanyaan)
	}

	if matrixSim[0].sim >= 0.8 {
		fmt.Println("DEBUG - Jawaban langsung ditemukan:", questions[matrixSim[0].index].Jawaban)
		return []string{questions[matrixSim[0].index].Jawaban}, []int{matrixSim[0].index}
	}


	var recommendations []string
	var indices []int

	maxRecommendations := 3
	if len(matrixSim) < maxRecommendations {
		maxRecommendations = len(matrixSim)
	}

	for i := 0; i < maxRecommendations; i++ {
		if matrixSim[i].sim >= 0.4 {
			recommendations = append(recommendations, questions[matrixSim[i].index].Pertanyaan)
			indices = append(indices, matrixSim[i].index)
		}
	}

	if len(recommendations) > 0 {
		recommendationText := "Apakah yang Anda maksud: " + strings.Join(recommendations, " atau ") + "?"
		fmt.Println("DEBUG - Rekomendasi diberikan:", recommendationText)
		return []string{recommendationText}, indices
	}

	fmt.Println("DEBUG - Tidak ada pertanyaan yang cukup mirip.")
	return []string{"Tidak ada pertanyaan yang cocok"}, []int{-1}
}
