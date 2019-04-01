# jwt2pem

This tiny tool converts jwt public keys into pem files.

## Usage

jwt2pem reads from stdin and writes to stdout.
There are no parameters.

Example:
```bash
jwt2pem < key.json > key.pem
```

## Format

jwt keys are expected to have the following format:

```json
{
  "kty": "RSA",
  "n": "some-base64-encoded-large-prime-factor",
  "e": "some-base64-encoded-not-so-large-prime-factor"
}
```

Here is an example key with `n=65537` and `e=4097`:
```json
{
  "kty": "RSA",
  "n": "AQAB",
  "e": "EAE"
}
```
