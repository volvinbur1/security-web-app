package web

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/volvinbur1/security-web-app/internal/cmn"
	"github.com/volvinbur1/security-web-app/internal/db"
	"github.com/volvinbur1/security-web-app/internal/web/aesgcm"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"os"
)

const saltSize = 16
const keysFile = "./bin/aes-keys/keys.txt"

func loginUser(dbMgr *db.Manager, loggingUser cmn.User) error {
	users, err := dbMgr.GetUsers()
	if err != nil {
		return err
	}

	for _, u := range users {
		if u.Login == loggingUser.Login {
			pwdBytes, err := hex.DecodeString(u.Password)
			if err != nil {
				return err
			}

			if bcrypt.CompareHashAndPassword(pwdBytes, []byte(u.PwdSalt+loggingUser.Login)) != nil {
				return errors.New("password incorrect")
			}

			log.Println(loggingUser.Login, "logged in.")
			return nil
		}
	}

	return errors.New("user not registered")
}

func Register(dbMgr *db.Manager, newUser cmn.User) error {
	users, err := dbMgr.GetUsers()
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Login == newUser.Login {
			return errors.New("user already registered")
		}
	}

	newUser.Guid = uuid.New().String()
	err = hashPassword(&newUser)
	if err != nil {
		return err
	}

	err = encryptUserData(&newUser)
	if err != nil {
		return err
	}

	err = dbMgr.AddUser(newUser)
	if err != nil {
		return err
	}

	log.Println(newUser.Login, "registered in.")
	return nil
}

func genSalt(password []byte) []byte {
	buf := make([]byte, saltSize, saltSize+sha1.Size)
	_, err := io.ReadFull(rand.Reader, buf)

	if err != nil {
		fmt.Printf("random read failed: %v", err)
		os.Exit(1)
	}

	hash := sha1.New()
	hash.Write(buf)
	hash.Write(password)
	return hash.Sum(buf)
}

func hashPassword(user *cmn.User) error {
	user.PwdSalt = hex.EncodeToString(genSalt([]byte(user.Password)))
	user.Password = user.PwdSalt + user.Password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}

	user.Password = hex.EncodeToString(hash)
	return nil
}

func encryptUserData(user *cmn.User) error {
	key, err := aesgcm.GenerateUniqueKey()
	if err != nil {
		return err
	}

	err = storeUserKey(user.Guid, key)
	if err != nil {
		return err
	}

	user.Surname, err = aesgcm.EncryptUserData(user.Surname, key)
	if err != nil {
		return err
	}

	user.Phone, err = aesgcm.EncryptUserData(user.Phone, key)
	if err != nil {
		return err
	}

	user.Email, err = aesgcm.EncryptUserData(user.Email, key)
	if err != nil {
		return err
	}

	return nil
}

func storeUserKey(guid, key string) error {
	err := os.MkdirAll(keysFile, 0600)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(keysFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.WriteString(guid + ";" + key); err != nil {
		return err
	}
	return nil
}
