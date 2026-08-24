# PKCE for Go

[![Go CI][go-ci-svg]][go-ci-url]
[![Go Lint][go-lint-svg]][go-lint-url]
[![Go SAST][go-sast-svg]][go-sast-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![Visualization][viz-svg]][viz-url]
[![License][license-svg]][license-url]

 [go-ci-svg]: https://github.com/grokify/go-pkce/actions/workflows/go-ci.yaml/badge.svg?branch=main
 [go-ci-url]: https://github.com/grokify/go-pkce/actions/workflows/go-ci.yaml
 [go-lint-svg]: https://github.com/grokify/go-pkce/actions/workflows/go-lint.yaml/badge.svg?branch=main
 [go-lint-url]: https://github.com/grokify/go-pkce/actions/workflows/go-lint.yaml
 [go-sast-svg]: https://github.com/grokify/go-pkce/actions/workflows/go-sast-codeql.yaml/badge.svg?branch=main
 [go-sast-url]: https://github.com/grokify/go-pkce/actions/workflows/go-sast-codeql.yaml
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/go-pkce
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/go-pkce
 [docs-mkdoc-svg]: https://img.shields.io/badge/Go-dev%20guide-blue.svg
 [docs-mkdoc-url]: https://grokify.github.io/go-pkce
 [viz-svg]: https://img.shields.io/badge/visualizaton-Go-blue.svg
 [viz-url]: https://mango-dune-07a8b7110.1.azurestaticapps.net/?repo=grokify%2Fgo-pkce
 [loc-svg]: https://tokei.rs/b1/github/grokify/go-pkce
 [repo-url]: https://github.com/grokify/go-pkce
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/go-pkce/blob/main/LICENSE

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