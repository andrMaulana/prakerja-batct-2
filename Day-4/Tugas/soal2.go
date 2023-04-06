package main

import "fmt"

type Student struct {
	Name  string
	Score []int
}

func (s *Student) GetAverageScore() float64 {
	var total int
	for _, score := range s.Score {
		total += score
	}
	return float64(total) / float64(len(s.Score))
}

func (s *Student) GetMinScore() int {
	min := s.Score[0]
	for _, score := range s.Score[1:] {
		if score < min {
			min = score
		}
	}
	return min
}

func (s *Student) GetMaxScore() int {
	max := s.Score[0]
	for _, score := range s.Score[1:] {
		if score > max {
			max = score
		}
	}
	return max
}
func main() {
	students := []Student{
		{Name: "Rizki", Score: []int{80, 75, 90}},
		{Name: "Kobar", Score: []int{70, 85, 65}},
		{Name: "Ismai", Score: []int{75, 80, 70}},
		{Name: "Umam", Score: []int{90, 95, 85}},
		{Name: "Yopan", Score: []int{60, 65, 70}},
	}

	// hitung skor rata-rata
	var total float64
	for _, student := range students {
		total += student.GetAverageScore()
	}
	averageScore := total / float64(len(students))
	fmt.Println("Skor rata-rata:", averageScore)

	// cari siswa dengan skor minimum
	var minScore int
	var minScoreStudent Student
	for _, student := range students {
		if student.GetMinScore() < minScore || minScore == 0 {
			minScore = student.GetMinScore()
			minScoreStudent = student
		}
	}
	fmt.Println("Siswa dengan skor minimum:", minScoreStudent.Name)

	// cari siswa dengan skor maksimum
	var maxScore int
	var maxScoreStudent Student
	for _, student := range students {
		if student.GetMaxScore() > maxScore {
			maxScore = student.GetMaxScore()
			maxScoreStudent = student
		}
	}
	fmt.Println("Siswa dengan skor maksimum:", maxScoreStudent.Name)
}
