# Overview
Golang client library for interacting with the [Busy Bar](https://busy.app/) via its HTTP API.

# Usage
## As a library
There's a client implemented in [client.go](./client/client.go)

As per the example CLI application, you can create it:

```go
// Create a new client with the specified endpoint
ctx := context.Background()
c := client.NewClient("http://" + ipAddress + "/")
```

And then use it to make different endpoint calls:

```go
status, err := c.GetStatus(ctx, nil)
if err != nil {
    fmt.Printf("Error getting device status: %v\n", err)
    return
}
status.PrettyPrint()
```

## CLI tool
Build it with `Make` or use `go run`:

```bash
 go run cli/*.go -c status

Status
┏━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃ Field            ┃ Value                     ┃
┡━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━━━━━┩
│ Firmware Version │ 1.1.1                     │
│ System Uptime    │ 00d 00h 21m 47s           │
│ System Boot Time │ 2026-07-29T16:31:40+01:00 │
│ Power State      │ discharging               │
│ Battery Charge   │ 100%                      │
│ Battery Voltage  │ 4179 mV                   │
│ Battery Current  │ 0 mA                      │
│ USB Voltage      │ 5112 mV                   │
└──────────────────┴───────────────────────────┘
```

It assumes the Busy Bar is connected over USB/Ethernet on the default address. To use over Wi-Fi, specify the address with `-ip` and the password with '-p'.

The flag `-c` is for the command you want to run. The currently implemented commands are:

- `status`: prints the current device status 
- `info`: prints the device hardware info
- `settings`: prints the current settings of the bar
- `version`: prints the current API version running on the bar.
- `firmware`: prints details of the installed firmware



# References
- [Online HTTP reference documentation](https://docs.busy.app/bar/dev/http-api)
- [HTTP API Spec (over USB/Ethernet to a local device)](http://10.0.4.20/docs/#/)