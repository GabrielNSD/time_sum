package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type TimeSum struct {
	latestAdded  []int
	totalSeconds int
	isSumming    bool
}

func (ts *TimeSum) addTime(timeStr string) error {
	// Try HH:MM:SS format first
	re := regexp.MustCompile(`^(\d{1,2}):(\d{2}):(\d{2})$`)
	matches := re.FindStringSubmatch(timeStr)
	if matches != nil {
		return ts.parseTimeWithSeconds(matches[1], matches[2], matches[3], timeStr)
	}

	// Try HH:MM or MM:SS format
	re = regexp.MustCompile(`^(\d{1,2}):(\d{2})$`)
	matches = re.FindStringSubmatch(timeStr)
	if matches != nil {
		// Check if it's MM:SS format (first number < 60) or HH:MM format
		firstNum, err := strconv.Atoi(matches[1])
		if err == nil && firstNum < 60 {
			// Likely MM:SS format
			return ts.parseMinutesSeconds(matches[1], matches[2], timeStr)
		} else {
			// Likely HH:MM format
			return ts.parseTimeWithMinutes(matches[1], matches[2], timeStr)
		}
	}

	return fmt.Errorf("invalid time format. Use HH:MM:SS, HH:MM, or MM:SS (e.g., 02:30:45, 02:30, 30:45)")
}

func (ts *TimeSum) parseTimeWithSeconds(hoursStr, minutesStr, secondsStr, timeStr string) error {
	hours, err := strconv.Atoi(hoursStr)
	if err != nil {
		return fmt.Errorf("invalid hours: %v", err)
	}

	minutes, err := strconv.Atoi(minutesStr)
	if err != nil {
		return fmt.Errorf("invalid minutes: %v", err)
	}

	seconds, err := strconv.Atoi(secondsStr)
	if err != nil {
		return fmt.Errorf("invalid seconds: %v", err)
	}

	if minutes >= 60 {
		return fmt.Errorf("minutes must be less than 60")
	}

	if seconds >= 60 {
		return fmt.Errorf("seconds must be less than 60")
	}

	if hours >= 24 {
		return fmt.Errorf("hours must be less than 24")
	}

	totalSeconds := hours*3600 + minutes*60 + seconds
	ts.totalSeconds += totalSeconds

	fmt.Printf("Added %s (total: %s)\n", timeStr, ts.formatTotal())
	return nil
}

func (ts *TimeSum) parseTimeWithMinutes(hoursStr, minutesStr, timeStr string) error {
	hours, err := strconv.Atoi(hoursStr)
	if err != nil {
		return fmt.Errorf("invalid hours: %v", err)
	}

	minutes, err := strconv.Atoi(minutesStr)
	if err != nil {
		return fmt.Errorf("invalid minutes: %v", err)
	}

	if minutes >= 60 {
		return fmt.Errorf("minutes must be less than 60")
	}

	if hours >= 24 {
		return fmt.Errorf("hours must be less than 24")
	}

	totalSeconds := hours*3600 + minutes*60
	ts.totalSeconds += totalSeconds

	fmt.Printf("Added %s (total: %s)\n", timeStr, ts.formatTotal())
	return nil
}

func (ts *TimeSum) parseMinutesSeconds(minutesStr, secondsStr, timeStr string) error {
	minutes, err := strconv.Atoi(minutesStr)
	if err != nil {
		return fmt.Errorf("invalid minutes: %v", err)
	}

	seconds, err := strconv.Atoi(secondsStr)
	if err != nil {
		return fmt.Errorf("invalid seconds: %v", err)
	}

	if minutes >= 60 {
		return fmt.Errorf("minutes must be less than 60")
	}

	if seconds >= 60 {
		return fmt.Errorf("seconds must be less than 60")
	}

	totalSeconds := minutes*60 + seconds
	ts.latestAdded = append(ts.latestAdded, totalSeconds)
	ts.totalSeconds += totalSeconds

	fmt.Printf("Added %s (total: %s)\n", timeStr, ts.formatTotal())
	return nil
}

func (ts *TimeSum) formatTotal() string {
	hours := ts.totalSeconds / 3600
	minutes := (ts.totalSeconds % 3600) / 60
	seconds := ts.totalSeconds % 60

	if hours > 0 {
		return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	} else if minutes > 0 {
		return fmt.Sprintf("%02d:%02d", minutes, seconds)
	} else {
		return fmt.Sprintf("00:%02d", seconds)
	}
}

func (ts *TimeSum) reset() {
	ts.totalSeconds = 0
	ts.isSumming = false
	fmt.Println("Time sum reset")
}

func (ts *TimeSum) showTotal() {
	if ts.totalSeconds == 0 {
		fmt.Println("No times added yet")
		return
	}
	fmt.Printf("Total time: %s\n", ts.formatTotal())
}

func (ts *TimeSum) undo() {
	if len(ts.latestAdded) > 0 {
		lastAdded := ts.latestAdded[len(ts.latestAdded)-1]
		ts.latestAdded = ts.latestAdded[:len(ts.latestAdded)-1]
		ts.totalSeconds -= lastAdded
	}
}

func printInstructions() {
	fmt.Println("Commands:")
	fmt.Println("  start - Start summing times")
	fmt.Println("  end   - End summing and show total")
	fmt.Println("  reset - Reset the sum")
	fmt.Println("  undo  - Undo the last time added")
	fmt.Println("  quit  - Exit the program")
	fmt.Println("  help  - Show this help")
	fmt.Println("Time formats:")
	fmt.Println("  HH:MM:SS (e.g., 02:30:45)")
	fmt.Println("  HH:MM    (e.g., 02:30)")
	fmt.Println("  MM:SS    (e.g., 30:45)")
}

func main() {
	timeSum := &TimeSum{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Time Sum CLI")
	printInstructions()
	fmt.Println()

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		switch strings.ToLower(input) {
		case "start":
			timeSum.totalSeconds = 0
			timeSum.isSumming = true
			fmt.Println("Started summing times. Enter times in HH:MM:SS, HH:MM, or MM:SS format or 'end' to finish.")

		case "end":
			if !timeSum.isSumming {
				fmt.Println("Not currently summing. Use 'start' to begin.")
				continue
			}
			timeSum.isSumming = false
			timeSum.showTotal()

		case "reset":
			timeSum.reset()

		case "quit", "exit":
			fmt.Println("Goodbye!")
			return

		case "undo":
			timeSum.undo()

		case "help":
			printInstructions()

		default:
			if timeSum.isSumming {
				// Try to parse as time
				if err := timeSum.addTime(input); err != nil {
					fmt.Printf("Error: %v\n", err)
					fmt.Println("Enter a valid time in HH:MM:SS, HH:MM, or MM:SS format or 'end' to finish.")
				}
			} else {
				fmt.Println("Not currently summing. Use 'start' to begin or 'help' for commands.")
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}
