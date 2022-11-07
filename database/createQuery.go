package database

import (
	"fmt"
	"net/http"
)

var query string
var firstFlag int = 0

func CreateQuery(r *http.Request) string {

	// postからデータを取得
	team := r.FormValue("team")
	position := r.FormValue("position")
	country := r.FormValue("country")
	age := r.FormValue("age")
	score := r.FormValue("score")
	assist := r.FormValue("assist")
	foot := r.FormValue("foot")
	height := r.FormValue("height")

	// クエリを初期化
	query = "SELECT name, team, position, country, age, score, assist, foot, height FROM players INNER JOIN teams ON players.team_id = teams.team_id"
	firstFlag = 0

	/* 検索条件によってクエリを生成 */

	// チーム条件
	valueCheck(team)

	switch team {
	case "Arsenal":
		query = query + "team = 'Arsenal'"
	case "Aston Villa":
		query = query + "team = 'Aston Villa'"
	case "Bournemouth":
		query = query + "team = 'Bournemouth'"
	case "Brighton & Hove Albion":
		query = query + "team = 'Brighton & Hove Albion'"
	case "Burnley":
		query = query + "team = 'Burnley'"
	case "Chelsea":
		query = query + "team = 'Chelsea'"
	case "Crystal Palace":
		query = query + "team = 'Crystal Palace'"
	case "Everton":
		query = query + "team = 'Everton'"
	case "Leicester City":
		query = query + "team = 'Leicester City'"
	case "Liverpool":
		query = query + "team = 'Liverpool'"
	case "Manchester City":
		query = query + "team = 'Manchester City'"
	case "Manchester-United":
		query = query + "team = 'Manchester-United'"
	case "Newcastle United":
		query = query + "team = 'Newcastle United'"
	case "Norwich City":
		query = query + "team = 'Norwich City'"
	case "Sheffield United":
		query = query + "team = 'Sheffield United'"
	case "Southampton":
		query = query + "team = 'Southampton'"
	case "Tottenham Hotspur":
		query = query + "team = 'Tottenham Hotspur'"
	case "Watford":
		query = query + "team = 'Watford'"
	case "West Ham United":
		query = query + "team = 'West Ham United'"
	case "Wolverhampton Wanderers":
		query = query + "team = 'Wolverhampton Wanderers'"
	}

	// ポジション条件
	valueCheck(position)

	switch position {
	case "fw":
		query = query + "position = 'FW'"
	case "mf":
		query = query + "position = 'MF'"
	case "df":
		query = query + "position = 'DF'"
	case "gk":
		query = query + "position = 'GK'"
	}

	// 国籍条件
	valueCheck(country)

	switch country {
	//
	}

	// 年齢条件
	valueCheck(age)

	switch age {
	//
	}

	// 得点条件
	valueCheck(score)

	switch score {
	//
	}

	// アシスト条件
	valueCheck(assist)

	switch assist {
	//
	}

	// 利き足条件
	valueCheck(foot)

	switch foot {
	//
	}

	// 身長条件
	valueCheck(height)

	switch height {
	//
	}

	fmt.Println(query)

	return query
}

func valueCheck(data interface{}) {
	if data != "" {
		if firstFlag == 0 {
			query = query + " WHERE "
			firstFlag = 1
		} else {
			query = query + " AND "
		}
	}
}
