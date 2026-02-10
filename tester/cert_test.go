package tester

import (
	"crypto/x509"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCert(t *testing.T) {
	cert, err := GenX509KeyPair("test.local")
	require.NoError(t, err)

	pcert, err := x509.ParseCertificate(cert.Certificate[0])
	require.NoError(t, err)
	require.Equal(t, "test.local", pcert.Subject.CommonName)
}
