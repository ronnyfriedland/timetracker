# Timetracker

Tracks every 15 minutes if system is running to store the beginning and end of using the system.

[![Anchore Syft SBOM scan](https://github.com/ronnyfriedland/timetracker/actions/workflows/anchore-syft.yml/badge.svg)](https://github.com/ronnyfriedland/timetracker/actions/workflows/anchore-syft.yml)
[![CodeQL](https://github.com/ronnyfriedland/timetracker/actions/workflows/codeql.yml/badge.svg)](https://github.com/ronnyfriedland/timetracker/actions/workflows/codeql.yml)

## Parameters

The timetracker application provides the following arguments which can be passed:

| Property      | Description                                                                |
|---------------|----------------------------------------------------------------------------|
| archivedata   | Enables archiving timetracker status to excel archive file, default: false |
| configpath    | Defines the location of the necessary files, default: /var/lib/timetracker |

## Execution

The application is triggered by a systemd timer which triggers the application via systemd unit.

*Note:* Running timetracker with ystemd unit uses the default property values. To change it you have to modify the unit file.

To enable the timer you have to (requires root privileges):

### enable the timer

```shell
systemctl enable timetracker.timer
```

### start the timer

```shell
systemctl start timetracker.timer
```

To verify if the timer is running you can check it using:

```shell
systemctl list-timers
```

### Show results

The aggregated data can be displayed using `journalctl`:
```shell
journalctl -u timetracker.service -t timetracker
```

## License

This application is published under the [MIT license](LICENSE).
