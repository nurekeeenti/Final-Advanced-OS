package utils

import (
    "log"
    "math/rand"
    "net/smtp"
    "strconv"
    "time"
)

var verificationCodes = make(map[string]string)

func SendVerificationCode(email string) string {
    rand.Seed(time.Now().UnixNano())
    code := strconv.Itoa(100000 + rand.Intn(900000))

    smtpHost := "smtp.gmail.com"
    smtpPort := "587"
    auth := smtp.PlainAuth("", "aldikoshka@gmail.com", "lmpo twun eihl xajw", smtpHost)

    msg := []byte("Subject: Подтверждение регистрации\r\n\r\nВаш код подтверждения: " + code)

    err := smtp.SendMail(smtpHost+":"+smtpPort, auth, "aldikoshka@gmail.com", []string{email}, msg)
    if err != nil {
        log.Println("Ошибка отправки письма:", err)
    }

    verificationCodes[email] = code
    return code
}

func VerifyCode(email, code string) bool {
    return verificationCodes[email] == code
}

func RemoveCode(email string) {
    delete(verificationCodes, email)
}
