# GO-EST ![image](https://github.com/acavella/go-est/blob/b011511623f79e697367088a17b517ee5d82ba3b/docs/Go-EST_Blue.png)
A RFC 7030 compliant EST client written in Go.

## Requirements
- 
- 

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