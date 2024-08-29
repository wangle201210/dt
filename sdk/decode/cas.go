package decode

func DoDecrypt(key string, data string) (res string, err error) {
	key += ":" + casKey
	encrypt, err := Decrypt(data, []byte(key))
	if err != nil {
		return "", err
	}
	res = string(encrypt)
	return
}
