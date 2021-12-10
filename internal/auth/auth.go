package auth

import (
	"crypto/rand"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/volvinbur1/security-web-app/internal/cmn"
	"github.com/volvinbur1/security-web-app/internal/db"
	"golang.org/x/crypto/bcrypt"
	"io"
	"os"
)

const saltSize = 16

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

	err = preValidatePassword(newUser.Password)
	if err != nil {
		return err
	}

	newUser.PwdSalt = string(genSalt([]byte(newUser.Password)))
	newUser.Password = newUser.PwdSalt + newUser.Password
	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		return err
	}

	newUser.Password = string(hash)
	return dbMgr.AddUser(newUser)
}

func preValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password too short")
	}

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
