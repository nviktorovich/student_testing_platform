package entities_test

import (
	"github.com/nviktorovich/student_testing_platform/internal/entities"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewUser_Successful(t *testing.T) {
	id := "test id"
	name := "test name"
	email := "test email"

	u := entities.NewUser(id, name, email)
	require.NotNil(t, u)

	require.Equal(t, id, u.GetID())
	require.Equal(t, name, u.GetName())
	require.Equal(t, email, u.GetEmail())
}
