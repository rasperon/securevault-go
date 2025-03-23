package auth

import (
    "crypto/rand"
    "fmt"
    "strconv"
)

// GenerateOTP rastgele bir 6 haneli doğrulama kodu (OTP) üretir
func GenerateOTP() string {
    otp := make([]byte, 6)
    if _, err := rand.Read(otp); err != nil {
        fmt.Println("Error generating OTP:", err)
        return ""
    }

    // OTP'yi 6 haneli sayıya dönüştürme
    var otpStr string
    for i := 0; i < 6; i++ {
        otpStr += strconv.Itoa(int(otp[i]) % 10)
    }

    return otpStr
}

// VerifyOTP kullanıcı tarafından girilen OTP'yi doğrular
func VerifyOTP(inputOTP, actualOTP string) bool {
    return inputOTP == actualOTP
}
