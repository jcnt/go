package main

import (
	"fmt"
	"math/rand"
	"os"
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

var owners []Owner
var pets []Pet
var db *sql.DB

var cfg = mysql.Config{
	User:                 os.Getenv("DBUSER"),
	Passwd:               os.Getenv("DBPASS"),
	Net:                  "tcp",
	Addr:                 "petclinic-db-mysql.petclinic:3306",
	DBName:               "petclinic",
	AllowNativePasswords: true,
}

func main() {

	loadOwners()
	loadPets()
	a, b := randData()

	fmt.Println(owners)
	fmt.Println()
	fmt.Println(pets)
	fmt.Println()
	fmt.Println(a)
	fmt.Println(b)

}

func loadOwners() {
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

func insertOwners() {

}

func randData() (Owner, Pet) {

	t := time.Now()
	rand.NewSource(t.UnixNano())

	first := []string{
		"Jack",
		"Harry",
		"Jacob",
		"Charlie",
		"Thomas",
		"George",
		"Oscar",
		"James",
		"William",
		"Jake",
		"Connor",
		"Callum",
		"Jacob",
		"Kyle",
		"Joe",
		"Reece",
		"Rhys",
		"Charlie",
		"Damian",
		"Noah",
		"Liam",
		"Mason",
		"Jacob",
		"William",
		"Ethan",
		"Michael",
		"Alexander",
		"James",
		"Daniel",
		"James",
		"John",
		"Robert",
		"Michael",
		"William",
		"David",
		"Richard",
		"Joseph",
		"Charles",
		"Thomas",
		"Larry",
	}

	last := []string{
		"JONES",
		"WILLIAMS",
		"TAYLOR",
		"BROWN",
		"DAVIES",
		"EVANS",
		"WILSON",
		"THOMAS",
		"JOHNSON",
		"ROBERTS",
		"ROBINSON",
		"THOMPSON",
		"WRIGHT",
		"WALKER",
		"WHITE",
		"EDWARDS",
		"HUGHES",
		"GREEN",
		"HALL",
		"LEWIS",
		"HARRIS",
		"CLARKE",
		"PATEL",
		"JACKSON",
		"WOOD",
		"TURNER",
		"MARTIN",
		"COOPER",
		"HILL",
		"WARD",
		"MORRIS",
		"MOORE",
		"CLARK",
		"LEE",
		"KING",
		"BAKER",
		"HARRISON",
		"MORGAN",
		"ALLEN",
		"JOHNSON",
	}

	street := []string{
		"Third",
		"First",
		"Fourth",
		"Park",
		"Fifth",
		"Main",
		"Sixth",
		"Oak",
		"Seventh",
		"Pine",
		"Maple",
		"Cedar",
		"Eighth",
		"Elm",
		"View",
		"Washington",
		"Ninth",
		"Lake",
		"Hill",
		"High",
		"Station",
		"Main",
		"Park",
		"Church",
		"London",
		"Victoria",
		"Green",
		"Manor",
		"Park",
		"Queens",
		"Grange",
		"Kings",
		"Kingsway",
		"Windsor",
		"Alexander",
		"York",
		"Broadway",
		"Springfield",
		"Richmond",
		"Taylor",
	}

	city := []string{
		"Sun Prairie",
		"McFarland",
		"Windsor",
		"Monona",
		"Waunakee",
		"Madison",
		"Sun Prairie",
		"McFarland",
		"Windsor",
		"Monona",
		"Waunakee",
		"Madison",
		"Sun Prairie",
		"McFarland",
		"Windsor",
		"Monona",
		"Waunakee",
		"Madison",
		"Sun Prairie",
		"McFarland",
		"Windsor",
		"Monona",
		"Waunakee",
		"Madison",
		"Sun Prairie",
		"McFarland",
		"Windsor",
		"Monona",
		"Waunakee",
		"Madison",
		"Sun Prairie",
		"McFarland",
		"Windsor",
		"Monona",
		"Waunakee",
		"Madison",
		"Sun Prairie",
		"McFarland",
		"Windsor",
		"Monona",
	}

	pet := []string{
		"Milo",
		"Teddy",
		"Buddy",
		"Alfie",
		"Max",
		"Charlie",
		"Bailey",
		"Reggie",
		"Bruno",
		"Hugo",
		"Luna",
		"Bella",
		"Lola",
		"Poppy",
		"Daisy",
		"Coco",
		"Ruby",
		"Molly",
		"Rosie",
		"Willow",
		"Milo",
		"Teddy",
		"Buddy",
		"Alfie",
		"Max",
		"Charlie",
		"Bailey",
		"Reggie",
		"Bruno",
		"Hugo",
		"Luna",
		"Bella",
		"Lola",
		"Poppy",
		"Daisy",
		"Coco",
		"Ruby",
		"Molly",
		"Rosie",
		"Willow",
	}

	rowner := Owner{
		strconv.Itoa(rand.Intn(10) + 20),
		first[rand.Intn(len(first))],
		last[rand.Intn(len(last))],
		strconv.Itoa(rand.Intn(100)) + " " + strconv.Itoa(rand.Intn(len(street))),
		city[rand.Intn(len(city))],
		strconv.Itoa(rand.Intn(1000000) + 1000000),
	}

	rpet := Pet{
		strconv.Itoa(rand.Intn(10) + 20),
		pet[rand.Intn(len(pet))],
		"2006-10-10",
		strconv.Itoa(rand.Intn(5) + 1),
		strconv.Itoa(rand.Intn(10) + 20),
	}

	return rowner, rpet
}
