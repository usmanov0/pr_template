package repo

import (
	"database/sql"
	"errors"
	"project-template/internal/entity"
)

type UserRepo interface {
	Create(user *entity.User) error
	GetById(id int) (*entity.User, error)
	GetAll() (*[]entity.User, error)
	PutUser(user *entity.User) (*entity.User, error)
	Delete(id int) error
}

type SqlUserRepo struct {
	db *sql.DB
}

func NewSQLUserRepo(db *sql.DB) *SqlUserRepo {
	return &SqlUserRepo{db: db}
}

func (r *SqlUserRepo) Create(user *entity.User) error {
	_, err := r.db.Exec("INSERT INTO users(Id, UserName,Email,Pinfl, PassWord VALUES (?, ?, ?, ?, ?)",
		user.Id, user.UserName, user.Email, user.PinFl, user.PassWord)
	return err
}

func (r *SqlUserRepo) GetById(id int) (*entity.User, error) {
	row := r.db.QueryRow("SELECT Id, UserName, PassWord, PinFl, Email FROM users WHERE Id = ?", id)
	var user entity.User
	err := row.Scan(&user.Id, &user.UserName, &user.PassWord, &user.PinFl, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *SqlUserRepo) GetAll() (*[]entity.User, error) {
	rows, err := r.db.Query("SELECT Id, UserName, PassWord, PinFl, Email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.Id, &user.UserName, &user.PassWord, &user.PinFl, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return &users, nil
}

func (r *SqlUserRepo) PutUser(user *entity.User) (*entity.User, error) {
	_, err := r.db.Exec("UPDATE users SET UserName=?, PassWord=?, PinFl=?, Email=? WHERE Id=?",
		user.UserName, user.PassWord, user.PinFl, user.Email, user.Id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *SqlUserRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE Id=?", id)
	return err
}
