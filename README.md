# CRON Expression Parser

A Cron Expression Parser is a tool designed to interpret and expand cron expressions, which are commonly used in Unix-like operating systems to schedule jobs or tasks to run at specific times. 


#### Key Components of a Cron Expression
A standard cron expression consists of five fields (excluding special strings), each representing a specific unit of time:

    1. Minute (0 - 59): Specifies the exact minute when the task should run.
    2. Hour (0 - 23): Specifies the hour of the day.
    3. Day of Month (1 - 31): Specifies the day of the month.
    4. Month (1 - 12): Specifies the month of the year.
    5. Day of Week (0 - 6): Specifies the day of the week.
## Usage

```bash
  go run cmd/main.go "* * * * * /cmd/to/execute"
```
    
## How to run
    1. Install golang (version > go 1.20) if not installed.
    2. Clone the git repo using the following command:
        git clone https://github.com/code-dagger/cron-parser.git
            or
        git clone git@github.com:code-dagger/cron-parser.git
    3. Navigate to the cloned directory.
    4. Run the following command in the terminal:
        
        > go run cmd/main.go "1 1 1-5 1 1 /cmd/to/execute"
        
        You will get the following output:

        minute        1 
        hour          1 
        day of month  1 2 3 4 5 
        month         1 
        day of week   1 
        command       /cmd/to/execute

## Example

```
> go run cmd/main.go "*/15 0 1,15 * 1-5 /usr/bin/find"

minute        0 15 30 45 
hour          0 
day of month  1 15 
month         1 2 3 4 5 6 7 8 9 10 11 12 
day of week   1 2 3 4 5 
command       /usr/bin/find
```

```
> go run cmd/main.go "0 0,12 1 */2 * /usr/bin/find"

minute        0 
hour          0 12 
day of month  1 
month         1 3 5 7 9 11 
day of week   0 1 2 3 4 5 6 
command       /usr/bin/find
```

```
> go run cmd/main.go "0 4 1-20/4 * * /usr/bin/find"

minute        0 
hour          4 
day of month  1 5 9 13 17 
month         1 2 3 4 5 6 7 8 9 10 11 12 
day of week   0 1 2 3 4 5 6 
command       /usr/bin/find
```

```
> go run cmd/main.go "*/5 4,6,8 1-20/4 */2 6 /usr/bin/find"

minute        0 5 10 15 20 25 30 35 40 45 50 55 
hour          4 6 8 
day of month  1 5 9 13 17 
month         1 3 5 7 9 11 
day of week   6 
command       /usr/bin/find
```