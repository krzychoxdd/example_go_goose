package migrations

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/go-faker/faker/v4"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upSeedSchema, downSeedSchema)
}

var density = 25

func upSeedSchema(tx *sql.Tx) error {

	runSeeder(tx)

	return nil
}

type DBUser struct {
	UserID    string `faker:"uuid_digit"`
	Email     string `faker:"email"`
	Password  string `faker:"password"`
	FirstName string `faker:"first_name"`
	LastName  string `faker:"last_name"`
}

type DBCategory struct {
	Name string `faker:"word"`
}

type DBAd struct {
	Title       string  `faker:"sentence"`
	Description string  `faker:"paragraph"`
	Latitude    float32 `faker:"lat"`
	Longitude   float32 `faker:"long`
}

func runSeeder(tx *sql.Tx) {

	user := DBUser{}
	category := DBCategory{}
	ad := DBAd{}

	var categoryIDs []int64
	for u := 0; u <= density*100; u++ {
		err := faker.FakeData(&category)
		if err != nil {
			fmt.Println(err)
		}

		res, err := tx.Exec(fmt.Sprintf(`INSERT INTO categories(name) VALUES('%s');`,
			strings.Title(category.Name),
		))

		if err != nil {
			panic(err)
		}

		if val, ok := res.LastInsertId(); ok == nil {
			categoryIDs = append(categoryIDs, val)
		}
	}

	var userIDs []string

	for i := 0; i <= density*10; i++ {
		err := faker.FakeData(&user)
		if err != nil {
			fmt.Println(err)
		}

		_, err = tx.Exec(fmt.Sprintf(`INSERT INTO users(user_id, email, password, first_name, last_name) VALUES("%s", "%s", "%s", "%s", "%s");`,
			user.UserID, user.Email, user.Password, user.FirstName, user.LastName,
		))

		if err != nil {
			panic(err)
		}

		_, err = tx.Exec(fmt.Sprintf(`INSERT INTO user_email_confirmations(user_id, code) VALUES('%s', '%s');`,
			user.UserID, fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String()+user.UserID))),
		))

		if err != nil {
			panic(err)
		}

		userIDs = append(userIDs, user.UserID)
	}

	var adIDs []int64
	for _, user_id := range userIDs {
		for j := 0; j <= density*rand.Intn(30); j++ {
			err := faker.FakeData(&ad)
			if err != nil {
				fmt.Println(err)
			}

			randomCatID := rand.Intn(len(categoryIDs))
			catID := categoryIDs[randomCatID]

			res, err := tx.Exec(fmt.Sprintf(`INSERT INTO ad(category_id, creator_id, title, description, location_lat_long) VALUES('%d', '%s', '%s', '%s', '%s');`,
				catID, user_id, ad.Title, ad.Description, fmt.Sprintf("%f,%f", ad.Latitude, ad.Longitude),
			))

			if err != nil {
				panic(err)
			}

			if val, ok := res.LastInsertId(); ok == nil {
				adIDs = append(adIDs, val)
			}
		}
	}
}

func downSeedSchema(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
