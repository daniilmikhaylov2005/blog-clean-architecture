package utils

import (
  "golang.org/x/crypto/bcrypt"
)

func CreateHash(password string) (string, error) {
  hash, err := bcrypt.GenerateFromPassword([]byte(password), 7)
  return string(hash), err
}

func CheckHashAndPassword(password, hash string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
  return err == nil
}
