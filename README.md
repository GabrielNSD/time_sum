# Time Sum CLI

A simple command-line interface tool for summing time durations in multiple formats including hours, minutes, and seconds.

## Features

- Interactive CLI for time summation
- Supports multiple time formats:
  - HH:MM:SS (hours, minutes, seconds)
  - HH:MM (hours, minutes)
  - MM:SS (minutes, seconds)
- Validates time format and ranges
- Real-time total calculation
- Input validation for hours (0-23), minutes (0-59), and seconds (0-59)
- Reset functionality to start over

## Usage

### Building the application

```bash
go build
```

### Running the application

```bash
./time_sum
```

### Commands

- `start` - Begin summing times
- `end` - Finish summing and display total
- `undo` - Undo the last sum
- `reset` - Reset the current sum to zero
- `quit` or `exit` - Exit the program
- `help` - Show available commands

### Time Formats

The tool supports three time formats:

1. **HH:MM:SS** - Hours, minutes, and seconds
   - `02:30:45` (2 hours 30 minutes 45 seconds)
   - `01:15:30` (1 hour 15 minutes 30 seconds)
   - `00:00:30` (30 seconds)

2. **HH:MM** - Hours and minutes
   - `02:30` (2 hours 30 minutes)
   - `1:45` (1 hour 45 minutes)
   - `00:15` (15 minutes)

3. **MM:SS** - Minutes and seconds
   - `30:45` (30 minutes 45 seconds)
   - `05:30` (5 minutes 30 seconds)
   - `00:45` (45 seconds)

**Note**: The tool automatically detects the format based on the number of colons and the value ranges.

### Example Session

```
Time Sum CLI
Commands:
  start - Start summing times
  end   - End summing and show total
  reset - Reset the sum
  undo  - Undo the last time added
  quit  - Exit the program
  help  - Show this help
Time formats:
  HH:MM:SS (e.g., 02:30:45)
  HH:MM    (e.g., 02:30)
  MM:SS    (e.g., 30:45)

> start
Started summing times. Enter times in HH:MM:SS, HH:MM, or MM:SS format or 'end' to finish.
> 02:30:45
Added 02:30:45 (total: 02:30:45)
> 01:15:30
Added 01:15:30 (total: 03:46:15)
> 30:45
Added 30:45 (total: 04:17:00)
> end
Total time: 04:17:00
> quit
Goodbye!
```

## Requirements

- Go 1.21.6 or later 