package models

import (
	"time"
	"src/github.com/astaxie/beego/orm"
	/*"src/github.com/golang/protobuf/conformance/internal/conformance_proto"*/
)


type User struct {
	Uid                            string `orm:"pk"`
	Username                       string
	Nickname		       string
	Email                          string `orm:"unique"`
	Phone                          string
	Location                       string
	Password                       string
	Registration_uid               string
	Registration_date              time.Time `orm:"auto_now_add;type(datetime)"`
	Password_reset_uid             string
	PhoneVerify                    string
	Reference_Link                 string
	Registers_count		    int32
	Click_count		    int32

}

type Profile struct {
	Uid                         string `orm:"pk"`
	Username                    string
	Reference_id                string
	Joined_date                 time.Time `orm:"auto_now_add;type(datetime)"`
 	Registers_count		    int32
	Click_count		    int32
	Password_reset_date         time.Time `orm:"auto_now_add;type(datetime)"`
	message                     string
	settings                    string
}

type Market struct {
	Uid                        string
	AssetName		   string
	NewPrice                   float32
	BoardNews                  string
	News		           string
	PromoteNews                string
}


func (account User) CopySignUpForm(form *FormUserSignUp) User {
  account.Username = form.Username
  account.Nickname = form.Nickname
  account.Email = form.Email
  account.Password = form.Password
  account.Phone = form.Phone

  return account
}


func (newsDetails Market) saveNewsForm(form *FormNews) Market{
  newsDetails.News = form.News
	newsDetails.BoardNews = form.BoardNews
	newsDetails.PromoteNews = form.PromoteNews
	newsDetails.AssetName = form.AssetName
	newsDetails.NewPrice = form.NewPrice

	return newsDetails
}


func (account User) CopyUpdateForm(form *FormUserUpdate) User {
  account.Username = form.Username
  account.Nickname = form.Nickname
  account.Email = form.Email
  account.Phone = form.Phone
  account.Location = form.Location

  return account
}

func (account User) VerifyPhone(form *FormPhoneVerification) User{
	account.Phone = form.Phone
	account.PhoneVerify = form.PhoneVerify

	return account
}

func init() {
	orm.RegisterModel(new(User), new(Profile), new(Market))
}
