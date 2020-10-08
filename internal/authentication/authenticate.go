package authentication

import (
	"github.com/spf13/viper"
)

// type CloudSettings struct {
// 	AwsProfile string
// 	AwsRegion  string
// }

// var AwsSettings CloudSettings

// //AWS Authenticate
// func Amazon() {

// 	AwsSettings.AwsProfile = viper.GetString("aws.profile")
// 	AwsSettings.AwsRegion = viper.GetString("aws.region")

// }

//below is a bit over-kill. Commented out way for now is better
//but it could be that below is future-proofed for more robust authentication
//if we come to it.

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
