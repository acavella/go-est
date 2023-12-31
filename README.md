![GO-EST](https://github.com/acavella/go-est/blob/0f93ff9904842bfda8dc09dca691380c1d72f0ea/assets/Go-EST_Blue.png)

![GitHub go.mod Go version (branch & subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/acavella/go-est/main)
![GitHub last commit (branch)](https://img.shields.io/github/last-commit/acavella/go-est/main)
![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/acavella/go-est/go.yml)


___
A RFC 7030 compliant EST client written in Go.

## References
- [RFC 7030](https://datatracker.ietf.org/doc/html/rfc7030)
- [RFC 8295](https://datatracker.ietf.org/doc/html/rfc8295)
- ISC CertAgent 7.0.9.x Administration Guide

## Installation
- Step 1
- Step 2

## Usage
- Step 1
- Step 2

## Process Diagram
```
┌────────────┐                     ┌────────────┐                      ┌────────────┐
│ EST Client │                     │ EST Server │                      │   EST CA   │
└─────┬──────┘                     └──────┬─────┘                      └──────┬─────┘
      │                                   │                                   │
      │                                   │                                   │
      │                                   │                                   │
      │    (EST) Request certification    │                                   │
      ├──────────────────────────────────►│                                   │
      │                                   │                                   │
      │             Trust chain           │                                   │
      │◄──────────────────────────────────┤                                   │
      │                                   │                                   │
      │  Validate chain                   │                                   │
      ├───────────────────┐               │                                   │
      │                   │               │                                   │
      │                   │               │                                   │
      │◄──────────────────┘               │                                   │
      │                                   │                                   │
      │Generate key and CSR               │                                   │
      ├───────────────────┐               │                                   │
      │                   │               │                                   │
      │                   │               │                                   │
      │◄──────────────────┘               │                                   │
      │                                   │                                   │
      │ (EST) PKCS#10 certificate request │                                   │
      ├──────────────────────────────────►│                                   │
      │                                   │                                   │
      │                                   │Validate client credent            │
      │                                   │(Certificate auth)                 │
      │                                   ├─────────────────────┐             │
      │                                   │                     │             │
      │                                   │                     │             │
      │                                   │◄────────────────────┘             │
      │                                   │                                   │
      │                                   │         Request certificate       │
      │                                   ├──────────────────────────────────►│
      │                                   │                                   │
      │                                   │              Certificate          │
      │                                   │◄──────────────────────────────────┤
      │                                   │                                   │
      │        PKCS#7 Certificate         │                                   │
      │◄──────────────────────────────────┤                                   │
      │                                   │                                   │
      │                                   │                                   │
      │                                   │                                   │
```
