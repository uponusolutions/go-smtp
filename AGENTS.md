# Agent guidelines

## Checks

- `make race` runs all tests with the race detector, `make lint` runs golangci-lint.
  Both must pass before a change is done.

## Options and defaults

- Default option values are written inline where the defaults are created
  (`client.DefaultConfig()`, `server.New()`), each with a short comment explaining
  the value (e.g. the RFC it comes from). Do not add `Default*` constants for them.
- The `With…` option doc comment states the default value ("Defaults to 5 minutes")
  and what zero means.
- Limits and timeouts are bounded by default; a zero (or negative) value disables them.
- Tests assert the literal default values.
