package repository

import (
	"context"
	"database/sql"
	"github.com/chahar4/aura/internal/entity"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Init() error {
	query := `create table if not exists users (
	id serial primary key,
	username varchar(100),
	email varchar(100),
	password varchar(250),
	created_time timestamp
	)`
	_, err := r.db.Exec(query)
	return err
}

func (r *UserRepository) AddUser(c context.Context, user *entity.User) (*entity.User, error) {
	query := "INSERT INTO users(username , email , password) VALUES ($1 , $2 , $3) returning id"

	var lastIdInserted int
	err := r.db.QueryRowContext(c, query, user.Username, user.Email, user.Password).Scan(&lastIdInserted)
	if err != nil {
		return &entity.User{}, err
	}
	user.ID = lastIdInserted
	return user, nil
}

func (r *UserRepository) GetUserByEmail(c context.Context, email string) (*entity.User, error) {

	query := "SELECT * FROM users WHERE email = $1"
	var user entity.User
	err := r.db.QueryRowContext(c, query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAllUsersByGuildID(c context.Context, guildID int) (*[]entity.User, error) {
	query := `SELECT u.* FROM users u JOIN user_guild ug ON u.id = ug.user_id WHERE ug.guild_id = $1;`
	rows, err := r.db.QueryContext(c, query, guildID)
	if err != nil {
		return nil, err
	}
	var users []entity.User
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return &users, nil
}
