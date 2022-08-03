# Icinga2 API exporter

Prometheus exporter for the Icinga2 API.

## Installation and Usage

The `icinga2_exporter` listens on HTTP port 9665 by default. See the --help output for more options.

Further configuration is done via environment variables.

### Environment Variables
Name | Description
-----|------------
`ICINGA2_EXPORTER_BASE_URL` | URL to the Icinga2 API (default "https://localhost:5665/v1")
`ICINGA2_EXPORTER_USERNAME` | Username for the Icinga2 API (default "root")
`ICINGA2_EXPORTER_PASSWORD` | Password for the Icinga2 API (default "password")
`ICINGA2_EXPORTER_TLS_INSECURE` | Skip TLS verification (default "false")
