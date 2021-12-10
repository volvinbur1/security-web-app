package cmn

type User struct {
	Guid string `bson:"guid,omitempty"`

	Name    string `bson:"name,omitempty"`
	Surname string `bson:"surname,omitempty"`
	Phone   string `bson:"phone,omitempty"`
	Email   string `bson:"email,omitempty"`

	Login    string `bson:"login,omitempty"`
	Password string `bson:"pwdhash,omitempty"`
	PwdSalt  string `bson:"pwdsalt,omitempty"`
}
