package repositories

import (
	"database/sql"
	"golang-beginner-22/models"
	// "fmt"
)

type TodoRepositoryDB struct {
	DB *sql.DB
}

func NewTodoRepositoryDB(db *sql.DB) *TodoRepositoryDB {
	return &TodoRepositoryDB{DB: db}
}

func (repo *TodoRepositoryDB) Create(todoInput *models.Todos, token string) (*models.Todos, error) {
	tx, err := repo.DB.Begin()
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

	userSQLStatement := `SELECT id FROM users WHERE token = $1`
	var userID int
	err = tx.QueryRow(userSQLStatement, token).Scan(&userID)
	if err != nil {
		return nil, err
	}

	if userID == 0 {
		return nil, sql.ErrNoRows
	}

	todoInput.UserID = userID
	// Create a new todo item
	var todo models.Todos
	sqlStatement := `INSERT INTO todos (user_id, description) VALUES ($1, $2) RETURNING id, description, todo_status`
	err = tx.QueryRow(sqlStatement, todoInput.UserID, todoInput.Description).Scan(&todo.ID, &todo.Description, &todo.TodoStatus)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return &todo, err
}

func (repo *TodoRepositoryDB) GetTodos(token string) ([]models.Todos, error) {
	tx, err := repo.DB.Begin()
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

	userSqlStatement := `SELECT id FROM users WHERE token = $1`
	var userID int
	err = tx.QueryRow(userSqlStatement, token).Scan(&userID)
	if err != nil {
		return nil, err
	}

	userId := userID
	// Create a new todo item
	var todos []models.Todos
	sqlStatement := `SELECT id, description, todo_status FROM todos WHERE user_id = $1 AND todo_status != 'deleted'`
	rows, err := tx.Query(sqlStatement, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var todo models.Todos
		err := rows.Scan(&todo.ID, &todo.Description, &todo.TodoStatus)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return todos, err
}

func (repo *TodoRepositoryDB) GetCount(token string) (*models.Todos, error) {
	userStatetement := `SELECT id FROM users WHERE token = $1`
	var userID int
	err := repo.DB.QueryRow(userStatetement, token).Scan(&userID)
	if err != nil {
		return nil, err
	}
	notStartedCount := `SELECT count(*) FROM todos WHERE user_id = $1 AND todo_status = 'not_started'`
	doneCount := `SELECT count(*) FROM todos WHERE user_id = $1 AND todo_status = 'done'`

	var todo models.Todos
	err = repo.DB.QueryRow(notStartedCount, userID).Scan(&todo.TotalTodoNotStarted)
	if err != nil {
		return nil, err
	}
	err = repo.DB.QueryRow(doneCount, userID).Scan(&todo.TotalTodoDone)
	if err != nil {
		return nil, err
	}
	return &todo, err
}

func (repo *TodoRepositoryDB) Update(todo *models.Todos) (*models.Todos, error) {
	tx, err := repo.DB.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	sqlStatement := `UPDATE todos SET todo_status = $2 WHERE id = $1 RETURNING id, description, todo_status`
	var updatedTodo models.Todos
	err = tx.QueryRow(sqlStatement, todo.ID, todo.TodoStatus).Scan(&updatedTodo.ID, &updatedTodo.Description, &updatedTodo.TodoStatus)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return &updatedTodo, nil
}

func (repo *TodoRepositoryDB) Delete(id int) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	sqlStatement := `UPDATE todos SET todo_status = 'deleted' WHERE id = $1`
	_, err = tx.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
