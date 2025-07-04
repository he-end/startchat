package serviceotp

import (
	"errors"
	"fmt"
	"sc/internal/logger"
	serviceemail "sc/internal/service/email"

	"go.uber.org/zap"
)

// you can custom the body email here
var sampleBody = `<!DOCTYPE html>
<html lang="id">
<head>
  <meta charset="UTF-8">
  <title>Kode OTP</title>
  <style>
    body {
      background-color: #f4f4f4;
      font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif;
      margin: 0;
      padding: 0;
    }
    .container {
      max-width: 600px;
      margin: 30px auto;
      background-color: #ffffff;
      padding: 30px;
      border-radius: 8px;
      box-shadow: 0 0 10px rgba(0,0,0,0.05);
    }
    .header {
      text-align: center;
      color: #333333;
    }
    .otp-box {
      background-color: #eaf4ff;
      color: #2c3e50;
      padding: 20px;
      margin: 20px 0;
      font-size: 28px;
      font-weight: bold;
      text-align: center;
      letter-spacing: 6px;
      border-radius: 6px;
    }
    .footer {
      font-size: 13px;
      text-align: center;
      color: #999999;
      margin-top: 30px;
    }
  </style>
</head>
<body>
  <div class="container">
    <h2 class="header">Verifikasi Email Anda</h2>
    <p>Gunakan kode OTP berikut untuk melanjutkan proses verifikasi akun Anda:</p>

    <div class="otp-box">
      %v
    </div>

    <p>Kode ini hanya berlaku selama <strong>5 menit</strong>. Jangan berikan kode ini kepada siapa pun, termasuk pihak yang mengaku sebagai tim kami.</p>

    <p>Jika Anda tidak meminta kode ini, abaikan email ini.</p>

    <div class="footer">
      &copy; 2025 Perusahaan Anda. Semua Hak Dilindungi.
    </div>
  </div>
</body>
</html>
`

func SendOTPWithGmail(otpCode, email string) error {
	// if you want to change the subject, replace it
	subject := "OTP"
	body := fmt.Sprintf(sampleBody, otpCode)
	if ok := serviceemail.SendEmailWithGmail([]string{email}, subject, body); !ok {
		logger.Error("error sending otp code", zap.String("email", email))
		return errors.New("error sending otp code")
	}
	return nil
}
