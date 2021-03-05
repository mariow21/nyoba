package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"mario/nyoba/config"
	"mario/nyoba/models"
	"time"

	// "gopkg.in/check.v1"
	// "github.com/google/go-querystring/query"
)

const(
	table = "user_mario"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll ...
func GetAll(ctx context.Context) ([]models.User, error){
	var users []models.User

	db,err := config.MySQL()

	if err != nil {
		log.Fatal("cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC",table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next(){
		var user models.User
		// var CreatedAt, updatedAt string

		if err = rowQuery.Scan(&user.ID,
			&user.NIP,
			&user.Name,
			&user.Position,
			&user.CreatedAt,
			&user.UpdatedAt,
			); err != nil {
			return nil, err
		}

		// user.CreatedAt, err = time.Parse(layoutDateTime,CreatedAt)

		// if err != nil {
		// 	log.Fatal(err)
		// }

		// user.UpdatedAt, err= time.Parse(layoutDateTime, updatedAt)

		// if err != nil {
		// 	return nil, err
		// }

		users = append(users, user)
	}

	return users, nil

}

// Insert ...
func Insert(ctx context.Context, use models.User) error{
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("cant Connect", err)
	}

	queryText  := fmt.Sprintf("INSERT INTO %v (nip, name, position, created_at, updated_at)values(%v,'%v', '%v', '%v', '%v')",table,
		use.NIP,
		use.Name,
		use.Position,
		time.Now().Format(layoutDateTime),
		time.Now().Format(layoutDateTime))

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Update ...
func Update(ctx context.Context, use models.User) error  {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("cant Connect", err)
	}

	queryText  := fmt.Sprintf("UPDATE %v set nip = %d, name = '%s', position = '%s', updated_at = '%v' where id = '%d'",table,
		use.NIP,
		use.Name,
		use.Position,
		time.Now().Format(layoutDateTime),
		use.ID,
	)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete ...
func Delete(ctx context.Context, use models.User) error {
	
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("cant Connect", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = '%d'", table, use.ID)
 
    s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println("check")
	if check == 0 {
		return errors.New("id tidak ada")
	}

	return nil
}