package cmd

import (
    "errors"
    "fmt"
    ucbapi "github.com/saltatory/go-appstoreconnect-cli/api"
    "github.com/dgrijalva/jwt-go"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "os"
    "time"
)
jwt
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "os"
)

const ASC_PRIVATE_KEY_ID = "ASC_PRIVATE_KEY_ID"
const ASC_ISSUER_ID = "ASC_ISSUER_ID"
const ASC_AUDIENCE = "appstoreconnect-v1"

var JwtIssuerId string
var JwtExpirationTime int32
var JwtKeyId string

var Token string

const ERROR_SUBCOMMAND string = "Choose a sub-command"

func GenerateToken() string {
    d,_ := time.ParseDuration("60m")
    expirationTime := time.Now().Add(d)
    token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
        "iss": JwtIssuerId,
        "exp": expirationTime,
        "aud": ASC_AUDIENCE,
    })
    token.Header["kid"] = JwtKeyId

    tokenString, err := token.SignedString(JwtKeyId)

    if(err!=nil){
        fmt.Println(err)
        os.Exit(1)
    }

    return tokenString
}

var rootCmd = &cobra.Command{
    Use:   "asc-cli",
    Short: "asc-cli makes App Store Connect easy",
    Long:  "asc-cli is a CLI for App Store Connect",
    RunE: func(cmd *cobra.Command, args []string) error {
        return errors.New(ERROR_SUBCOMMAND)
    },
    PersistentPreRun: func(cmd *cobra.Command, args []string) {
        //Config.DefaultHeader["Authorization"] = "Bearer " + Token
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

}

func init() {
    viper.AutomaticEnv()
    rootCmd.PersistentFlags().StringVarP(&JwtKeyId, "key", "k", viper.GetString(ASC_PRIVATE_KEY_ID), "App Store Connect private key (required)")
    rootCmd.PersistentFlags().StringVarP(&JwtIssuerId, "iss", "i", viper.GetString(ASC_ISSUER_ID), "App Store Connect issuer id (required)")

    Token = GenerateToken()
}
