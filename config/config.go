package config

type Creds struct {
	Host string
	Port string
}

var Verification_listener_creds = Creds{
	Host: "0.0.0.0:",
	Port: "60321",
}

type BdCredentials struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

var BdCred = BdCredentials{
	Host:     "185.231.153.68",
	Port:     "5432",
	Database: "postgredb",
	Username: "admin",
	Password: "admin",
}
