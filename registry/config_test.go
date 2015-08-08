package registry

import (
	"testing"
)

func TestValidateMirror(t *testing.T) {
	valid := []string{
		"http://mirror-1.com",
		"https://mirror-1.com",
		"http://localhost",
		"https://localhost",
		"http://localhost:5000",
		"https://localhost:5000",
		"http://127.0.0.1",
		"https://127.0.0.1",
		"http://127.0.0.1:5000",
		"https://127.0.0.1:5000",
	}

	invalid := []string{
		"!invalid!://%as%",
		"ftp://mirror-1.com",
		"http://mirror-1.com/",
		"http://mirror-1.com/?q=foo",
		"http://mirror-1.com/v1/",
		"http://mirror-1.com/v1/?q=foo",
		"http://mirror-1.com/v1/?q=foo#frag",
		"http://mirror-1.com?q=foo",
		"https://mirror-1.com#frag",
		"https://mirror-1.com/",
		"https://mirror-1.com/#frag",
		"https://mirror-1.com/v1/",
		"https://mirror-1.com/v1/#",
		"https://mirror-1.com?q",
	}

	for _, address := range valid {
		if ret, err := ValidateMirror(address); err != nil || ret == "" {
			t.Errorf("ValidateMirror(`"+address+"`) got %s %s", ret, err)
		}
	}

	for _, address := range invalid {
		if ret, err := ValidateMirror(address); err == nil || ret != "" {
			t.Errorf("ValidateMirror(`"+address+"`) got %s %s", ret, err)
		}
	}
}

func TestValidateFullyQualifiedCmd(t *testing.T) {
	valid := map[string]int{
		"pull":   1,
		"push":   1,
		"search": 1,
		"login":  1,
		"all":    len(ValidFullyQualifiedCmds),
	}

	invalid := []string{
		"foo", "bar", "", "pull push",
	}

	for cmd, num := range valid {
		if ret, err := ValidateFullyQualifiedCmd(cmd); err != nil || len(ret) != num {
			t.Errorf("ValidateFullyQualifiedCmd(`"+cmd+"`) got %v %s", ret, err)
		}
	}

	for _, cmd := range invalid {
		if ret, err := ValidateFullyQualifiedCmd(cmd); err == nil || len(ret) != 0 {
			t.Errorf("ValidateFullyQualifiedCmd(`"+cmd+"`) got %s %s", ret, err)
		}
	}
}
