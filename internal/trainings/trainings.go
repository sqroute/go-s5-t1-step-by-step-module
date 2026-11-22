package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {

	parts := strings.Split(datastring, ",")

	if len(parts) != 3 {
		return errors.New("incorrect datastring format")
	}

	steps, err := strconv.Atoi(parts[0])

	if err != nil {
		return fmt.Errorf("error parse steps from datastring: %w", err)
	}

	if steps <= 0 {
		return errors.New("steps must be positive")
	}

	t.Steps = steps
	t.TrainingType = parts[1]

	duration, err := time.ParseDuration(parts[2])

	if err != nil {
		return fmt.Errorf("error parse duration from datastring: %w", err)
	}

	if duration <= 0 {
		return errors.New("duration must be positive")
	}

	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {

	distance := spentenergy.Distance(t.Steps, t.Height)
	meenSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":

		calories, err = spentenergy.RunningSpentCalories(
			t.Steps,
			t.Weight,
			t.Height,
			t.Duration,
		)

		if err != nil {
			return "", err
		}

	case "Ходьба":

		calories, err = spentenergy.WalkingSpentCalories(
			t.Steps,
			t.Weight,
			t.Height,
			t.Duration,
		)

		if err != nil {
			return "", err
		}

	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	msg := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, meenSpeed, calories)

	return msg, nil

}
