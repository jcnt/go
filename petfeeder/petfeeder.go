package main

import (
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strconv"
	"time"

	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type Owner struct {
	Id      string
	First   string
	Last    string
	Address string
	City    string
	Phone   string
}

type Pet struct {
	Id    string
	Name  string
	Birth string
	Type  string
	Oid   string
}

var (
	owners []Owner
	pets   []Pet
	db     *sql.DB
)

var cfg = mysql.Config{
	User:                 os.Getenv("PETDBUSER"),
	Passwd:               os.Getenv("PETDBPASS"),
	Net:                  "tcp",
	Addr:                 "petclinic-db-mysql.petclinic:3306",
	DBName:               "petclinic",
	AllowNativePasswords: true,
}

func main() {

	maxRows, err := strconv.Atoi(os.Getenv("PETLIMIT"))
	if err != nil {
		fmt.Println("PETLIMIT ENV is missing")
	}

	interval, err := strconv.Atoi(os.Getenv("PETINT"))
	if err != nil {
		fmt.Println("PETINT ENV is missing")
	}

	for {

		loadOwners()
		loadPets()

		if maxRows == 0 || len(owners) < maxRows {
			soid, spid := getNextId()
			rowner, rpet := randData(soid, spid)
			insertData(rowner, rpet)
			time.Sleep(time.Duration(interval) * time.Second)
		} else {
			deleteRows(len(owners))
			time.Sleep(time.Duration(interval) * time.Second)
		}
	}

}

func loadOwners() {
	owners = []Owner{}
	pets = []Pet{}

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		fmt.Println("SQL Open failed")
	}

	defer db.Close()

	rows, err := db.Query("SELECT * from owners")
	if err != nil {
		fmt.Println("SQL SELECT failed")
	}

	defer rows.Close()

	for rows.Next() {
		var owner Owner
		if err := rows.Scan(&owner.Id, &owner.First, &owner.Last, &owner.Address, &owner.City, &owner.Phone); err != nil {
			fmt.Println("Read failed")
		}
		owners = append(owners, owner)
	}

}

func loadPets() {
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		fmt.Println("SQL Open failed")
	}

	defer db.Close()

	rows, err := db.Query("SELECT * from pets")
	if err != nil {
		fmt.Println("SQL SELECT failed")
	}

	defer rows.Close()

	for rows.Next() {
		var pet Pet
		if err := rows.Scan(&pet.Id, &pet.Name, &pet.Birth, &pet.Type, &pet.Oid); err != nil {
			fmt.Println("Read failed")
		}
		pets = append(pets, pet)
	}
}

func insertData(rowner Owner, rpet Pet) (sql.Result, sql.Result) {

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		fmt.Println("SQL Open failed")
	}

	defer db.Close()

	resultOwner, err := db.Exec("INSERT INTO owners VALUES (?, ?, ?, ?, ?, ?)",
		rowner.Id, rowner.First, rowner.Last, rowner.Address, rowner.City, rowner.Phone)
	if err != nil {
		fmt.Println("Owner Insert Failed")
	}

	resultPet, err := db.Exec("INSERT INTO pets VALUES (?,?,?,?,?)",
		rpet.Id, rpet.Name, rpet.Birth, rpet.Type, rpet.Oid)
	if err != nil {
		fmt.Println("Pet Insert Failed")
	}

	return resultOwner, resultPet

}

func getNextId() (string, string) {
	var oid, pid int

	maxOid, _ := strconv.Atoi(owners[len(owners)-1].Id)
	if maxOid == len(owners) {
		oid = maxOid
	} else {
		for _, v := range owners {
			id, _ := strconv.Atoi(v.Id)
			if id-oid == 1 {
				oid++
			}
		}
	}

	maxPid, _ := strconv.Atoi(pets[len(pets)-1].Id)
	if maxPid == len(pets) {
		pid = maxPid
	} else {
		for _, v := range pets {
			id, _ := strconv.Atoi(v.Id)
			if id-pid == 1 {
				pid++
			}
		}
	}

	finalOid := strconv.Itoa(oid + 1)
	finalPid := strconv.Itoa(pid + 1)

	return finalOid, finalPid

}

func deleteRows(row int) (sql.Result, sql.Result) {

	var rows []int
	var a, b sql.Result

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		fmt.Println("SQL Open failed")
	}

	defer db.Close()

	for i := 0; i <= row/5; {
		new := rand.Intn(row) + 1
		if !slices.Contains(rows, new) && new != 6 {
			rows = append(rows, new)
		}
		i = len(rows) + 1
	}

	for _, v := range rows {
		a, err = db.Exec("DELETE FROM pets WHERE owner_id= (?)", v)

		if err != nil {
			fmt.Println("Pet Delete Failed")
		}

		b, err = db.Exec("DELETE FROM owners WHERE id = (?)", v)

		if err != nil {
			fmt.Println("Owner Insert Failed")
		}

	}
	return a, b
}

func randData(soid string, spid string) (Owner, Pet) {

	t := time.Now()
	rand.NewSource(t.UnixNano())

	first := []string{"Jack", "Harry", "Jacob", "Charlie", "Thomas", "George", "Oscar", "James", "William", "Jake", "Connor", "Callum", "Jacob", "Kyle", "Joe", "Reece", "Rhys", "Charlie", "Damian", "Noah", "Liam", "Mason", "Jacob", "William", "Ethan", "Michael", "Alexander", "James", "Daniel", "James", "John", "Robert", "Michael", "William", "David", "Richard", "Joseph", "Charles", "Thomas", "Larry"}
	last := []string{"Jones", "Williams", "Taylor", "Brown", "Davies", "Evans", "Wilson", "Thomas", "Johnson", "Roberts", "Robinson", "Thompson", "Wright", "Walker", "White", "Edwards", "Hughes", "Green", "Hall", "Lewis", "Harris", "Clarke", "Patel", "Jackson", "Wood", "Turner", "Martin", "Cooper", "Hill", "Ward", "Morris", "Moore", "Clark", "Lee", "King", "Baker", "Harrison", "Morgan", "Allen", "Johnson"}
	street := []string{"Third", "First", "Fourth", "Park", "Fifth", "Main", "Sixth", "Oak", "Seventh", "Pine", "Maple", "Cedar", "Eighth", "Elm", "View", "Washington", "Ninth", "Lake", "Hill", "High", "Station", "Main", "Park", "Church", "London", "Victoria", "Green", "Manor", "Park", "Queens", "Grange", "Kings", "Kingsway", "Windsor", "Alexander", "York", "Broadway", "Springfield", "Richmond", "Taylor"}
	stType := []string{"St.", "Rd.", "Dr.", "Av.", "Exp."}
	city := []string{"Sun Prairie", "McFarland", "Windsor", "Monona", "Waunakee", "Madison"}
	pet := []string{"Milo", "Teddy", "Buddy", "Alfie", "Max", "Charlie", "Bailey", "Reggie", "Bruno", "Hugo", "Luna", "Bella", "Lola", "Poppy", "Daisy", "Coco", "Ruby", "Molly", "Rosie", "Willow", "Milo", "Teddy", "Buddy", "Alfie", "Max", "Charlie", "Bailey", "Reggie", "Bruno", "Hugo", "Luna", "Bella", "Lola", "Poppy", "Daisy", "Coco", "Ruby", "Molly", "Rosie", "Willow"}

	year, _, _ := time.Now().Date()

	var month string
	if m := rand.Intn(12) + 1; m < 10 {
		month = "0" + strconv.Itoa(m)
	} else {
		month = strconv.Itoa(m)
	}

	var day string
	var dmax int
	switch month {
	case "01", "03", "05", "07", "08", "10", "12":
		dmax = 31
	case "02":
		dmax = 28
	default:
		dmax = 30
	}
	if d := rand.Intn(dmax) + 1; d < 10 {
		day = "0" + strconv.Itoa(d)
	} else {
		day = strconv.Itoa(d)
	}

	birth := "20" + strconv.Itoa(rand.Intn(12)+year-2012) + "-" + month + "-" + day

	rowner := Owner{
		soid,
		first[rand.Intn(len(first))],
		last[rand.Intn(len(last))],
		strconv.Itoa(rand.Intn(5000)) + " " + street[rand.Intn(len(street))] +
			" " + stType[rand.Intn(len(stType))],
		city[rand.Intn(len(city))],
		strconv.Itoa(rand.Intn(999999999) + 6000000000),
	}

	rpet := Pet{
		spid,
		pet[rand.Intn(len(pet))],
		birth,
		strconv.Itoa(rand.Intn(5) + 1),
		soid,
	}

	return rowner, rpet
}
