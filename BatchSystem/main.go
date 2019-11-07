package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type task struct {
	duration int
	indices  []int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	duration := sc.Text()
	sc.Scan()
	user := sc.Text()

	var durations []int
	for _, d := range strings.Split(duration, ",") {
		di, _ := strconv.Atoi(d)
		durations = append(durations, di)
	}
	users := strings.Split(user, ",")

	// fmt.Println(durations)
	// fmt.Println(users)
	n := len(durations)
	var userTasks = map[string]task{}
	for i := 0; i < n; i++ {
		u := users[i]
		d := durations[i]
		if t, ok := userTasks[u]; ok {
			t.duration += d
			t.indices = append(t.indices, i)
			userTasks[u] = t
		} else {
			t := task{duration: d, indices: []int{i}}
			userTasks[u] = t
		}
	}
	// fmt.Println(userTasks)

	var tasks []task
	for _, t := range userTasks {
		tasks = append(tasks, t)
	}
	// fmt.Println(tasks)

	sort.Slice(tasks, func(i, j int) bool {
		if tasks[i].duration == tasks[j].duration {
			return tasks[i].indices[0] < tasks[j].indices[0]
		} else {
			return tasks[i].duration < tasks[j].duration
		}
	})
	// fmt.Println(tasks)

	var orders []int
	for _, t := range tasks {
		for _, i := range t.indices {
			orders = append(orders, i)
		}
	}

	fmt.Println(orders)
}
