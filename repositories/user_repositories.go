package repositories

import (
	"database/sql"
	"golang-beginner-22/models"

	"github.com/google/uuid"
)

type UserRepositoryDB struct {
	db *sql.DB
}

func NewUserRepositoryDB(db *sql.DB) *UserRepositoryDB {
	return &UserRepositoryDB{db: db}
}

func (r *UserRepositoryDB) Create(name, email, password, token string) (*models.User, error) {
	var user models.User

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	// SQL query to create a new user
	sqlStatement := `INSERT INTO users (name, email, password, token) VALUES ($1, $2, $3, $4) RETURNING id, email, name, token`
	err = tx.QueryRow(sqlStatement, name, email, password, token).Scan(&user.ID, &user.Email, &user.Name, &user.Token)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryDB) Login(email, password string) (*models.User, error) {
	var user models.User

	// SQL query to check if the user exists and password is correct
	sqlStatement := `SELECT id, email FROM users WHERE email = $1 AND password = $2`
	err := r.db.QueryRow(sqlStatement, email, password).Scan(&user.ID, &user.Email)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	user.Token = uuid.NewString()
	updateStatement := `UPDATE users SET token = $1 WHERE id = $2`
	_, err = r.db.Exec(updateStatement, user.Token, user.ID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryDB) GetAll() ([]models.User, error) {
	rows, err := r.db.Query("SELECT id, email, name, status FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.Status)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepositoryDB) GetById(id int) (*models.User, error) {
	var user models.User

	sqlStatement := `SELECT id, email, name, password, token FROM users WHERE id = $1`
	err := r.db.QueryRow(sqlStatement, id).Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.Token)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}
