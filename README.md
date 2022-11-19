# OEMEdit: A CLI Tool for Editing Windows OEM Information

ARCHIVED: The latest Windows 10 doesn't support this feature.

Note: `Logo`, `Manufacturer`, `Model`, `SupportHours` and `SupportPhone` are [deprecated in Get Help app](https://docs.microsoft.com/en-us/windows-hardware/customize/desktop/unattend/microsoft-windows-shell-setup-oeminformation) but are still used in "System" control panel.


## What's this?

A tool for doing this:

![docs/system.png](https://raw.githubusercontent.com/ebiiim/oemedit/main/docs/system.png)

### How does the tool change OEM information?

Just editing a few registry values.

## Usage

### Get OEM information

1. Run the app with `get` subcommand:
```ps1
OEMEdit.exe get
```

2. Then you can get results like:
```yaml
OEMInformation:
  Logo: ""
  Manufacturer: ""
  Model: ""
  SupportHours: ""
  SupportPhone: ""
  SupportURL: ""
```

3. You can backup your current OEM info:
```ps1
OEMedit.exe get > current.yaml
```

### Set OEM information

1. Edit config like:
```yaml
OEMInformation:
  Logo: C:\path\to\logo.bmp
  Manufacturer: you
  Model: super cool model name here
  SupportHours: ""
  SupportPhone: ""
  SupportURL: ""
```

2. Run the app with `set` subcommand and input config to stdin:
```ps1
OEMEdit.exe set < oem.yaml
```
