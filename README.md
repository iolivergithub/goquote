# goquote

WARNING: Work in progress but this will more or less replicate the tpm2tools tpm2_quote with a few additions.

Goquote will return a formatted TPMS_ATTEST structure from your TPM for a given set of PCR registers and signed with a given key, typically the TPM2.0 Attestation Key.

## Compiling

Firstly ensure that the go-tpm libraries are installed

```bash
go get github.com/google/go-tpm/tpm2
```

then compile and run (if build is successful)

```bash
go build goquote.go
./goquote PARAMETERS GO HERE
```
