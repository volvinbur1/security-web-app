package cmn

type User struct {
	Guid string `bson:"guid"`

	Name    string `bson:"name"`
	Surname string `bson:"surname"`
	Phone   string `bson:"phone"`

	Login    string `bson:"login"`
	Password string `bson:"pwdhash"`
	PwdSalt  string `bson:"pwdsalt"`
}
