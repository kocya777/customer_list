/*
モジュール名：

	main.go

機能名：

	顧客名簿管理システム

機能概要：

	顧客名簿を管理する。
	１．ログイン画面
		ログインIDとパスワードを入力して、ログイン後に顧客名簿一覧画面を表示する。
	２．顧客名簿一覧画面
		顧客名簿一覧の閲覧、削除・キーワード検索が出来る。
	３．詳細情報画面
		顧客名情報の新規登録、更新、閲覧が出来る。

作成日付：

	2024/03/29

作成者：

	k.f
*/

package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// メンバーID構造体
type USER_ID struct {
	ITEM0 int       // プライマリーキー
	ITEM1 string    // ユーザID
	ITEM2 string    // ユーザPW
	ITEM3 string    // 作成者
	ITEM4 time.Time // 作成日付
}

// 顧客名簿構造体
type MEMBER struct {
	Cnt         int
	Regflg      int
	Primary_key string
	Items       []CUSTOMER_MASTER
}
type CUSTOMER_MASTER struct {
	ITEM0  int       // プライマリーキー
	ITEM1  string    // 氏名
	ITEM2  string    // フリガナ
	ITEM3  string    // 住所
	ITEM4  int       // 年齢
	ITEM5  string    // 固定電話
	ITEM6  string    // 携帯電話
	ITEM7  string    // メールアドレス
	ITEM8  string    // 備考欄
	ITEM9  string    // 作成者
	ITEM10 time.Time // 作成日付
}

// メッセージ構造体
type Msg struct {
	Message string
}

var st_msg Msg = Msg{""}
var p_db *sql.DB = nil

/*
関数名：

	main

関数仕様：

	main処理を行う。

入力：

	無し

出力：

	無し
*/
func main() {
	port := "8080"
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js/"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img/"))))
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/management_list", handleList)
	http.HandleFunc("/management_editregistration", handleEdit)

	log.Print(http.ListenAndServe(":"+port, nil))
}

/*
関数名：

	handleIndex

関数仕様：

	ログイン画面処理を行う。

入力：

	w http.ResponseWriter
	r *http.Request)

出力：

	無し

エラー時：

	CALL panic()
*/
func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 入力フォームを返す
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, st_msg)
	}
	if r.Method == "POST" {
		// 必ず必要
		r.ParseForm()
		st_msg.Message = ""

		s_username := r.Form.Get("username")
		s_userpasswd := r.Form.Get("userpasswd")

		if len(s_username) == 0 || len(s_userpasswd) == 0 {
			st_msg.Message = "入力内容に誤りがあります。"
		}

		if st_msg.Message == "" {
			// データベースオープン
			var err error
			if p_db == nil {
				// データベースオープン
				p_db, err = sql.Open("mysql", "fujiwara:nice4970@tcp(localhost:3306)/golang_user?parseTime=true")
				if err != nil {
					log.Fatal(err)
				}
			}

			// 顧客登録数 読み込み
			s_sql := "SELECT count(*) FROM user_id where ITEM1='" + s_username + "' and ITEM2='" + s_userpasswd + "'"
			rows, err := p_db.Query(s_sql)
			//_ = rows
			if err != nil {
				p_db.Close()
				panic(err.Error())
			}
			defer rows.Close()

			var i_counter int
			for rows.Next() {
				err := rows.Scan(&i_counter)
				if err != nil {
					p_db.Close()
					panic(err.Error())
				}
			}

			if i_counter == 1 {
				http.Redirect(w, r, "/management_list", http.StatusMovedPermanently)
			} else {
				st_msg.Message = "ログイン名又は、パスワードが違います。"
				http.Redirect(w, r, "/login", http.StatusMovedPermanently)
			}
		} else {
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
		}
	}
}

/*
関数名：

	handleList

関数仕様：

	顧客名簿一覧画面処理を行う。

入力：

	w http.ResponseWriter
	r *http.Request)

出力：

	無し

エラー時：

	無し
*/
func handleList(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 顧客リスト読み込み
		st_m := sub_handleList(0, r)
		// 入力フォームを返す
		t, _ := template.ParseFiles("management_list.html")
		t.Execute(w, st_m)
	}
	if r.Method == "POST" {
		// 必ず必要
		r.ParseForm()

		action := r.Form.Get("action")
		switch action {
		// 新規登録
		case "0":
			http.Redirect(w, r, "/management_editregistration", http.StatusMovedPermanently)
		// キーワード検索
		case "1":
			st_m := sub_handleList(1, r)

			// 入力フォームを返す
			t, _ := template.ParseFiles("management_list.html")
			t.Execute(w, st_m)
		// 詳細
		case "2":
			primary_key := r.Form.Get("primary_key")

			str_url := "/management_editregistration?primary_key=" + primary_key
			http.Redirect(w, r, str_url, http.StatusMovedPermanently)
		// 削除
		case "3":
			sub_handleListDel(r)
			http.Redirect(w, r, "/management_list", http.StatusMovedPermanently)
		// 処理終了
		case "4":
			st_msg.Message = ""
			p_db.Close()
			p_db = nil
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
		}
	}
}

/*
関数名：

	bool sub_handleListDel

関数仕様：

	顧客削除処理

入力：

	r *http.Request

出力：

	bool true

エラー時：

	無し
*/
func sub_handleListDel(r *http.Request) bool {
	primary_key := r.Form.Get("primary_key")

	s_sql := "DELETE FROM customer_master WHERE ITEM0=?"
	rows, err := p_db.Prepare(s_sql)
	if err != nil {
		p_db.Close()
		panic(err.Error())
	}
	rows.Exec(primary_key)
	defer rows.Close()

	return true
}

/*
関数名：

	MEMBER sub_handleList

関数仕様：

	顧客名簿一覧画面のイベント処理（全顧客データ取得、キーワード検索、詳細）のデータ取得を行う。

入力：

	r *http.Request

出力：

	顧客名簿リスト

エラー時：

	CALL panic()
*/
func sub_handleList(serch_flg int, r *http.Request) MEMBER {
	var err error
	var s_sql string
	var s_keyword string
	var st_cm CUSTOMER_MASTER
	var i_counter int = 0
	var st_m MEMBER = MEMBER{}

	switch serch_flg {
	// 全顧客データ取得
	case 0:
		s_sql = "SELECT * FROM customer_master ORDER BY ITEM0 ASC;"
	// キーワード検索
	case 1:
		s_keyword = r.Form.Get("keyword")
		// キーワード検索入力なしか？
		if s_keyword == "" {
			s_sql = "SELECT * FROM customer_master ORDER BY ITEM0 ASC;"
		} else {
			// キーワード検索入力ありか？
			// ある場合、全ての項目に対して検索を行う。
			s_sql = "SELECT * FROM customer_master WHERE ITEM1 LIKE \"%" + s_keyword +
				"%\" or ITEM2 LIKE \"%" + s_keyword +
				"%\" or ITEM3 LIKE \"%" + s_keyword +
				"%\" or ITEM4 LIKE \"%" + s_keyword +
				"%\" or ITEM5 LIKE \"%" + s_keyword +
				"%\" or ITEM6 LIKE \"%" + s_keyword +
				"%\" or ITEM7 LIKE \"%" + s_keyword +
				"%\" or ITEM8 LIKE \"%" + s_keyword +
				"%\" or ITEM9 LIKE \"%" + s_keyword + "%\""
		}
	// 詳細
	case 2:
		primary_key := r.FormValue("primary_key")
		if primary_key == "" {
			return st_m
		}

		s_sql = "SELECT * FROM customer_master  WHERE ITEM0=" + primary_key
	}

	rows, err := p_db.Query(s_sql)
	if err != nil {
		p_db.Close()
		panic(err.Error())
	}

	for rows.Next() {
		err := rows.Scan(&st_cm.ITEM0, &st_cm.ITEM1, &st_cm.ITEM2, &st_cm.ITEM3, &st_cm.ITEM4, &st_cm.ITEM5, &st_cm.ITEM6, &st_cm.ITEM7, &st_cm.ITEM8, &st_cm.ITEM9, &st_cm.ITEM10)
		// 配列に新しく追加
		st_m.Items = append(st_m.Items, CUSTOMER_MASTER{st_cm.ITEM0, st_cm.ITEM1, st_cm.ITEM2, st_cm.ITEM3, st_cm.ITEM4, st_cm.ITEM5, st_cm.ITEM6, st_cm.ITEM7, st_cm.ITEM8, st_cm.ITEM9, st_cm.ITEM10})
		if err != nil {
			p_db.Close()
			panic(err.Error())
		}

		i_counter++
	}
	defer rows.Close()
	st_m.Cnt = i_counter

	return st_m
}

/*
関数名：

	handleEdit

関数仕様：

	詳細情報画面処理を行う。

入力：

	w http.ResponseWriter
	r *http.Request

出力：

	無し

エラー時：

	CALL panic()
*/
func handleEdit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 入力フォームを返す
		st_m := sub_handleList(2, r)
		if st_m.Cnt != 0 {
			st_m.Regflg = 1
			st_m.Primary_key = r.FormValue("primary_key")
		} else {
			st_m.Regflg = 0
		}

		// 入力フォームを返す
		t, _ := template.ParseFiles("management_editregistration.html")
		t.Execute(w, st_m)
	}
	if r.Method == "POST" {
		// 必ず必要
		r.ParseForm()

		action := r.Form.Get("action")
		switch action {
		// 顧客リスト読み込み
		case "0":
			http.Redirect(w, r, "/management_list", http.StatusMovedPermanently)
		// データ登録 or データ更新処理
		case "1":
			var s_sql string

			regflg := r.Form.Get("regflg")
			primary_key := r.Form.Get("primary_key")
			switch regflg {
			// データ登録
			case "0":
				s_sql = "INSERT INTO customer_master VALUES(" +
					"NULL," +
					"\"" + r.Form.Get("namae") + "\"," +
					"\"" + r.Form.Get("furigana") + "\"," +
					"\"" + r.Form.Get("zyusyo") + "\"," +
					r.Form.Get("nenrei") + "," +
					"\"" + r.Form.Get("koteidenwa") + "\"," +
					"\"" + r.Form.Get("keitaidenwa") + "\"," +
					"\"" + r.Form.Get("mailadoresu") + "\"," +
					"\"" + r.Form.Get("biko") + "\"," +
					"\"system\"," +
					"now());"
			// データ更新
			case "1":
				s_sql = "UPDATE customer_master " +
					"SET ITEM1=\"" + r.Form.Get("namae") + "\"," +
					"ITEM2=\"" + r.Form.Get("furigana") + "\"," +
					"ITEM3=\"" + r.Form.Get("zyusyo") + "\"," +
					"ITEM4=" + r.Form.Get("nenrei") + "," +
					"ITEM5=\"" + r.Form.Get("koteidenwa") + "\"," +
					"ITEM6=\"" + r.Form.Get("keitaidenwa") + "\"," +
					"ITEM7=\"" + r.Form.Get("mailadoresu") + "\"," +
					"ITEM8=\"" + r.Form.Get("biko") + "\"," +
					"ITEM9=\"system\"," +
					"ITEM10=now() WHERE ITEM0=" + primary_key
			}
			rows, err := p_db.Prepare(s_sql)
			if err != nil {
				p_db.Close()
				panic(err.Error())
			}
			rows.Exec()
			defer rows.Close()

			// 顧客リスト読み込み
			http.Redirect(w, r, "/management_list", http.StatusMovedPermanently)
		}
	}
}
