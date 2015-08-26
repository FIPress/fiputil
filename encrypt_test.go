package fiputil

import "testing"

func TestEncrypt(t *testing.T) {
	msg := "hi"
	cipher, err := Encrypt([]byte(msg), genKey())
	if err != nil {
		t.Error(err)
	}
	if len(cipher) != 20 {
		t.Error("encrypt failed")
	}
}

func TestPwd(t *testing.T) {
	p := []byte("abc")
	pwd, err := HashPwd(p)
	if err != nil {
		t.Error("pwd failed:", err)
	}

	ok := CheckPwd(p, pwd)
	if !ok {
		t.Error("Check pwd failed")
	}

	ok = CheckPwd([]byte("a"), pwd)
	if ok {
		t.Error("Check pwd error, should failed")
	}
}
