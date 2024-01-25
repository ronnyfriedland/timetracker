# Timetracker

Tracks every 5 minutes if system is running to store the beginning and end of using the system.

## Parameters

The timetracker application provides the following arguments which can be passed:

| Property      | Description                                                                |
|---------------|----------------------------------------------------------------------------|
| configpath    | Defines the location of the necessary files, default: /var/lib/timetracker |

## Execution

The application is triggered by a systemd timer which triggers the application via systemd unit.

## License

This application is published under the [MIT license](LICENSE).