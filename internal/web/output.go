package web

import (
	"errors"
	"github.com/volvinbur1/security-web-app/internal/cmn"
	"github.com/volvinbur1/security-web-app/internal/db"
	"github.com/volvinbur1/security-web-app/internal/web/aesgcm"
)

func getInfoAboutUser(currentUserUid string, dbMgr *db.Manager) (cmn.User, error) {
	users, err := dbMgr.GetUsers()
	if err != nil {
		return cmn.User{}, err
	}

	for _, user := range users {
		if user.Guid != currentUserUid {
			continue
		}

		err = decryptUserData(&user)
		if err != nil {
			return cmn.User{}, err
		}
		return user, nil
	}

	return cmn.User{}, errors.New("user with such id not found")
}

func decryptUserData(user *cmn.User) error {
	key, err := getUserKey(user.Guid)
	if err != nil {
		return err
	}

	user.Surname, err = aesgcm.DecryptData(user.Surname, key)
	if err != nil {
		return err
	}
	user.Phone, err = aesgcm.DecryptData(user.Phone, key)
	if err != nil {
		return err
	}
	user.Email, err = aesgcm.DecryptData(user.Email, key)
	if err != nil {
		return err
	}

	return nil
}
