package cmn

type User struct {
	Login   string `bson:"login"`
	Name    string `bson:"name"`
	Surname string `bson:"surname"`

	Password string `bson:"pwdhash"`
	PwdSalt  string `bson:"pwdsalt"`
}