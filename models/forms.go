package models

import (
  "regexp"
  "src/github.com/astaxie/beego/validation"
)


type FormUserSignUp struct {
  Username           string `valid:"Required;MinSize(2);Alpha"`
  Nickname           string `valid:"Required;MinSize(2);Alpha"`
  Phone              string `valid:"Required;MinSize(12);MaxSize(12)"`
  Email           string `valid:"Email"`
  Password        string `valid:"MinSize(8);MaxSize(36)"`
  ConfirmPassword string `valid:"Required"`
}

type FormUserUpdate struct {
  Username           string `valid:"Required;MinSize(2);Alpha"`
  Nickname            string `valid:"Required;MinSize(2);Alpha"`
  Email           string `valid:"Email"`
  Phone          string
  Location         string
  CurrentPassword string `valid:"Required"`
}

type FormPhoneVerification struct {
  Phone   string
  PhoneVerify string
}

type FormUserPasswordUpdate struct {
  NewPassword              string `valid:"MinSize(8);MaxSize(36)"`
  NewPasswordConfirmation  string `valid:"Required"`
  CurrentPassword          string `valid:"Required"`
}

type FormNews struct {
  News               string `valid:"MinSize(10);MaxSize(256)"`
  BoardNews          string `valid:"Required"`
  PromoteNews        string `valid:"Required"`
  AssetName		   string `valid:"Required"`
  NewPrice                 float32 `valid:"Required"`
}

func (form *FormUserPasswordUpdate) Valid(v *validation.Validation) {
  // Check passwords match
  if form.NewPassword != form.NewPasswordConfirmation {
    v.SetError("NewPassword", "Passwords don't match")
    v.SetError("NewPasswordConfirmation", "Passwords don't match")
  }

  // Check password contain capital letter
  r, _ := regexp.Compile(`[A-Z]`)

  if !r.MatchString(form.NewPassword) {
    v.SetError("NewPassword", "Password must contain at least one capital letter")
  }

  // Check password contain lowercase letter
  r, _ = regexp.Compile(`[a-z]`)

  if !r.MatchString(form.NewPassword) {
    v.SetError("NewPassword", "Password must contain at least one lower case letter")
  }

  // Check password contain number
  r, _ = regexp.Compile(`[0-9]`)

  if !r.MatchString(form.NewPassword) {
    v.SetError("NewPassword", "Password must contain at least one number")
  }

}//end FormPasswordUpdate.valid func


func (form *FormUserSignUp) Valid(v *validation.Validation) {

  // Check passwords match
  if form.Password != form.ConfirmPassword {
    v.SetError("Password", "Passwords don't match")
    v.SetError("ConfirmPassword", "Passwords don't match")
  }

  // Check password contain capital letter
  r, _ := regexp.Compile(`[A-Z]`)

  if !r.MatchString(form.Password) {
    v.SetError("Password", "Password must contain at least one capital letter")
  }

  // Check password contain lowercase letter
  r, _ = regexp.Compile(`[a-z]`)

  if !r.MatchString(form.Password) {
    v.SetError("Password", "Password must contain at least one lower case letter")
  }

  // Check password contain number
  r, _ = regexp.Compile(`[0-9]`)

  if !r.MatchString(form.Password) {
    v.SetError("Password", "Password must contain at least one number")
  }

}//end FormAccountSignUp validation



func (form *FormUserUpdate) Valid(v *validation.Validation) {

  valid := validation.Validation{}

  if form.Phone != "" {
    matched, err := regexp.MatchString("^\\d{3}\\-\\d{3}\\-\\d{4}$", form.Phone1)

    if err != nil {
      v.SetError("Phone1", "Something wierd is going on here.")
    } else if !matched {
      v.SetError("Phone1", "Phone should be something like: 647-123-4567")
    }
  } // end Phone validation if block


  for _, err := range valid.Errors {
    v.SetError(err.Key, err.Message)
  }

}//end FormAccountUpdate validation
