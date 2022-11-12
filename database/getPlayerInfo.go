package database

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
)

type playerInfo struct {
	Name     string
	Team     string
	Position string
	Country  string
	Age      int
	Number   int
	Foot     string
	Height   int
}

var SearchResult []playerInfo

func GetPlayerInfo(r *http.Request, db *sql.DB) {

	rows, err := db.QueryContext(context.Background(), createQuery(r))

	if err != nil {
		fmt.Println("選手情報の取得に失敗しました")
		panic(err.Error())
	} else {
		fmt.Println("選手情報の取得に成功しました")
	}

	defer rows.Close()

	SearchResult = []playerInfo{}

	for rows.Next() {
		var players playerInfo
		err := rows.Scan(
			&players.Name,
			&players.Team,
			&players.Position,
			&players.Country,
			&players.Age,
			&players.Number,
			&players.Foot,
			&players.Height,
		)
		if err != nil {
			panic(err.Error())
		}
		SearchResult = append(SearchResult, players)
	}

	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}
}

/*
クエリを作成する
*/
func createQuery(r *http.Request) string {

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
	query := "SELECT name, team, position, country, age, number, foot, height FROM players INNER JOIN teams ON players.team_id = teams.team_id"

	// クエリの条件が最初のものかを判定
	firstFlag := true

	/* 検索条件によってクエリを生成 */

	// チーム条件
	if team != "" {
		firstCheck(team, &firstFlag, &query)
		query = query + "team = '" + team + "'"
	}

	// ポジション条件
	if position != "" {
		firstCheck(position, &firstFlag, &query)
		query = query + "position = '" + position + "'"
	}

	// 国籍条件
	if country != "" {
		firstCheck(country, &firstFlag, &query)

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
	}

	// 年齢条件
	if age != "" {
		firstCheck(age, &firstFlag, &query)

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
	}

	// 背番号条件
	if number != "" {
		firstCheck(number, &firstFlag, &query)

		for i := 1; i < 100; i++ {
			if number == strconv.Itoa(i) {
				query = query + "number = " + strconv.Itoa(i)
				break
			}
		}
	}

	// キャプテン条件
	if captain != "" {
		firstCheck(captain, &firstFlag, &query)

		if captain == "captain" {
			query = query + "captain = 1"
		}
	}

	// 利き足条件
	if foot != "" {
		firstCheck(foot, &firstFlag, &query)
		query = query + "foot = '" + foot + "'"
	}

	// 身長条件
	if height != "" {
		firstCheck(height, &firstFlag, &query)

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
	}

	fmt.Println("作成されたSQLのクエリ:" + query)

	return query
}

/*
クエリの条件が複数ある場合、条件の間にANDを入れる
*/
func firstCheck(data interface{}, firstFlag *bool, query *string) {
	if *firstFlag {
		*query = *query + " WHERE "
		*firstFlag = false
	} else {
		*query = *query + " AND "
	}
}
