![image](https://github.com/acavella/go-est/blob/0f93ff9904842bfda8dc09dca691380c1d72f0ea/assets/Go-EST_Blue.png)
# GO-EST 
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
