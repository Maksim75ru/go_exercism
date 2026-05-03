package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var totalCount int
    for dayCount := 0; dayCount < len(birdsPerDay); dayCount++ {
        totalCount += birdsPerDay[dayCount]
    }
    return totalCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var totalCount int

    startDay := week * 7 - 7
    endDay :=  week * 7 - 1

    for day := startDay; day <= endDay; day++ {
        totalCount += birdsPerDay[day]
    }

    return totalCount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for dayCount := 0; dayCount < len(birdsPerDay); dayCount+=2 {
        birdsPerDay[dayCount] += 1
    }    
	return birdsPerDay
}
