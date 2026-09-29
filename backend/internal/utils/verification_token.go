package utils

import (
    "crypto/rand"
    "crypto/sha256"
    "encoding/hex"
)

func GenerateVerificationToken() (
    token string,
    tokenHash string,
    err error,
) {
    bytes := make([]byte, 32)

    if _, err = rand.Read(bytes); err != nil {
        return "", "", err
    }

    // token ที่ส่งผ่าน URL ในอีเมล
    token = hex.EncodeToString(bytes)

    // hash ที่เก็บใน database
    hash := sha256.Sum256([]byte(token))
    tokenHash = hex.EncodeToString(hash[:])

    return token, tokenHash, nil
}

func HashVerificationToken(
    token string,
) string {
    hash := sha256.Sum256([]byte(token))

    return hex.EncodeToString(hash[:])
}