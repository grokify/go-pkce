# PKCE for Go

[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

 [build-status-svg]: https://github.com/grokify/go-pkce/workflows/go%20build/badge.svg
 [build-status-url]: https://github.com/grokify/go-pkce/actions
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/go-pkce
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/go-pkce
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/go-pkce
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/go-pkce
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/go-pkce/blob/master/LICENSE

`go-pkce` package contains an implementation for OAuth 2.0 PKCE spec, [IETF RFC 7636](https://datatracker.ietf.org/doc/html/rfc7636).

## Installation

```
go get github.com/grokify/go-pkce
```

Or you can manually git clone the repository to
`$(go env GOPATH)/src/github.com/grokify/go-pkce`.

## Usage

```
import("github.com/grokify/go-pkce")

// Create a code_verifier with default 32 byte length.
codeVerifier := NewCodeVerifier()

// Create a code_verifier with a custom length (32-96 bytes)
codeVerifier, err := NewCodeVerifierWithLength(96)

// Create a code_challenge using `S256`
codeChallenge := CodeChallengeS256(codeVerifier)
```

## Usage with `oauth2`

```
import(
    "github.com/grokify/go-pkce"
    "golang.org/x/oauth2
)

// Create a code_verifier with default 32 byte length.
codeVerifier := NewCodeVerifier()

// Create a code_challenge using `S256`
codeChallenge := CodeChallengeS256(codeVerifier)

// Create authorization_code URL using `oauth2.Config`
authURL := oauth2Config.AuthCodeURL(
    "myState",
    oauth2.SetAuthURLParam(pkce.ParamCodeChallenge, codeChallenge),
    oauth2.SetAuthURLParam(pkce.ParamCodeChallengeMethod, pkce.MethodS256))
```