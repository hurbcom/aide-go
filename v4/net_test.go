package v4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDSN2MAP(t *testing.T) {
	dsn := "MyUser:MyPassword@tcp(MyDatabaseHost.net:3306)/MyDatabaseName?param1=1&param2=2"
	actual := DSN2MAP(dsn)

	assert.Equal(t, "MyUser", actual["user"])
	assert.Equal(t, "MyPassword", actual["passwd"])
	assert.Equal(t, "tcp", actual["net"])
	assert.Equal(t, "MyDatabaseHost.net:3306", actual["addr"])
	assert.Equal(t, "MyDatabaseName", actual["dbname"])
	assert.Equal(t, "param1=1&param2=2", actual["params"])
}

func TestDSN2Publishable(t *testing.T) {
	dsn := "MyUser:MyPassword@tcp(MyDatabaseHost.net:3306)/MyDatabaseName?param1=1&param2=2"
	expected := "MyUser@tcp(MyDatabaseHost.net:3306)/MyDatabaseName?param1=1&param2=2"

	assert.Equal(t, expected, DSN2Publishable(dsn))
}
