# PKCE for Go

[![CI Status][build-status-svg]][build-status-url]
[![Lint Status][lint-status-svg]][lint-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Code Coverage][codecov-svg]][codecov-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

 [build-status-svg]: https://github.com/grokify/go-pkce/actions/workflows/ci.yaml/badge.svg?branch=master
 [build-status-url]: https://github.com/grokify/go-pkce/actions/workflows/ci.yaml
 [lint-status-svg]: https://github.com/grokify/go-pkce/actions/workflows/lint.yaml/badge.svg?branch=master
 [lint-status-url]: https://github.com/grokify/go-pkce/actions/workflows/lint.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/go-pkce
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/go-pkce
 [codecov-svg]: https://codecov.io/gh/grokify/go-pkce/branch/master/graph/badge.svg
 [codecov-url]: https://codecov.io/gh/grokify/go-pkce
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

```go
import("github.com/grokify/go-pkce")

func main() {
  // Create a code_verifier with default 32 byte length.
  codeVerifier := pkce.NewCodeVerifier(-1)

  // Create a code_verifier with a custom length (32-96 bytes)
  codeVerifier, err := pkce.NewCodeVerifierWithLength(96)

  // Create a code_challenge using `S256`
  codeChallenge := pkce.CodeChallengeS256(codeVerifier)
}
```

## Usage with `oauth2`

```go
import(
  "context"

  "github.com/grokify/go-pkce"
  "golang.org/x/oauth2"
)

func main() {
  // Create a code_verifier with default 32 byte length.
  codeVerifier := pkce.NewCodeVerifier()

  // Create a code_challenge using `S256`
  codeChallenge := pkce.CodeChallengeS256(codeVerifier)

  // Create authorization_code URL using `oauth2.Config`
  authURL := oauth2Config.AuthCodeURL(
    "myState",
    oauth2.SetAuthURLParam(pkce.ParamCodeChallenge, codeChallenge),
    oauth2.SetAuthURLParam(pkce.ParamCodeChallengeMethod, pkce.MethodS256))

  // ... retrieve authorization_code ...

  // Exchange the authorization_code for a token with PKCE.
  token, err := oauth2Config.Exchange(
    context.Background(),
    "myCode",
    oauth2.SetAuthURLParam(pkce.ParamCodeVerifier, codeVerifier),
  )
}
```

## Similar projects

* [`github.com/nirasan/go-oauth-pkce-code-verifier`](https://github.com/nirasan/go-oauth-pkce-code-verifier)