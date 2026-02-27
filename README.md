# Icinga2 API exporter

Prometheus exporter for the Icinga2 API.

## Installation and Usage

The `icinga2_exporter` listens on HTTP port 9665 by default.
See the `-help` output for more options.

## Collectors

By default only the `CIB` metrics of the status API are collected.

There are more collectors that can be activated via the CLI.
The tables below list all existing collectors.

| Collector     | Flag       |
| ------------- | ---------- |
| APIListener   | `-collector.apilistener` |
