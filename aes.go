// Encryption - AES-GCM 端到端加密
// Copyright (c) 2021-present, b3log.org
//
// Encryption is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
)

var GcmNonce = []byte("Jl8aXipNicI8")

func AESGCMEncryptBinBytes(data []byte, passwd string) (ret []byte, err error) {
	key := keyHash(passwd)
	block, err := aes.NewCipher([]byte(key))
	if nil != err {
		return
	}

	aesgcm, err := cipher.NewGCM(block)
	if nil != err {
		return
	}

	ret = aesgcm.Seal(nil, GcmNonce, data, nil)
	return
}

func AESGCMDecryptBinBytes(cryptData []byte, passwd string) (ret []byte, err error) {
	key := keyHash(passwd)
	block, err := aes.NewCipher([]byte(key))
	if nil != err {
		return
	}

	aesgcm, err := cipher.NewGCM(block)
	if nil != err {
		return
	}

	ret, err = aesgcm.Open(nil, GcmNonce, cryptData, nil)
	return
}

func keyHash(passwd string) string {
	hasher := md5.New()
	hasher.Write([]byte(passwd))
	return hex.EncodeToString(hasher.Sum(nil))
}
