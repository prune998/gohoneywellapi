# PeakHoursD

This is an HTTP(s) server acting as a Peak Hour API

## API

There's only one API endpoint: `/peakhours`

Answer if a Json Payload fomated as a stream of start/end blocks.
Time is UTC and should contain the timezone information.

```json
[
  {
    "start": "2020-11-13T1720:00:02-05:00",
    "end": "2020-11-13T23:00:02-05:00",
  },
  {
    "start": "2020-11-14T1720:00:02-05:00",
    "end": "2020-11-14T23:00:02-05:00",
  }
]
```

## Usage
