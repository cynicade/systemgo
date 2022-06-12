package processselector

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/ktr0731/go-fuzzyfinder"
)

func parseLoad(s string) LoadT {
	// loaded, not-found, bad-setting, error, masked
	switch s {
	case LoadT.String(Loaded):
		return Loaded
	case LoadT.String(NotFound):
		return NotFound
	case LoadT.String(BadSetting):
		return BadSetting
	case LoadT.String(Error):
		return Error
	case LoadT.String(Masked):
		return Masked
	}
	return UndefinedLoad
}

func parseActive(s string) ActiveT {
	// active, reloading, inactive, failed, activating, deactivating
	switch s {
	case ActiveT.String(Active):
		return Active
	case ActiveT.String(Reloading):
		return Reloading
	case ActiveT.String(Inactive):
		return Inactive
	case ActiveT.String(Failed):
		return Failed
	case ActiveT.String(Activating):
		return Activating
	case ActiveT.String(Deactivating):
		return Deactivating
	}
	return UndefinedActive
}

func getUnits() []Unit {
	out, err := exec.Command("systemctl", "list-units").Output()
	if err != nil {
		log.Fatal(err)
	}

	var units = []Unit{}
	outString := string(out)
	outSplit := strings.Split(outString, "\n")

	for idx, unit := range outSplit {
		if idx == 0 {
			continue
		}

		if idx >= len(outSplit)-6 {
			break
		}

		unitClean := strings.TrimSpace(unit)
		re := regexp.MustCompile("[ \t]+")
		unitSplit := re.Split(unitClean, 5)
		if len(unitSplit) < 5 {
			continue
		}
		// unit load active sub description
		temp := Unit{unitSplit[0], parseLoad(unitSplit[1]), parseActive(unitSplit[2]), unitSplit[3], unitSplit[4]}
		units = append(units, temp)
	}

	return units
}

func UnitSelector() Unit {
	units := getUnits()

	idx, findErr := fuzzyfinder.Find(
		units,
		func(i int) string {
			return units[i].Name
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}

			boldGreen := color.New(color.FgGreen, color.Bold).SprintFunc()
			boldYellow := color.New(color.FgYellow, color.Bold).SprintFunc()
			boldRed := color.New(color.FgRed, color.Bold).SprintFunc()

			var loadStr string
			var activeStr string

			load := units[i].Load
			if load == Loaded {
				loadStr = boldGreen(load.String())

			} else if load == Error {
				loadStr = boldRed(load.String())
			}

			active := units[i].Active
			if active == Active {
				activeStr = boldGreen(active.String())
			} else if active == Inactive {
				activeStr = boldYellow(active.String())
			} else if active == Failed {
				activeStr = boldRed(active.String())
			}

			return fmt.Sprintf("Unit: %s\nLoad: %s\nActive: %s\nSub: %s\nDescription: %s",
				units[i].Name,
				loadStr,
				activeStr,
				units[i].Sub,
				units[i].Description,
			)
		}),
	)
	if findErr != nil {
		log.Fatal(findErr)
	}

	return units[idx]
}
