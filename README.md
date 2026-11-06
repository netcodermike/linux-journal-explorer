# Linux Journal Explorer

Query a bounded systemd journal selection and emit structured, filtered JSON records.

[![Build and verify](https://github.com/netcodermike/linux-journal-explorer/actions/workflows/verify.yml/badge.svg)](https://github.com/netcodermike/linux-journal-explorer/actions/workflows/verify.yml)

**Language:** Go · **Platform:** Linux-focused · **App type:** Linux journalctl filtering and session reports

First-release learning implementation, with explicit scope below. This is one independently runnable app, not a placeholder in a combined repository.

## Features

- Unit and since filters
- 1–500 record bound
- Stable normalized JSON fields
- No service changes

## Requirements

Go 1.22+ and the documented Linux facilities. No external Go modules.

## Quick start

```sh
git clone https://github.com/netcodermike/linux-journal-explorer.git
cd linux-journal-explorer
go build -o app .
./app --since today --limit 20
./app --unit example.service --since yesterday --limit 100
```

Commands run from the repository root. Lines marked for a particular operating system apply only there. Language tools are prerequisites, not installed by the app. Close desktop windows normally; terminal/server apps stop with Ctrl+C unless their usage notes say otherwise.

## Example workflow

Run ./app --since today --limit 20, then add --unit a-known.service. Compare normalized messages to journalctl and redirect JSON to a separate report if desired.

## Architecture, configuration, and data

main.go calls journalctl with an argument array, --no-pager, JSON output, a capped record count, and optional unit/since expressions. It parses realtime timestamp microseconds, unit, priority, and message without executing a shell. Tests parse synthetic records and reject malformed JSON. journalctl permissions determine visibility. --help lists flags.

## Verification

Build with the quick-start command where needed, then run:

```sh
go test ./...
```

Manual acceptance: Run ./app --since today --limit 20, then add --unit a-known.service. Compare normalized messages to journalctl and redirect JSON to a separate report if desired. Also test malformed inputs and an unwritable output destination.

See [validation record](docs/VALIDATION.md) for observed checks and their limits. CI results verify only what the workflow runs; they do not establish native GUI correctness or production readiness.

## Limitations and next steps

Linux/systemd only. No live following, journal deletion, service management, graphical timeline, rich binary-message decoding, or timestamp timezone conversion. A single over-2-MiB row or malformed JSON fails the query. Logs may contain private data; exports are ordinary JSON.

## Troubleshooting

Use the exact platform/toolchain above. Preserve originals and use synthetic fixtures first. Errors are displayed in the UI or stderr. A headless machine can run the model tests but cannot verify native GUI/audio behavior. Output paths are chosen explicitly; keep backups of important data.

## Source map

- `go.mod`
- `main.go`
- `main_test.go`

## Development and license

Read [contributing](CONTRIBUTING.md) before changing the app, and [security/privacy](SECURITY.md) before supplying private data. Use synthetic fixtures for tests. No telemetry, paid API key, or account is required by the app itself. Dependencies, when present, are stated in Requirements.

[MIT license](LICENSE), copyright 2026 netcodermike.
