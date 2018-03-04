package main

import (
	"fmt"
	"strings"
	"time"

	"database/sql"
	"encoding/csv"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/robbert229/jwt"
)

type Blog struct {
	visitas           string
	cliques           string
	compartilhamentos string
}

func main() {
	// Jwt()
	loadCSV()
}

func loadCSVsaveMysql(registro Blog) {
	db, err := sql.Open("mysql", "root:root@/golang")
	check(err)

	defer db.Close()

	_, er := db.Exec("INSERT INTO golang (visitas, cliques, compartilhamentos) VALUES(?,?,?)", registro.visitas, registro.cliques, registro.compartilhamentos)
	check(er)

	fmt.Println("Tudo certo!")
}

func loadCSV() {
	file, err := os.Open("stats.csv")
	check(err)

	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '|'

	records, err := reader.ReadAll()
	check(err)

	for i := 1; i < len(records); i++ {
		v := records[i][1]
		cl := records[i][2]
		com := records[i][3]
		dados := Blog{visitas: v, cliques: cl, compartilhamentos: com}
		loadCSVsaveMysql(dados)
	}
}

func Jwt() {
	secret := "ThisIsMySuperSecret"
	algorithm := jwt.HmacSha256(secret)

	claims := jwt.NewClaim()
	claims.Set("Role", "Admin")
	claims.SetTime("exp", time.Now().Add(time.Minute))

	token, err := algorithm.Encode(claims)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Token: %s\n", token)

	if algorithm.Validate(token) != nil {
		panic(err)
	}

	loadedClaims, err := algorithm.Decode(token)
	if err != nil {
		panic(err)
	}

	role, err := loadedClaims.Get("Role")
	if err != nil {
		panic(err)
	}

	roleString, ok := role.(string)
	if !ok {
		panic(err)
	}

	if strings.Compare(roleString, "Admin") == 0 {
		//user is an admin
		fmt.Println("User is an admin")
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
