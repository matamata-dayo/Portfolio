package database

import (
	"fmt"
	"net/http"
	"strconv"
)

var query string
var firstFlag int = 0

func CreateQuery(r *http.Request) string {

	// postからデータを取得
	team := r.FormValue("team")
	position := r.FormValue("position")
	country := r.FormValue("country")
	age := r.FormValue("age")
	number := r.FormValue("number")
	captain := r.FormValue("captain")
	foot := r.FormValue("foot")
	height := r.FormValue("height")

	// クエリを初期化
	query = "SELECT name, team, position, country, age, number, foot, height FROM players INNER JOIN teams ON players.team_id = teams.team_id"
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
	case "Manchester United":
		query = query + "team = 'Manchester United'"
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
	case "Ireland":
		query = query + "country = 'アイルランド'"
	case "America":
		query = query + "country = 'アメリカ'"
	case "Algeria":
		query = query + "country = 'アルジェリア'"
	case "Argentina":
		query = query + "country = 'アルゼンチン'"
	case "Albania":
		query = query + "country = 'アルバニア'"
	case "Italy":
		query = query + "country = 'イタリア'"
	case "England":
		query = query + "country = 'イングランド'"
	case "Wales":
		query = query + "country = 'ウェールズ'"
	case "Ukraine":
		query = query + "country = 'ウクライナ'"
	case "Uruguay":
		query = query + "country = 'ウルグアイ'"
	case "Egypt":
		query = query + "country = 'エジプト'"
	case "Netherlands":
		query = query + "country = 'オランダ'"
	case "Cameroon":
		query = query + "country = 'カメルーン'"
	case "Gabon":
		query = query + "country = 'ガボン'"
	case "Ghana":
		query = query + "country = 'ガーナ'"
	case "Korea":
		query = query + "country = '韓国'"
	case "Guinea":
		query = query + "country = 'ギニア'"
	case "Greece":
		query = query + "country = 'ギリシャ'"
	case "Croatia":
		query = query + "country = 'クロアチア'"
	case "Columbia":
		query = query + "country = 'コロンビア'"
	case "Switzerland":
		query = query + "country = 'スイス'"
	case "Sweden":
		query = query + "country = 'スウェーデン'"
	case "Scotland":
		query = query + "country = 'スコットランド'"
	case "Spain":
		query = query + "country = 'スペイン'"
	case "Slovakia":
		query = query + "country = 'スロバキア'"
	case "Senegal":
		query = query + "country = 'セネガル'"
	case "Denmark":
		query = query + "country = 'デンマーク'"
	case "Germany":
		query = query + "country = 'ドイツ'"
	case "Japan":
		query = query + "country = '日本'"
	case "Norway":
		query = query + "country = 'ノルウェー'"
	case "France":
		query = query + "country = 'フランス'"
	case "Brazil":
		query = query + "country = 'ブラジル'"
	case "Belgium":
		query = query + "country = 'ベルギー'"
	case "Portugal":
		query = query + "country = 'ポルトガル'"
	case "Mali":
		query = query + "country = 'マリ'"
	case "Morocco":
		query = query + "country = 'モロッコ'"
	}

	// 年齢条件
	valueCheck(age)

	switch age {
	case "under19":
		query = query + "age <= 19"
	case "between20-25":
		query = query + "age >= 20 AND age <= 25"
	case "between26-29":
		query = query + "age >= 26 AND age <= 29"
	case "over30":
		query = query + "age >= 30"
	}

	// 背番号条件
	valueCheck(number)

	for i := 1; i < 100; i++ {
		if number == strconv.Itoa(i) {
			query = query + "number = " + strconv.Itoa(i)
			break
		}
	}

	// キャプテン条件
	valueCheck(captain)

	if captain == "captain" {
		query = query + "captain = 1"
	}

	// 利き足条件
	valueCheck(foot)

	switch foot {
	case "right":
		query = query + "foot = 'right'"
	case "left":
		query = query + "foot = 'left'"
	}

	// 身長条件
	valueCheck(height)

	switch height {
	case "under169":
		query = query + "height <= 169"
	case "between170-179":
		query = query + "height >= 170 AND height <= 179"
	case "between180-189":
		query = query + "height >= 180 AND height <= 189"
	case "between190-199":
		query = query + "height >= 190 AND height <= 199"
	case "over200":
		query = query + "height >= 200"
	}

	fmt.Println("作成されたSQLのクエリ:" + query)

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
