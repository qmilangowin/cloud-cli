package authentication

import (
	"github.com/spf13/viper"
)

//Authenticator
type Authenticator interface {
	initialize()
}

type AwsSettings struct {
	AwsProfile string
	AwsRegion  string
}

var AwsAuth AwsSettings

func (a *AwsSettings) initialize() {

	AwsAuth.AwsProfile = viper.GetString("aws.profile")
	AwsAuth.AwsRegion = viper.GetString("aws.region")

}

func authenticate(a Authenticator) {
	a.initialize()
}

func Amazon() {
	r := AwsSettings{}
	authenticate(&r)
}
