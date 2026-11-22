package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func calculateSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 {
		return 0, errors.New("steps must be greater than zero")
	}

	if weight <= 0 {
		return 0, errors.New("weight must be greater than zero")
	}

	if height <= 0 {
		return 0, errors.New("height must be greater than zero")
	}

	if duration <= 0 {
		return 0, errors.New("duration must be greater than zero")
	}

	meanSpeed := MeanSpeed(steps, height, duration)

	calories := (weight * meanSpeed * duration.Minutes()) / float64(minInH)

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	calories, err := calculateSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	return calories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	calories, err := calculateSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {

	if duration <= 0 {
		return 0
	}

	durationHours := duration.Hours()
	distance := Distance(steps, height)

	return distance / durationHours

}

func Distance(steps int, height float64) float64 {
	stepLenght := height * stepLengthCoefficient
	return float64(steps) * stepLenght / float64(mInKm)
}
