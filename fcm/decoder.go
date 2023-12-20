package fcm

import "encoding/base64"

func DecodeFirebaseAuthKey(encoded string) ([]byte, error) {
	decodedKey, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	return decodedKey, nil
}
