package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	// db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:23306)/fun_admin?charset=utf8mb4")
	db, err = sql.Open("mysql", "rdsroot:funip2019!db@#@tcp(mysql57.rdsmp9pujzy5quy.rds.bj.baidubce.com:3306)/fun_admin?charset=utf8mb4")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	// query := "select phone,user_ext,type from merchant_ip_coupons where ug4_id = 148 and status=1"
	query := "select a.phone,a.type,a.user_ext,a.updated_at,b.ext from " +
		"merchant_ip_coupons as a left join merchant_ip_prizes as b " +
		"on a.merchant_ip_prize_id = b.id where a.ug4_id = 148 and a.status=1"
	// query := "select a.phone,a.type,a.user_ext,a.updated_at,b.ext from " +
	// 	"merchant_ip_coupons as a left join merchant_ip_prizes as b " +
	// 	"on a.merchant_ip_prize_id = b.id where a.status=1 limit 0,10"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	type ABC struct {
		Phone     string
		Type      int64
		UserExt   sql.NullString
		UpdatedAt sql.NullString
		Ext       sql.NullString
	}

	for rows.Next() {
		var row ABC
		if err := rows.Scan(&row.Phone,
			&row.Type, &row.UserExt, &row.UpdatedAt, &row.Ext); err != nil {
			log.Fatal(err)
		}

		var t string
		if row.Type == 12 {
			t, _ = extractTitle(row.Ext.String)
		} else if row.Type == 1 {
			t = "电影票"
		}

		address, _ := extractAddress(row.UserExt.String)

		fmt.Printf("%s,%s,%s,%s\n", row.UpdatedAt.String, row.Phone, t, address)
	}
}

func extractAddress(s string) (string, error) {
	var ue interface{}
	if err := json.Unmarshal([]byte(s), &ue); err != nil {
		return "", err
	}
	var address string
	addr, ok := ue.(map[string]interface{})["address"]
	if ok {
		address = addr.(string)
	}
	return address, nil
}

func extractTitle(s string) (string, error) {
	var ue interface{}
	if err := json.Unmarshal([]byte(s), &ue); err != nil {
		return "", err
	}
	var title string
	addr, ok := ue.(map[string]interface{})["title"]
	if ok {
		title = addr.(string)
	}
	return title, nil
}

/**

02 24 34 34 33 39 34 35 63 38 66 66 62 64 35 33 31 65 63 32 32 62 30 36
31 64 64 61 39 31 37 63 31 62 5F 31 5F 30 02 FF 00 00 10 00 03 06 7D 01
9C 00 00 03 06 78 01 96 19 01 03 06 6E 01 94 19 01 03 06 60 01 95 19 01
0306540196190103063F019E190103062901A9190103060F01B9190103060101C319010305EE01D419010305DE01E719010305D501F419010305C9020919010305C2021C19010305C0022119010305BD023619010305BE024319010305C1025319010305C3025619010305CC026219010305D5026819010305DF026C19010305EC027019010306020270190103060F026E1901030625026C190103063C02671901030653025F1901030662025919010306780250190103068D02471901030698024219010306A3023B19010306B8023219010306C3022D19010306CB022819010306D2022219010306D4021F19010306D5021919010306D1021219010306C6020819010306BA020219010306A501FC190103068C01F6190103067901F3190103065A01F0190103063901F0190103062101F119010305FB01F519010305D501FB19010305AE020219010305960208190103056E02121901030548021D190103052E02271901030507023919010304E2024C19010304CA025B19010304A502721901030482028D190103046B02A2190103044902C3190103043402DF190103041B02FF190103041B02FF0000



02 24 34 34 33 39 34 35 63 38 66 66 62 64 35 33 31 65 63 32 32 62 30 36
31 64 64 61 39 31 37 63 31 62 5F 31 5F 30
02 FF 00 00 10 00 03 06 7D 01 9C 00 00 03 06780196190103066E01941901030660019519010306540196190103063F019E190103062901A9190103060F01B9190103060101C319010305EE01D419010305DE01E719010305D501F419010305C9020919010305C2021C19010305C0022119010305BD023619010305BE024319010305C1025319010305C3025619010305CC026219010305D5026819010305DF026C19010305EC027019010306020270190103060F026E1901030625026C190103063C02671901030653025F1901030662025919010306780250190103068D02471901030698024219010306A3023B19010306B8023219010306C3022D19010306CB022819010306D2022219010306D4021F19010306D5021919010306D1021219010306C6020819010306BA020219010306A501FC190103068C01F6190103067901F3190103065A01F0190103063901F0190103062101F119010305FB01F519010305D501FB19010305AE020219010305960208190103056E02121901030548021D190103052E02271901030507023919010304E2024C19010304CA025B19010304A502721901030482028D190103046B02A2190103044902C3190103043402DF190103041B02FF190103041B02FF0000


**/
