package oauth2

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	AccountsTable = "oauth_users"
)

type Users interface {
	Create() (*OauthClients, error)
	Update()
	Get() (OauthUsers, error)
	Delete()
	Verify()
}

func NewCreateOauthUsers(username string, mail string, password string) *OauthUsers {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return &OauthUsers{Username: username, Mail: mail, Password: string(hash)}
}

func (a *OauthUsers) Create() (*OauthClients, error) {
	// TODO - Verify that the structure data, username email and password match the criteria.
	//if err := validator.New().Struct(a); err != nil {
	//	fmt.Println(err)
	//	return errors.New("FAILED_TO_VALIDATOR")
	//}

	if err := db.AutoMigrate(&OauthUsers{}); err != nil {
		return nil, err
	}

	if err := db.Debug().
		Table(AccountsTable).
		Where("username = ? ", a.Username).
		Or("mail = ?", a.Mail).
		First(&OauthUsers{}); err != nil {
		if err.Error != gorm.ErrRecordNotFound {
			return nil, errors.New("ACCOUNT ALREADY EXISTS")
		}
		if err := db.Table(AccountsTable).Create(&a).Error; err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (a *OauthUsers) Update() {
	//TODO implement me
	panic("implement me")
}

func NewOauthUsersId(id uint) *OauthUsers {
	return &OauthUsers{
		Model: gorm.Model{
			ID: id,
		},
	}
}

func (a *OauthUsers) Get() (*OauthUsers, error) {
	if err := db.Debug().
		Table(AccountsTable).
		Where("id = ?", a.ID).
		First(&a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

func (a *OauthUsers) Delete() {
	//TODO implement me
	panic("implement me")
}

func NewVerify(username string) *OauthUsers {
	return &OauthUsers{Username: username}
}

func (a *OauthUsers) Verify(password string) (*OauthUsers, error) {
	if err := db.Debug().
		Table(AccountsTable).
		Where("username = ?", a.Username).
		First(&a).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password)); err != nil {
		return nil, err
	}

	return &OauthUsers{
		Model: gorm.Model{
			ID: a.ID,
		},
		Username: a.Username,
		Mail:     a.Mail,
	}, nil
}
