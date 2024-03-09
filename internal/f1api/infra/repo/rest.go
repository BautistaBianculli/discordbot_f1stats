package repo

import (
	"BotDiscordGO/internal/f1api/infra/domain"
	"BotDiscordGO/internal/server/infra/config"
	"encoding/json"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"log"
	"net/http"
)

const (
	maxDiscordLenght = 2000
	errorMaxLenght   = "The table is longer but I can't show the whole message because it exceeds the discord message limit."
)

type F1Repository struct {
	Config *config.Config
	Client config.HttpClient
}

func (f *F1Repository) GetDriverTable(y string) string {
	url := fmt.Sprintf(f.Config.DriverTableUrl, y)

	resp, err := http.Get(url)

	if err != nil {
		log.Println(err)
		return "No hay data para ese año"
	}

	defer resp.Body.Close()

	var f1Driver domain.F1Drivers

	err = json.NewDecoder(resp.Body).Decode(&f1Driver)
	if err != nil {
		log.Println(err)
		return ""
	}
	if len(f1Driver.Data.StandingsTable.StandingsLists) == 0 {
		return fmt.Sprintf("No data from year %s", y)
	}

	t := table.NewWriter()
	t.AppendHeader(table.Row{"Position", "Pilot", "Constructor", "Points", "Wins"})

	var rows []table.Row

	for _, driverStanding := range f1Driver.Data.StandingsTable.StandingsLists[0].DriverStandings {
		row := table.Row{
			driverStanding.Position,
			fmt.Sprintf("%s %s", driverStanding.Driver.GivenName, driverStanding.Driver.FamilyName),
			driverStanding.Constructors[0].Name,
			driverStanding.Points,
			driverStanding.Wins,
		}
		rows = append(rows, row)
	}
	t.AppendRows(rows)

	msg := fmt.Sprintf("Year: %s\nRaces: %s\n```\n%s\n```", f1Driver.Data.StandingsTable.Season, f1Driver.Data.StandingsTable.StandingsLists[0].Round, t.Render())
	if len(msg) > maxDiscordLenght {
		t.ResetRows()
		rows = []table.Row{}
		for i, driverStanding := range f1Driver.Data.StandingsTable.StandingsLists[0].DriverStandings {
			row := table.Row{
				driverStanding.Position,
				fmt.Sprintf("%s %s", driverStanding.Driver.GivenName, driverStanding.Driver.FamilyName),
				driverStanding.Constructors[0].Name,
				driverStanding.Points,
				driverStanding.Wins,
			}
			rows = append(rows, row)
			if i == 15 {
				t.AppendRows(rows)
				msg = fmt.Sprintf("Year: %s\nRaces: %s\n```\n%s\n%s\n```", f1Driver.Data.StandingsTable.Season, f1Driver.Data.StandingsTable.StandingsLists[0].Round, t.Render(), errorMaxLenght)
				break
			}
		}
	}
	return msg
}
