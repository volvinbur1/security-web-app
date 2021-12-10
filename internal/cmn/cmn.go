package cmn

type User struct {
	Login string `bson:"login"`

	Name    string `bson:"name"`
	Surname string `bson:"surname"`
	Phone   string `bson:"phone"`

	EncryptionKey string `bson:"key"`

	Password string `bson:"pwdhash"`
	PwdSalt  string `bson:"pwdsalt"`
}
