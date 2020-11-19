package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
}

//r nya kecil artinya tidak bisa di akses bagian repo/kode yang lain, jika besar maka bisa diluar
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}
