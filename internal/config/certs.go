package config

import (
	"os"

	"github.com/pkg/errors"
)

const (
	keyPathEnv  = "TESTER_KEY_FILE"
	certPathEnv = "TESTER_CERT_FILE"
)

type Certificator interface {
	Key() string
	Cert() string
}

type certConfig struct {
	keyPath  string
	certPath string
}

func NewCertConfig() (Certificator, error) {
	keyPath := os.Getenv(keyPathEnv)
	if keyPath == "" {
		return nil, errors.Errorf("can't get key file at %s", keyPath)
	}

	if err := fileCheck(keyPath); err != nil {
		return nil, err
	}

	certPath := os.Getenv(certPathEnv)
	if certPath == "" {
		return nil, errors.Errorf("can't get cert file at %s", certPath)
	}

	if err := fileCheck(certPath); err != nil {
		return nil, err
	}

	certs := certConfig{
		keyPath:  keyPath,
		certPath: certPath,
	}

	return &certs, nil
}

func (c *certConfig) Key() string {
	return c.keyPath
}

func (c *certConfig) Cert() string {
	return c.certPath
}

func fileCheck(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return errors.Errorf("file %s not exist", path)
	}
	if err != nil {
		return errors.Wrapf(err, "can't open file %s", path)
	}

	return nil
}
