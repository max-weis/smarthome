package internal

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ziflex/lecho"
)

// Custom errors
var (
	ErrBadRequest          = errors.New("bad request")
	ErrNotFound            = errors.New("not found")
	ErrInternalServerError = errors.New("internal server error")
)

// CustomError wraps a standard error with a status code
type CustomError struct {
	Err        error
	StatusCode int
}

func (ce *CustomError) Error() string {
	return ce.Err.Error()
}

// NewCustomError creates a new CustomError
func NewCustomError(err error, statusCode int) *CustomError {
	return &CustomError{
		Err:        err,
		StatusCode: statusCode,
	}
}

// ProblemDetails represents the RFC 7807 problem details
type ProblemDetails struct {
	Type     string      `json:"type"`
	Title    string      `json:"title"`
	Status   int         `json:"status"`
	Detail   string      `json:"detail,omitempty"`
	Instance string      `json:"instance,omitempty"`
	Extra    interface{} `json:"extra,omitempty"`
}

// errorHandler is a custom error handling middleware supporting RFC 7807
func errorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			problem := ProblemDetails{
				Type:     "about:blank",
				Title:    "Internal Server Error",
				Status:   http.StatusInternalServerError,
				Instance: c.Request().RequestURI,
			}

			var ce *CustomError
			if errors.As(err, &ce) {
				problem.Status = ce.StatusCode
				problem.Title = http.StatusText(ce.StatusCode)
				problem.Detail = ce.Err.Error()
			} else if he, ok := err.(*echo.HTTPError); ok {
				problem.Status = he.Code
				problem.Title = http.StatusText(he.Code)
				problem.Detail = fmt.Sprintf("%v", he.Message)
			} else {
				problem.Detail = err.Error()
			}

			switch {
			case errors.Is(err, ErrBadRequest):
				problem.Type = "https://example.com/errors/bad-request"
			case errors.Is(err, ErrNotFound):
				problem.Type = "https://example.com/errors/not-found"
			case errors.Is(err, ErrInternalServerError):
				problem.Type = "https://example.com/errors/internal-server-error"
			}

			if !c.Response().Committed {
				if c.Request().Method == http.MethodHead {
					err = c.NoContent(problem.Status)
				} else {
					c.Response().Header().Set(echo.HeaderContentType, "application/problem+json")
					err = c.JSON(problem.Status, problem)
				}
			}

			c.Logger().Error(err)
		}

		return nil
	}
}

func NewEchoServer() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(errorHandler)

	e.Logger = lecho.New(os.Stdout)

	return e
}
