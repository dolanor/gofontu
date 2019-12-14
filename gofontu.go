// SPDX-License-Identifier: Unlicense OR MIT

// Package gofont registers the Go fonts in the font registry.
//
// See https://blog.golang.org/go-fonts for a description of the
// fonts, and the golang.org/x/image/font/gofont packages for the
// font data.
package gofontu

import (
	"fmt"

	"gioui.org/font"
	"gioui.org/font/opentype"
	"gioui.org/text"
	"github.com/dolanor/gofontu/ubuntub"
	"github.com/dolanor/gofontu/ubuntubi"
	"github.com/dolanor/gofontu/ubuntul"
	"github.com/dolanor/gofontu/ubuntuli"
	"github.com/dolanor/gofontu/ubuntum"
	"github.com/dolanor/gofontu/ubuntumi"
	"github.com/dolanor/gofontu/ubuntumonob"
	"github.com/dolanor/gofontu/ubuntumonobi"
	"github.com/dolanor/gofontu/ubuntumonor"
	"github.com/dolanor/gofontu/ubuntumonori"
	"github.com/dolanor/gofontu/ubuntur"
	"github.com/dolanor/gofontu/ubunturi"
)

func Register() {
	register(text.Font{}, ubuntur.TTF)
	register(text.Font{Style: text.Italic}, ubunturi.TTF)
	register(text.Font{Weight: text.Bold}, ubuntub.TTF)
	register(text.Font{Style: text.Italic, Weight: text.Bold}, ubuntubi.TTF)
	register(text.Font{Weight: text.Medium}, ubuntum.TTF)
	register(text.Font{Weight: text.Medium, Style: text.Italic}, ubuntumi.TTF)
	register(text.Font{Variant: "Mono"}, ubuntumonor.TTF)
	register(text.Font{Variant: "Mono", Weight: text.Bold}, ubuntumonob.TTF)
	register(text.Font{Variant: "Mono", Weight: text.Bold, Style: text.Italic}, ubuntumonobi.TTF)
	register(text.Font{Variant: "Mono", Style: text.Italic}, ubuntumonori.TTF)
	register(text.Font{Variant: "Smallcaps"}, ubuntul.TTF)
	register(text.Font{Variant: "Smallcaps", Style: text.Italic}, ubuntuli.TTF)
}

func register(fnt text.Font, ttf []byte) {
	face, err := opentype.Parse(ttf)
	if err != nil {
		panic(fmt.Sprintf("failed to parse font: %v", err))
	}
	fnt.Typeface = "ubuntu"
	font.Register(fnt, face)
}
