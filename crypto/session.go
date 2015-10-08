package crypto

import (
  "crypto/rand"
  "golang.org/x/crypto/salsa20"
)

var Session CryptoSession

func init() {
  Session = NewSession()
}

type CryptoSession struct {
  sessionKey *[32]byte
  nonce []byte
}

func NewSession() CryptoSession {
  nonce := make([]byte, 8)
  if _, err := rand.Read(nonce); err != nil {
    panic(err)
  }

  var key [32]byte
  if _, err := rand.Read(key[:]); err != nil {
    panic(err)
  }

  return CryptoSession{
    sessionKey: &key,
    nonce: nonce,
  }
}

func (session *CryptoSession) Encrypt(plaintext string) []byte {
  in := make([]byte, 64)
  copy(in, []byte(plaintext))
  out := make([]byte, 64)
  salsa20.XORKeyStream(out, in, Session.nonce, Session.sessionKey)
  return out
}

func (session *CryptoSession) Decrypt(ciphertext []byte) string {
  out := make([]byte, len(ciphertext))
  salsa20.XORKeyStream(out, ciphertext, Session.nonce, Session.sessionKey)
  return string(out)
}
