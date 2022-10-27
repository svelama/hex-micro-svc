package shortener

import (
	"errors"
	"time"

	"github.com/google/uuid"
	errs "github.com/pkg/errors"
	"gopkg.in/dealancer/validate.v2"
)

var (
	ErrRirectNotFound = errors.New("Redirect Not Found")
	ErrRirectInvalidd = errors.New("Redirect In Valid")
)

type redirectService struct {
	repo RedirectRepository
}

func NewRedirectService(repo RedirectRepository) *redirectService {
	return &redirectService{repo}
}

func (rs *redirectService) Find(code string) (*Redirect, error) {
	return rs.repo.Find(code)
}

func (rs *redirectService) Store(redirect *Redirect) error {

	// Perform validation
	if err := validate.Validate(redirect); err != nil {
		return errs.Wrap(ErrRirectInvalidd, "service.redirect.store")
	}

	redirect.Code = uuid.New().String()
	redirect.CreatedAt = time.Now().UTC().Unix()
	return rs.repo.Store(redirect)
}
