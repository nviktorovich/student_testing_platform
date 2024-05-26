package entities

import (
	"github.com/pkg/errors"
)

const (
	SingleQuestion QuestionType = iota + 1
	MultiQuestion
)

type QuestionType int

// Question is a base entity of testing platform
type Question struct {
	questionType QuestionType
	id           string
	question     string
	variants     map[string]bool
	userAnswers  []string
}

func NewQuestion(opts ...QuestionOption) (*Question, error) {
	settings := QuestionOptions{}
	settings.QuestionOptionsApply(opts...)

	if !isQuestionValid(settings.question) {
		return nil, errors.Wrapf(ErrInvalidParam, "not correct question: %s", settings.question)
	}

	if !isIDValid(settings.id) {
		return nil, errors.Wrapf(ErrInvalidParam, "not correct id: %s", settings.id)
	}

	if !isQuestionTypeValid(settings.questionType) {
		return nil, errors.Wrapf(ErrInvalidParam, "not correct question type: %v", settings.questionType)
	}

	if !isAnswerVariantsValid(settings.variants) {
		return nil, errors.Wrapf(ErrInvalidParam, "not correct answer variants: %v", settings.variants)
	}

	q := &Question{
		questionType: settings.questionType,
		id:           settings.id,
		question:     settings.question,
		variants:     settings.variants,
	}
	return q, nil
}

type QuestionOptions struct {
	questionType QuestionType
	id           string
	question     string
	variants     map[string]bool
}

type QuestionOption func(*QuestionOptions)

func WithQuestionType(questionType QuestionType) QuestionOption {
	return func(o *QuestionOptions) {
		o.questionType = questionType
	}
}

func WithQuestionID(id string) QuestionOption {
	return func(o *QuestionOptions) {
		o.id = id
	}
}

func WithQuestion(q string) QuestionOption {
	return func(o *QuestionOptions) {
		o.question = q
	}
}

func WithQuestionVariants(v map[string]bool) QuestionOption {
	return func(o *QuestionOptions) {
		o.variants = v
	}
}

func (o *QuestionOptions) QuestionOptionsApply(opts ...QuestionOption) {
	for _, opt := range opts {
		opt(o)
	}
}

func isQuestionValid(q string) bool {
	return !(q == "")
}

func isIDValid(id string) bool {
	return !(id == "")
}

func isAnswerVariantsValid(ans map[string]bool) bool {
	return !(ans == nil || len(ans) != 4)
}

func isQuestionTypeValid(t QuestionType) bool {
	return t == SingleQuestion ||
		t == MultiQuestion
}

//nolint:gosimple //ok
func (q *Question) SetUserAnswer(answers []string) error {
	if answers == nil || len(answers) == 0 {
		return errors.Wrapf(ErrInvalidParam, "incorrect answers: %v", answers)
	}
	q.userAnswers = answers
	return nil
}
