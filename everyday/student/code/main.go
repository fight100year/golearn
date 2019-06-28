package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	mysqlURL string
	initDB   bool
	help     bool
)

func init() {
	flag.BoolVar(&help, "h", false, "the help")
	flag.BoolVar(&initDB, "r", false, "create/reset db")
	flag.StringVar(&mysqlURL, "mysql", "root:123456@tcp(172.168.10.119)/?charset=utf8",
		"data source name, eg: root:123456@/tcp(ip:port)student?charset=utf8")
}

func main() {
	flag.Parse()
	if help {
		usage()
		return
	}

	if initDB {
		resetDB()
		return
	}

	fmt.Println("start...")

	startServer()

	// errfun := func(w http.ResponseWriter, r *http.Request) {
	//     io.WriteString(w, "usage: ip:port/[add|delete|update|query]")
	// }
	//
	// addfun := func(w http.ResponseWriter, r *http.Request) {
	//     b, _ := ioutil.ReadAll(r.Body)
	//     defer r.Body.Close()
	//     io.WriteString(w, string(b))
	// }
	//
	// delfun := func(w http.ResponseWriter, r *http.Request) {
	//     b, _ := ioutil.ReadAll(r.Body)
	//     defer r.Body.Close()
	//     io.WriteString(w, string(b))
	// }
	//
	// updatefun := func(w http.ResponseWriter, r *http.Request) {
	//     b, _ := ioutil.ReadAll(r.Body)
	//     defer r.Body.Close()
	//     io.WriteString(w, string(b))
	// }
	//
	// queryfun := func(w http.ResponseWriter, r *http.Request) {
	//     b, _ := ioutil.ReadAll(r.Body)
	//     defer r.Body.Close()
	//     io.WriteString(w, string(b))
	// }
	//
	// http.HandleFunc("/", errfun)
	// http.HandleFunc("/add", addfun)
	// http.HandleFunc("/delete", delfun)
	// http.HandleFunc("/update", updatefun)
	// http.HandleFunc("/query", queryfun)
	//
	// log.Fatal(http.ListenAndServe("172.17.0.2:9000", nil))
}

func usage() {
	fmt.Fprintf(os.Stdout, `student version: student/0.0.1
	usage: student [-h|reset] [-mysql dsn]
	
	options:
`)
	flag.PrintDefaults()
}

// Service 提供http能力、db能力
type Service struct {
	db *sql.DB
}

func resetDB() error {
	fmt.Println("resetDB start...")
	db, err := sql.Open("mysql", mysqlURL)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return err
	}

	const (
		sqldb    = `CREATE DATABASE IF NOT EXISTS student CHARACTER SET = utf8mb4;`
		sqltable = `CREATE TABLE IF NOT EXISTS student.info(
id INT NOT NULL AUTO_INCREMENT COMMENT '自增id, 用于标识一个学生',
name VARCHAR(30) NOT NULL COMMENT '名字',
gender INT NOT NULL DEFAULT 1 COMMENT '1 男 2 女 ',
age INT NOT NULL COMMENT '年龄',
class VARCHAR(30) NOT NULL COMMENT '班级',
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 123456 COMMENT '学生信息表';`
		sqltuncate = `TRUNCATE TABLE student.info`
	)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if _, err := tx.ExecContext(ctx, sqldb); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("rollback failed: %v", rollbackErr)
		}
		log.Fatalf("create database failed :%v", err)

		return err
	}

	if _, err := tx.ExecContext(ctx, sqltable); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("rollback failed: %v", rollbackErr)
		}
		log.Fatalf("create table failed :%v", err)

		return err
	}

	if _, err := tx.ExecContext(ctx, sqltuncate); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("rollback failed: %v", rollbackErr)
		}
		log.Fatalf("clean table failed :%v", err)

		return err
	}

	if err := tx.Commit(); err != nil {
		log.Fatalf("commit failed: %v", err)
	}

	fmt.Println("resetDB success...")
	return nil
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	db := s.db
	switch r.URL.Path {
	default:
		http.Error(w, "not found", http.StatusNotFound)
		return
	case "/add":
		name := r.FormValue("name")
		gender, _ := strconv.Atoi(r.FormValue("gender"))
		age, _ := strconv.Atoi(r.FormValue("age"))
		class := r.FormValue("class")
		if gender != 1 && gender != 2 {
			gender = 1
		}

		fmt.Printf("add: [name:%v, gender:%v, age:%v, class:%v]\n", name,
			gender, age, class)

		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
			w.Write([]byte(fmt.Sprintf("add failed: %v", err)))

			return
		}

		_, err = tx.ExecContext(r.Context(), `INSERT INTO 
student.info(name, gender, age, class) 
VALUES(?,?,?,?);`, name, gender, age, class)
		if err != nil {
			log.Fatalf("add failed: %v", err)
			fmt.Println(err)

			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Fatalf("rollback failed: %v", rollbackErr)
			}

			w.Write([]byte(fmt.Sprintf("add failed: %v", err)))

			return
		}

		var newID int
		err = tx.QueryRowContext(r.Context(),
			`SELECT last_insert_id() FROM student.info;`).Scan(&newID)
		if err != nil {
			log.Fatalf("add failed: %v", err)
			fmt.Println(err)

			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Fatalf("rollback failed: %v", rollbackErr)
			}

			w.Write([]byte(fmt.Sprintf("add failed: %v", err)))

			return
		}

		if err := tx.Commit(); err != nil {
			log.Fatalf("commit failed: %v", err)
		}

		w.Write([]byte(strconv.Itoa(newID)))

		return
	case "/delete":
		return
	case "/update":
		return
	case "/query":
		return
	case "/status":
		ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
		defer cancel()

		err := db.PingContext(ctx)
		if err != nil {
			http.Error(w, fmt.Sprintf("db down: %v", err), http.StatusFailedDependency)
			return
		}

		w.Write([]byte("mysql server is alive"))
		w.WriteHeader(http.StatusOK)
		return
	}
}

func startServer() {
	db, err := sql.Open("mysql", mysqlURL)
	if err != nil {
		log.Fatal(err)
		return
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	s := &Service{db: db}

	http.ListenAndServe("172.17.0.2:9000", s)
}
