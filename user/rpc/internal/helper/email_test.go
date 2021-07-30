package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendMail(t *testing.T) {
	err := ServerEmail.SendCode("1070259395@qq.com", "123456")
	assert.Nil(t, err)
}
