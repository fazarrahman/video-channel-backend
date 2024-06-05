package entity

type Users struct {
	Id           int64  `db:"id" json:"id"`
	Username     string `db:"username" json:"username"`
	PasswordHash []byte `db:"password_hash" json:"password_hash,omitempty"`
	Email        string `db:"email" json:"email"`
	Password     string `db:"-" json:"password,omitempty"`
}
