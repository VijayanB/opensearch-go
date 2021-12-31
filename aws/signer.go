package aws

import (
	"bytes"

	"io/ioutil"
	"net/http"
	"time"

	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
)

func sign(req *http.Request, region string, serviceName string, signer *v4.Signer) (err error) {
	if req.Body == nil {
		_, err = signer.Sign(req, nil, serviceName, region, time.Now().UTC())
		return
	}
	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}
	_, err = signer.Sign(req, bytes.NewReader(buf), serviceName, region, time.Now().UTC())
	return err
}

//SignRequest signs the request using SigV4
func SignRequest(req *http.Request, config *Config) error {
	signer := v4.NewSigner(config.Credentials)
	return sign(req, config.Region, config.Service, signer)
}

