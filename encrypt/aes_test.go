package encrypt

import "testing"

func TestPKCS7Padding(t *testing.T) {
	tests := []struct {
		chiperText string
		res        string
	}{
		{"hello world", "hello worlduuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuu"},
	}

	for _, item := range tests {
		r := PKCS7Padding([]byte(item.chiperText), 128)
		if string(r) != item.res {
			t.Errorf("expected %s, but has %s", item.res, string(r))
		}
	}
}

func TestPKCS7UnPadding(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello worlduuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuu", "hello world"},
	}

	for _, item := range tests {
		actual, err := PKCS7UnPadding([]byte(item.input))
		if err != nil {
			t.Fatal(err)
		}
		if string(actual) != item.expected {
			t.Errorf("expected %s, but has %s", item.expected, string(actual))
		}
	}
}

func TestEncrypt(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "upCuxEuUIQvBPtKF01AvenjNxy2MNaEz46/0IQ2HfPQ="},
	}

	for _, item := range tests {
		encdata, err := AesEcrypt([]byte(item.input), PwdKey)
		if err != nil {
			t.Fatal(err)
		}
		actual, err := EnPwdCode(encdata)
		if err != nil {
			t.Fatal(err)
		}
		if actual != item.expected {
			t.Errorf("expected %s, but is %s", item.expected, actual)
		}
	}
}

func TestDecrypt(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"upCuxEuUIQvBPtKF01AvenjNxy2MNaEz46/0IQ2HfPQ=", "hello"},
	}

	for _, item := range tests {
		decdata, err := DePwdCode(item.input)
		if err != nil {
			t.Fatal(err)
		}
		actual, err := AesDeCrypt(decdata, PwdKey)
		if err != nil {
			t.Fatal(err)
		}
		if string(actual) != item.expected {
			t.Errorf("expected %s, actual %s", item.expected, string(actual))
		}
	}
}
