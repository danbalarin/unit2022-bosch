package auth

import (
	"github.com/jackc/pgconn"
	"github.com/pkg/errors"
	"github.com/unit2022-bosch/teapot/backend/internal/entity"
	"gorm.io/gorm"
)

type authGormRepository struct {
	db *gorm.DB
}

func NewAuthGormRepository(db *gorm.DB) IAuthRepository {
	return &authGormRepository{
		db: db,
	}
}

func (repo authGormRepository) insertUser(user *entity.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return ErrUserAlreadyExists
			}
		}

		return errors.Wrap(err, "failed to insert user")
	}
	return nil
}

func (repo authGormRepository) findUser(id uint) (*entity.User, error) {
	user := entity.User{}
	res := repo.db.Find(&user, id)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			// User not found
			return nil, nil
		}
		return nil, errors.Wrap(res.Error, "failed to find user")
	}

	return &user, nil
}

func (repo authGormRepository) findUserByEmail(email string) (*entity.User, error) {
	user := entity.User{}
	res := repo.db.Where("email = ?", email).Find(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			// User not found
			return nil, nil
		}
		return nil, errors.Wrap(res.Error, "failed to find user")
	}

	return &user, nil
}
