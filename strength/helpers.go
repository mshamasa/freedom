package strength

import "sort"

func (a StrengthList) Len() int { return len(a) }
func (a StrengthList) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a StrengthList) Less(i, j int) bool {
  return a[i].Date > a[j].Date
}

func sortWorkouts(workouts []Workout) []Strength {
  workoutsMap := make(map[int64][]Workout)
  for i := 0; i < len(workouts); i++ {
    date := workouts[i].Date
    workoutsMap[date] = append(workoutsMap[date], workouts[i])
  }

  strengthList := StrengthList{}
  for k, v := range workoutsMap {
    strength := Strength {
      Date: k,
      WorkoutList: v,
    }
    strengthList = append(strengthList, strength)
  }

  sort.Sort(StrengthList(strengthList))
  return strengthList
}

func generateWorkouts(date int64, userId string, workouts []Workout) []Workout {
  workoutList := WorkoutList{}
  for i := 0; i < len(workouts); i++ {
    workout := Workout {
      UserId: userId,
      Exercise: workouts[i].Exercise,
      Weight: workouts[i].Weight,
      Sets: workouts[i].Sets,
      Reps: workouts[i].Reps,
      Completed: workouts[i].Completed,
      Date: date,
    }
    workoutList = append(workoutList, workout)
  }
  return workoutList
}

