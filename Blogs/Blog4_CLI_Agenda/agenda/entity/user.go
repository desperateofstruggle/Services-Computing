package entity

// User Struct
type User struct {
	Username string
	Password string
	Phone    string
	Email    string
}

// Init ...
func (mUser *User) Init(tName, tPassword, tEmail, tPhone string) {
	mUser.Username = tName
	mUser.Password = tPassword
	mUser.Email = tEmail
	mUser.Phone = tPhone
}

// GainUser ...
func (mUser *User) GainUser(u User) {
	mUser.Username = u.Username
	mUser.Password = u.Password
	mUser.Email = u.Email
	mUser.Phone = u.Phone
}

// GetUsername ...
func (mUser User) GetUsername() string {
	return mUser.Username
}

// SetGetUsername ...
func (mUser User) SetGetUsername(tname string) {
	mUser.Username = tname
}

// GetPassword ...
func (mUser User) GetPassword() string {
	return mUser.Password
}

// SetPassword ...
func (mUser User) SetPassword(tpassword string) {
	mUser.Password = tpassword
}

// GetEmail ...
func (mUser User) GetEmail() string {
	return mUser.Email
}

// SetEmail ...
func (mUser User) SetEmail(temail string) {
	mUser.Email = temail
}

// GetPhone ...
func (mUser User) GetPhone() string {
	return mUser.Phone
}

// SetPhone ...
func (mUser User) SetPhone(tphone string) {
	mUser.Phone = tphone
}

// IsUserExisted ...
func IsUserExisted(usern string, users []User) bool {
	for _, i := range users {
		if usern == i.GetUsername() {
			return true
		}
	}
	return false
}
