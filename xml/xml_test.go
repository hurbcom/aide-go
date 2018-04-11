package xml

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateXMLWithXSD(t *testing.T) {
	err := ValidateXMLWithXSD("fixtures/transaction.xml", "fixtures/transaction.xsd")
	assert.Nil(t, err)
}
