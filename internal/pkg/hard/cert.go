// Path: internal/pkg/hard
// FileName: cert.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/12/29$ 17:04$

package hard

import (
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/pem"
	log "github.com/sirupsen/logrus"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
	"math/big"
)

func GetInfoFromCert(certBase64 string, itype uint) (info []byte, err error) {
	certBytes, err := base64.StdEncoding.DecodeString(certBase64)
	if err != nil {
		log.Println("base64 decode error:", err)
		return
	}
	block := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	}
	certObj, err := parseCertificate(pem.EncodeToMemory(block))
	if err != nil {
		log.Println("parse cert error:", err)
		return
	}
	switch itype {
	case 1:
		info = getPubkey(certObj)
	case 2:
		info = getDName(certObj, 2)
	case 3:
		info = getDName(certObj, 3)
	case 4:
		info = getDName(certObj, 4)
	default:
		return
	}
	return
}
func parseCertificate(cert []byte) (*x509.Certificate, error) {
	return x509.ReadCertificateFromPem(cert)
}

func getPubkey(cert *x509.Certificate) (pubkey []byte) {
	sm2pub := &sm2.PublicKey{}
	switch pub := cert.PublicKey.(type) {
	case *ecdsa.PublicKey:
		sm2pub = &sm2.PublicKey{
			Curve: pub.Curve,
			X:     pub.X,
			Y:     pub.Y,
		}
	}
	pubKeySubject := make([]byte, 65)
	pubKeySubject[0] = 0x04
	// copy(pubKeySubject[1:33], sm2pub.X.Bytes())
	// copy(pubKeySubject[33:], sm2pub.Y.Bytes())
	copy(pubKeySubject[33-len(sm2pub.X.Bytes()):33], sm2pub.X.Bytes())
	copy(pubKeySubject[65-len(sm2pub.Y.Bytes()):], sm2pub.Y.Bytes())
	return pubKeySubject
}

func getDName(cert *x509.Certificate, ctype uint) (cn []byte) {
	switch ctype {
	case 2:
		return []byte(cert.Subject.CommonName)
	case 3:
		return []byte(filterNil(cert.Subject.OrganizationalUnit))
	case 4:
		return []byte(StringSn(cert.SerialNumber))
	default:
		return
	}
}
func filterNil(value []string) (result string) {
	if value == nil || len(value) == 0 {
		return ""
	} else {
		return value[0]
	}
}

func StringSn(sn *big.Int) string {
	return sn.Text(16)
}
