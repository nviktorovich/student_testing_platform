package entities_test

import (
	"github.com/nviktorovich/student_testing_platform/internal/entities"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewQuestion_Failed(t *testing.T) {
	t.Parallel()

	q, err := entities.NewQuestion()
	require.Nil(t, q)
	require.ErrorIs(t, err, entities.ErrInvalidParam)
	require.Contains(t, err.Error(), "not correct question")

	q, err = entities.NewQuestion(entities.WithQuestion(""))
	require.Nil(t, q)
	require.ErrorIs(t, err, entities.ErrInvalidParam)
	require.Contains(t, err.Error(), "not correct question")

	q, err = entities.NewQuestion(entities.WithQuestion("correct question"))
	require.Nil(t, q)
	require.ErrorIs(t, err, entities.ErrInvalidParam)
	require.Contains(t, err.Error(), "not correct id")

	q, err = entities.NewQuestion(entities.WithQuestion("correct question"),
		entities.WithQuestionID(""))
	require.Nil(t, q)
	require.ErrorIs(t, err, entities.ErrInvalidParam)
	require.Contains(t, err.Error(), "not correct id")

	q, err = entities.NewQuestion(entities.WithQuestion("correct question"),
		entities.WithQuestionID("correctID"))
	require.Nil(t, q)
	require.ErrorIs(t, err, entities.ErrInvalidParam)
	require.Contains(t, err.Error(), "not correct question type")

	q, err = entities.NewQuestion(entities.WithQuestion("correct question"),
		entities.WithQuestionID("correctID"),
		entities.WithQuestionType(entities.QuestionType(4)))
	require.Nil(t, q)
	require.ErrorIs(t, err, entities.ErrInvalidParam)
	require.Contains(t, err.Error(), "not correct question type")

	q, err = entities.NewQuestion(entities.WithQuestion("correct question"),
		entities.WithQuestionID("correctID"),
		entities.WithQuestionType(entities.MultiQuestion))
	require.Nil(t, q)
	require.ErrorIs(t, err, entities.ErrInvalidParam)
	require.Contains(t, err.Error(), "not correct answer variants")

	var nilMap map[string]bool
	q, err = entities.NewQuestion(entities.WithQuestion("correct question"),
		entities.WithQuestionID("correctID"),
		entities.WithQuestionType(entities.MultiQuestion),
		entities.WithQuestionVariants(nilMap))
	require.Nil(t, q)
	require.ErrorIs(t, err, entities.ErrInvalidParam)
	require.Contains(t, err.Error(), "not correct answer variants")

	q, err = entities.NewQuestion(entities.WithQuestion("correct question"),
		entities.WithQuestionID("correctID"),
		entities.WithQuestionType(entities.MultiQuestion),
		entities.WithQuestionVariants(map[string]bool{"1": true, "2": true, "3": true, "4": true, "5": false}))
	require.Nil(t, q)
	require.ErrorIs(t, err, entities.ErrInvalidParam)
	require.Contains(t, err.Error(), "not correct answer variants")
}

func TestNewQuestion_Successful(t *testing.T) {
	t.Parallel()

	q, err := entities.NewQuestion(
		entities.WithQuestion("correct question"),
		entities.WithQuestionID("correctID"),
		entities.WithQuestionType(entities.MultiQuestion),
		entities.WithQuestionVariants(map[string]bool{"1": true, "2": true, "3": true, "4": true}),
	)

	require.NoError(t, err)
	require.NotNil(t, q)
}

func TestQuestion_SetUserAnswer(t *testing.T) {
	t.Parallel()

	q, err := entities.NewQuestion(
		entities.WithQuestion("correct question"),
		entities.WithQuestionID("correctID"),
		entities.WithQuestionType(entities.MultiQuestion),
		entities.WithQuestionVariants(map[string]bool{"1": true, "2": true, "3": true, "4": true}),
	)

	require.NoError(t, err)
	require.NotNil(t, q)

	err = q.SetUserAnswer(nil)
	require.ErrorIs(t, err, entities.ErrInvalidParam)
	require.Contains(t, err.Error(), "incorrect answers")

	err = q.SetUserAnswer([]string{})
	require.ErrorIs(t, err, entities.ErrInvalidParam)
	require.Contains(t, err.Error(), "incorrect answers")

	err = q.SetUserAnswer([]string{"1", "2"})
	require.NoError(t, err)
}
