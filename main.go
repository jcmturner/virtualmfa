package main

import (
	"crypto/sha1"
	"fmt"
	"os"
	"strings"
	"syscall"
	"time"

	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
	term "golang.org/x/crypto/ssh/terminal"
	"gopkg.in/jcmturner/gootp.v1"
)

func main() {
	mfa, err := secretPrompt()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading MFA secret: %v", err)
		os.Exit(1)
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	p := mpb.New(
		// override default (80) width
		mpb.WithWidth(60),
		// override default "[=>-]" format
		mpb.WithFormat("╢▌▌░╟"),
		// override default 120ms refresh rate
		mpb.WithRefreshRate(time.Millisecond*60),
	)

	bar, err := newOTP(mfa, p)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error generating OTP: %v", err)
		os.Exit(1)
	}

	for {
		time.Sleep(time.Second)
		bar.Increment()
		if bar.Completed() {
			bar, err = newOTP(mfa, p)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error generating OTP: %v", err)
				os.Exit(1)
			}
		}
	}
}

func newOTP(mfa string, p *mpb.Progress) (*mpb.Bar, error) {
	otp, c, err := gootp.TOTPNow(mfa, sha1.New, 6)
	if err != nil {
		return nil, err
	}
	bar := p.AddBar(int64(c),
		mpb.PrependDecorators(
			decor.StaticName(otp, len(otp)+1, decor.DidentRight),
		),
		mpb.BarRemoveOnComplete(),
	)
	return bar, nil
}

func secretPrompt() (string, error) {
	fmt.Print("Enter MFA Secret: ")
	b, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	s := string(b)
	strings.TrimSpace(s)
	fmt.Println("")
	return s, nil
}
