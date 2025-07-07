package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Variable represents a <Variable ... /> entry
type Variable struct {
	Name    string
	Type    string
	Default string
}

// Action represents an <Action ...>...</Action> entry
type Action struct {
	Name   string
	Params []string
}

// Trigger represents a <Trigger ...>...</Trigger> entry
type Trigger struct {
	Name    string
	Actions []Action
}

type SpawnAction struct {
	Line        int
	TriggerName string
	Wave        string
	Unit        string
	SpawnPoint  string
	Owner       string
}

func main() {
	file, err := os.Open("2nd.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var triggerName string
	var lineNum int = 0
	var results []SpawnAction
	inTrigger := false
	inAction := false
	currentActionLine := 0
	currentParams := []string{}
	currentActionName := ""

	for scanner.Scan() {
		line := scanner.Text()
		lineNum++
		trim := strings.TrimSpace(line)

		if strings.HasPrefix(trim, "<Trigger ") {
			inTrigger = true
			// Имя триггера может быть на этой же строке
			if strings.Contains(trim, "<Name>") && strings.Contains(trim, "</Name>") {
				start := strings.Index(trim, "<Name>") + len("<Name>")
				end := strings.Index(trim, "</Name>")
				if start < end {
					triggerName = trim[start:end]
				}
			}
			continue
		}
		if inTrigger && strings.HasPrefix(trim, "<Name>") && strings.HasSuffix(trim, "</Name>") && !inAction {
			triggerName = strings.TrimSuffix(strings.TrimPrefix(trim, "<Name>"), "</Name>")
			continue
		}
		if inTrigger && strings.HasPrefix(trim, "<Action ") {
			inAction = true
			currentActionLine = lineNum
			currentParams = []string{}
			currentActionName = ""
			// Check if <Name>...</Name> is in the same line (rare, but for robustness)
			if strings.Contains(trim, "<Name>") && strings.Contains(trim, "</Name>") {
				start := strings.Index(trim, "<Name>") + len("<Name>")
				end := strings.Index(trim, "</Name>")
				if start < end {
					currentActionName = trim[start:end]
				}
			}
			continue
		}
		if inAction && strings.HasPrefix(trim, "<Name>") && strings.HasSuffix(trim, "</Name>") {
			currentActionName = strings.TrimSuffix(strings.TrimPrefix(trim, "<Name>"), "</Name>")
			continue
		}
		if inAction && strings.HasPrefix(trim, "<Param>") && strings.HasSuffix(trim, "</Param>") {
			param := strings.TrimSuffix(strings.TrimPrefix(trim, "<Param>"), "</Param>")
			currentParams = append(currentParams, param)
			continue
		}
		if inAction && strings.HasPrefix(trim, "</Action>") {
			if currentActionName == "a_spawnUnitGroupToZone" {
				wave := ""
				unit := ""
				spawnPoint := ""
				owner := ""
				if len(currentParams) > 1 {
					wave = currentParams[1]
				}
				if len(currentParams) > 3 {
					unit = currentParams[3]
				}
				if len(currentParams) > 5 {
					spawnPoint = currentParams[5]
				}
				if len(currentParams) > 6 {
					owner = currentParams[6]
				}
				results = append(results, SpawnAction{
					Line:        currentActionLine,
					TriggerName: triggerName,
					Wave:        wave,
					Unit:        unit,
					SpawnPoint:  spawnPoint,
					Owner:       owner,
				})
			}
			inAction = false
			continue
		}
		if inTrigger && strings.HasPrefix(trim, "</Trigger>") {
			inTrigger = false
			continue
		}
	}

	fmt.Printf("%-5s| %-20s| %-10s| %-20s| %-15s| %-15s\n", "LINE", "TRIGGER", "WAVE", "UNIT TO SPAWN", "SPAWN POINT", "APPLIED OWNER")
	for _, r := range results {
		fmt.Printf("%04d | %-20s| %-10s| %-20s| %-15s| %-15s\n", r.Line, r.TriggerName, r.Wave, r.Unit, r.SpawnPoint, r.Owner)
	}
} 
