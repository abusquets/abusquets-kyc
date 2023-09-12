package core_repositories

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	_uuid "github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	errors "github.com/abusquets/ab-kyc/internal/app/errors"
	core_entities "github.com/abusquets/ab-kyc/internal/core/domain/entities"
	core_ports "github.com/abusquets/ab-kyc/internal/core/domain/ports/repositories"
)

type userRepository struct {
	core_ports.IUserRepository
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) core_ports.IUserRepository {
	return &userRepository{db: db}
}

func (ur userRepository) GetByIDMem(uuid string) (*core_entities.User, *errors.AppError) {
	userUuid, err := _uuid.Parse(uuid)
	if err != nil {
		return nil, errors.NewNotFoundError("User not found")
	}

	user := core_entities.User{
		ID:        1,
		Uuid:      userUuid,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		Email:     "abusquets@gmail.com",
	}
	fmt.Println(user)
	return &user, nil
}

func (ur userRepository) GetByIDOld(uuid string) (*core_entities.User, *errors.AppError) {

	var err error

	var query = `SELECT * FROM kyc.users WHERE uuid = $1`
	var user core_entities.User
	err = ur.db.Get(&user, query, uuid)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("User not found")
		} else {
			log.Error().Err(err).Msg("Error while scanning User")
			return nil, errors.NewUnexpectedError("Unexpected database error")
		}
	}
	return &user, nil
}

func (ur userRepository) GetByID(uuid string) (*core_entities.User, *errors.AppError) {
	user := new(core_entities.User)

	query := sq.Select(
		"*",
	).From(
		"users",
	).Where(
		sq.Eq{"uuid": uuid},
	)
	stmt, args, err := query.ToSql()
	if err != nil {
		return nil, errors.NewUnexpectedError("SQL error: GET USER")
	}
	log.Debug().Msg(stmt)

	err = ur.db.Get(&user, stmt, args...)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("User not found")
		} else {
			log.Error().Err(err).Msg("Error while scanning user")
			return nil, errors.NewUnexpectedError("Unexpected database error")
		}
	}
	return user, nil

}

func (ur userRepository) Create(userData core_ports.CreateUserInDTO) (user *core_entities.User, error *errors.AppError) {
	user = new(core_entities.User)
	stmt, args, err := sq.
		Insert("users").
		PlaceholderFormat(sq.Dollar).
		Columns("email", "password", "is_active", "is_admin").
		Values(userData.Email, userData.Password, userData.IsActive, userData.IsAdmin).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		log.Error().Err(err).Msg("SQL error: CREATE USER")
		return nil, errors.NewUnexpectedError("SQL error: CREATE USER")
	}
	log.Debug().Msg(stmt)

	err = ur.db.QueryRowx(stmt, args...).Scan(&user.ID, &user.Uuid, &user.Email, &user.Password, &user.IsActive, &user.IsAdmin, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			msg := fmt.Sprintf("DB error: the User %s already exists", userData.Email)
			log.Error().Err(err).Msg(msg)
			return nil, errors.NewDuplicatedError(msg)
		} else {
			log.Error().Err(err).Msg("DB error: CREATE USER")
			return nil, errors.NewUnexpectedError("DB error: CREATE USER")
		}
	}
	return user, nil
}
