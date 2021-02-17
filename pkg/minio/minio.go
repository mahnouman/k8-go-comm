package minio

import (
	"github.com/minio/minio-go"
	"github.com/minio/minio-go/pkg/credentials"
)

//Mino settings and credentials
const (
	Endpoint        = "play.min.io"                              //to be replaced with real ones
	AccessKeyID     = "Q3AM3UQ867SPQQA43P2F"                     //to be replaced with real ones
	SecretAccessKey = "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG" //to be replaced with real ones
	UseSSL          = true
)

//Get mini client
func getMinioInstance() (*minio.Client, error) {

	Client, err := minio.New(Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(AccessKeyID, SecretAccessKey, ""),
		Secure: UseSSL,
	})

	return Client, err
}