import datetime


def get_date_and_time(raw_entry):
    fields = raw_entry.split()
    str_date = fields[0][1:len(fields[0])]
    str_time = fields[1][0:len(fields[1])-1]
    return str_date, str_time


def convert_to_datetime(str_date, str_time):
    ymd = str_date.split("-")
    year = int(ymd[0])
    month = int(ymd[1])
    day = int(ymd[2])

    hm = str_time.split(":")
    hour = int(hm[0])
    minute = int(hm[1])

    return datetime.datetime(year, month, day, hour, minute)


def convert_entry_to_datetime(raw_entry):
    # [1518-07-28 00:10] falls asleep
    str_date, str_time = get_date_and_time(raw_entry)
    return convert_to_datetime(str_date, str_time)


def time_sort(raw_entry):
    return convert_entry_to_datetime(raw_entry).timestamp()


def get_ordered_log_entries(filename):
    raw_log = open(filename)
    log = raw_log.readlines()
    log.sort(key=time_sort)
    return log


# Log entries are of the form:
# [yyyy-MM-dd hh:mm] Guard #x begins shift
# [yyyy-MM-dd hh:mm] falls asleep
# [yyyy-MM-dd hh:mm] wakes up
def check_valid_guard(current_guard):
    if current_guard == -1:
        print("Error: no current guard")


def check_awake(fell_asleep_time, awoke_time):
    if awoke_time != -1:
        print("Error: not awake")
    if fell_asleep_time != -1:
        print("Error: already asleep")


def check_asleep(fell_asleep_time, awoke_time):
    if awoke_time != -1:
        print("Error: awake")
    if fell_asleep_time == -1:
        print("Error: already awake")


def get_minute(entry):
    fields = entry.split(":")
    raw_minute = fields[1]
    minute = raw_minute[0:2]
    return int(minute)


def awake_for_an_hour():
    awake = []
    for minute in range(60):
        awake.append(0)
    return awake


def add_sleep_minutes(current_guard_sleeps, fell_asleep_time, awoke_time):
    for minute in range(fell_asleep_time, awoke_time):
        current_guard_sleeps[minute] += 1


def add_sleep(current_guard, fell_asleep_time, awoke_time, guard_sleeps):
    if current_guard not in guard_sleeps:
        guard_sleeps[current_guard] = awake_for_an_hour()
    add_sleep_minutes(guard_sleeps[current_guard], fell_asleep_time, awoke_time)


def process_entry(entry, current_guard, fell_asleep_time, awoke_time, guard_sleeps):
    fields = entry.split()

    action = fields[2]
    if action == "Guard":
        check_awake(fell_asleep_time, awoke_time)
        raw_current_guard = fields[3]
        current_guard = int(raw_current_guard[1:len(raw_current_guard)])
        fell_asleep_time = -1
        awoke_time = -1
    elif action == "falls":
        check_valid_guard(current_guard)
        check_awake(fell_asleep_time, awoke_time)
        fell_asleep_time = get_minute(entry)
        awoke_time = -1
    elif action == "wakes":
        check_valid_guard(current_guard)
        check_asleep(fell_asleep_time, awoke_time)
        awoke_time = get_minute(entry)
        add_sleep(current_guard, fell_asleep_time, awoke_time, guard_sleeps)
        fell_asleep_time = -1
        awoke_time = -1
    else:
        print("Error, can't parse entry: " + entry)
    return current_guard, fell_asleep_time, awoke_time


def calculate_guard_sleeps(log):
    guard_sleeps = {}
    current_guard = -1
    fell_asleep_time = -1
    awoke_time = -1
    for entry in log:
        current_guard, fell_asleep_time, awoke_time = process_entry(entry, current_guard, fell_asleep_time, awoke_time, guard_sleeps)
    return guard_sleeps


def sum_up_sleep_minutes(guard_sleep_time):
    minutes_asleep = 0
    for minute in guard_sleep_time:
        minutes_asleep += minute
    return minutes_asleep


def find_sleepiest_guard(guard_sleeps):
    sleepiest_guard = -1
    sleepiest_guard_minutes_asleep = 0
    for guard in guard_sleeps:
        minutes_asleep = sum_up_sleep_minutes(guard_sleeps[guard])
        print("Guard ", guard, " sleeps for ", minutes_asleep, " minutes")
        if minutes_asleep > sleepiest_guard_minutes_asleep:
            sleepiest_guard = guard
            sleepiest_guard_minutes_asleep = minutes_asleep
    return sleepiest_guard


def find_minute_most_asleep(guard_sleep_time):
    minute_most_asleep = 0
    for minute in range(60):
        if guard_sleep_time[minute] > guard_sleep_time[minute_most_asleep]:
            minute_most_asleep = minute
    return minute_most_asleep


def part1(filename):
    log = get_ordered_log_entries(filename)
    guard_sleeps = calculate_guard_sleeps(log)
    for guard in guard_sleeps:
        print(guard, ": ", guard_sleeps[guard])
    sleepiest_guard = find_sleepiest_guard(guard_sleeps)
    minute_most_asleep = find_minute_most_asleep(guard_sleeps[sleepiest_guard])
    print(sleepiest_guard, " * ", minute_most_asleep, " = ", sleepiest_guard * minute_most_asleep)


def find_most_popular_minute_with_guard(guard_sleeps):
    sleepiest_guard = next(iter(guard_sleeps))
    sleepiest_minute = guard_sleeps[sleepiest_guard][0]
    for guard in guard_sleeps:
        minute = find_minute_most_asleep(guard_sleeps[guard])
        if guard_sleeps[guard][minute] > guard_sleeps[sleepiest_guard][sleepiest_minute]:
            sleepiest_guard = guard
            sleepiest_minute = minute
    return sleepiest_guard, sleepiest_minute


def part2(filename):
    log = get_ordered_log_entries(filename)
    guard_sleeps = calculate_guard_sleeps(log)
    for guard in guard_sleeps:
        print(guard, ": ", guard_sleeps[guard])
    guard, sleepiest_minute = find_most_popular_minute_with_guard(guard_sleeps)
    print(guard, " * ", sleepiest_minute, " = ", guard * sleepiest_minute)


# part1("input.txt")
# part1("test.txt")
part2("input.txt")
# part2("test.txt")