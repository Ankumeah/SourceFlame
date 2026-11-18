package otp

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"
)

type OTP struct {
	otp      uint
	issuedAt time.Time
}

type OTPManager struct {
	otps map[string]OTP
	mu   sync.Mutex
}

func NewOTPManager() *OTPManager {
	otpManager := &OTPManager{
		otps: map[string]OTP{},
		mu:   sync.Mutex{},
	}

	go otpManager.startCleaner()

	return otpManager
}

func (o *OTPManager) NewOTP(email string) (OTP, error) {
	o.mu.Lock()
	defer o.mu.Unlock()

	num, err := rand.Int(rand.Reader, big.NewInt(9000))
	if err != nil {
		return OTP{}, err
	}

	o.otps[email] = OTP{
		otp:      uint(num.Uint64() + 1000),
		issuedAt: time.Now(),
	}

	return o.otps[email], nil
}

func (o *OTPManager) CheckOTP(email string, otp uint) bool {
	o.mu.Lock()
	defer o.mu.Unlock()

	savedOTP, ok := o.otps[email]

	if !ok ||
		time.Now().After(savedOTP.issuedAt.Add(otpLifetime)) ||
		otp != savedOTP.otp {

		return false
	}

	delete(o.otps, email)

	return true
}

func (o *OTPManager) startCleaner() {
	tick := time.Tick(cleanTick)

	for {
		select {
		case <-tick:
			func() {
				o.mu.Lock()
				defer o.mu.Unlock()

				for email, otp := range o.otps {
					if time.Now().After(otp.issuedAt.Add(otpLifetime)) {
						delete(o.otps, email)
					}
				}
			}()
		}
	}
}

func (o *OTP) GetEmailContent(user string) []byte {
	return []byte(fmt.Sprintf(otpEmailContent, user, o.otp, otpLifetime))
}
