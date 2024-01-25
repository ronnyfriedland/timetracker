# Timetracker

Tracks every 5 minutes if system is running to store the beginning and end of using the system.

## Parameters

The timetracker application provides the following arguments which can be passed:

| Property      | Description                                                                |
|---------------|----------------------------------------------------------------------------|
| configpath    | Defines the location of the necessary files, default: /var/lib/timetracker |

## Execution

The application is triggered by a systemd timer which triggers the application via systemd unit.

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

## License

This application is published under the [MIT license](LICENSE).