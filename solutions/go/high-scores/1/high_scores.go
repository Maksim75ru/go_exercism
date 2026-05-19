package highscores

import (
    "sort"
)

type HighScores struct{
    scores []int
}

// NewHighScores returns a new HighScores object.
func NewHighScores(scores []int) *HighScores {
	return &HighScores{scores: scores}
}

// Scores returns all the scores.
func (s *HighScores) Scores() []int {
    return s.scores
}

// Latest returns the latest (last) score.
func (s *HighScores) Latest() int {
	return s.scores[len(s.scores) - 1]
}

// PersonalBest returns the best (highest) score.
func (s *HighScores) PersonalBest() int {
	max := s.scores[0]
    for _, n := range s.scores {
        if n > max {
            max = n
        }
    }
    return max
}

// TopThree returns the top three scores.
func (s *HighScores) TopThree() []int {
    if len(s.scores) == 0 {
        return []int{}
    }
    
    // Копируем и сортируем по убыванию
    sorted := make([]int, len(s.scores))
    copy(sorted, s.scores)
    sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
    
    // Возвращаем первые 3 или меньше
    if len(sorted) > 3 {
        return sorted[:3]
    }
    return sorted
}
