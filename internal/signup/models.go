package signup

import (
	"context"

	"geeks-accelerator/oss/saas-starter-kit/internal/account"
	"geeks-accelerator/oss/saas-starter-kit/internal/user"
	"geeks-accelerator/oss/saas-starter-kit/internal/user_account"
	"github.com/jmoiron/sqlx"
)

// Repository defines the required dependencies for Signup.
type Repository struct {
	DbConn      *sqlx.DB
	User        *user.Repository
	UserAccount *user_account.Repository
	Account     *account.Repository
}

// NewRepository creates a new Repository that defines dependencies for Signup.
func NewRepository(db *sqlx.DB, user *user.Repository, userAccount *user_account.Repository, account *account.Repository) *Repository {
	return &Repository{
		DbConn:      db,
		User:        user,
		UserAccount: userAccount,
		Account:     account,
	}
}

// SignupRequest contains information needed perform signup.
type SignupRequest struct {
	Account SignupAccount `json:"account" validate:"required"` // Account details.
	User    SignupUser    `json:"user" validate:"required"`    // User details.
}

// SignupAccount defined the details needed for account.
type SignupAccount struct {
	Name     string  `json:"name" validate:"required,unique" example:"Company {RANDOM_UUID}"`
	Address1 string  `json:"address1" validate:"required" example:"221 Tatitlek Ave"`
	Address2 string  `json:"address2" validate:"omitempty" example:"Box #1832"`
	City     string  `json:"city" validate:"required" example:"Valdez"`
	Region   string  `json:"region" validate:"required" example:"AK"`
	Country  string  `json:"country" validate:"required" example:"USA"`
	Zipcode  string  `json:"zipcode" validate:"required" example:"99686"`
	Timezone *string `json:"timezone" validate:"omitempty" example:"America/Anchorage"`
}

// SignupUser defined the details needed for user.
type SignupUser struct {
	FirstName       string `json:"first_name" validate:"required" example:"Gabi"`
	LastName        string `json:"last_name" validate:"required" example:"May"`
	Email           string `json:"email" validate:"required,email,unique" example:"{RANDOM_EMAIL}"`
	Password        string `json:"password" validate:"required" example:"SecretString"`
	PasswordConfirm string `json:"password_confirm" validate:"required,eqfield=Password" example:"SecretString"`
}

// SignupResult response signup with created account and user.
type SignupResult struct {
	Account *account.Account `json:"account"`
	User    *user.User       `json:"user"`
}

// SignupResponse represents the user and account created for signup that is returned for display.
type SignupResponse struct {
	Account *account.AccountResponse `json:"account"`
	User    *user.UserResponse       `json:"user"`
}

// Response transforms SignupResult to SignupResponse that is used for display.
// Additional filtering by context values or translations could be applied.
func (m *SignupResult) Response(ctx context.Context) *SignupResponse {
	if m == nil {
		return nil
	}

	r := &SignupResponse{}
	if m.Account != nil {
		r.Account = m.Account.Response(ctx)
	}
	if m.User != nil {
		r.User = m.User.Response(ctx)
	}

	return r
}
