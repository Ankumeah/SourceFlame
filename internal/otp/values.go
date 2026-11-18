package otp

import (
	"time"
)

const otpLifetime = time.Minute * 10
const cleanTick = time.Minute * 5

var otpEmailContent = `
Hi %s!

Here is your OTP for SourceFlame
%d

This will expire in %s

If you were not expecting this email then feel free to ignore it!

Good day! :]
`
