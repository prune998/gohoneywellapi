# GoHwCTL

This is a command line tool to configure the server and CRUD the schedule

## Schedule format

```json
{
  startdate: "2020-11-13T23:00:02", // date formated as <year>-<month>-<day>T<hour>:<min>:<sec>
  temp: 22.3,                       // float celcius temp
  tempunit: "celcius",              // celcius or farheinehit
}
```

## command line

```sh
Usage of ./gohwctl:
  -action="add": action like add, del
  -key="": key used to talk to the server
  -loglevel="warning": the log level to display (debug,info,error,warning)
  -server="http://localhost": URL to connect to the server
  -startdate="": start date for the change, like <year>-<month>-<day>T<hour>:<min>:<sec>
  -temp=20: Temp to add
  -unit="celcius": temperature unit
  -version=false: Show version and quit
  -y=false: apply without asking
```

To schedule a 20 celcius starting at 23:00:02:

```sh
./gohwctl -startdate=2020-11-13T23:00:02 -y
```