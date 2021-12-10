package auth

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/volvinbur1/security-web-app/internal/cmn"
	"github.com/volvinbur1/security-web-app/internal/db"
	"golang.org/x/crypto/bcrypt"
	"io"
	"os"
	"unicode"
)

const saltSize = 16

func LoginUser(dbMgr *db.Manager, loggingUser cmn.User) error {
	users, err := dbMgr.GetUsers()
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Login == loggingUser.Login {
			pwdBytes, err := hex.DecodeString(user.Password)
			if err != nil {
				return err
			}

			if bcrypt.CompareHashAndPassword(pwdBytes, []byte(user.PwdSalt+loggingUser.Login)) != nil {
				return errors.New("password incorrect")
			}

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

	err = preValidatePassword(newUser.Password)
	if err != nil {
		return err
	}

	newUser.PwdSalt = hex.EncodeToString(genSalt([]byte(newUser.Password)))
	newUser.Password = newUser.PwdSalt + newUser.Password
	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		return err
	}

	newUser.Password = hex.EncodeToString(hash)
	return dbMgr.AddUser(newUser)
}

func preValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password too short")
	}

	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for _, ch := range password {
		if unicode.IsUpper(ch) {
			hasUpper = true
			continue
		}

		if unicode.IsLower(ch) {
			hasLower = true
			continue
		}
		if unicode.IsNumber(ch) {
			hasNumber = true
			continue
		}

		if unicode.IsPunct(ch) || unicode.IsSymbol(ch) {
			hasSpecial = true
		}
	}

	if hasNumber && hasSpecial && hasLower && hasUpper {
		return nil
	}

	return errors.New("password should contain numbers, upper and lower characters, special characters")
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
