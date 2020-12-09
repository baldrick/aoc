#   -->A--->B--
#  /    \      \
# C      -->D----->E
#  \           /
#   ---->F-----

# task next
# C-AF
# A-BD
# B-E
# D-E
# F-E
#
# task pre-reqs
# A<-C
# F<-C
# B<-A
# D<-A
# E<-BDF
#
# start with C because it has no pre-requisites
# done = C
# available tasks = AF
# all A pre-reqs (C) are done so done = CA
# available tasks = BDF
# all B pre-reqs (A) are done so done = CAB
# available tasks = DEF
# all D pre-reqs (A) are done so done = CABD
# available tasks = EF
# task E pre-reqs (BDF) not all done so skip
# task F pre-reqs (C) are done so done = CABDF
# available tasks = E
# CABDFE
#
# if C in pre, insert A into its post
# else insert C-A into list

# start C, next = AF
# start A => next = BDF
# start B => next = DEF
# start D => next = EF
# complete = CB, next = DF


def insert_task(task_list, task):
    if task in task_list:
        return

    insert_at = len(task_list)
    for index in range(len(task_list)):
        if task_list[index] > task:
            insert_at = index
            break
    task_list.insert(insert_at, task)


def add_task(tasks, prereqs, task, next_task):
    if task in tasks.keys():
        insert_task(tasks[task], next_task)
    else:
        tasks[task] = [next_task]

    if next_task in prereqs.keys():
        insert_task(prereqs[next_task], task)
    else:
        prereqs[next_task] = [task]


def process_instructions(instructions):
    tasks = {}
    prereqs = {}
    for instruction in instructions:
        words = instruction.split()
        task = words[1]
        next_task = words[7]
        add_task(tasks, prereqs, task, next_task)
    return tasks, prereqs

# start with C because it has no pre-requisites
# done = C
# available tasks = AF
# all A pre-reqs (C) are done so done = CA
# available tasks = BDF
# all B pre-reqs (A) are done so done = CAB
# available tasks = DEF
# all D pre-reqs (A) are done so done = CABD
# available tasks = EF
# task E pre-reqs (BDF) not all done so skip
# task F pre-reqs (C) are done so done = CABDF
# available tasks = E


def find_root_tasks(tasks, prereqs):
    all_tasks = set(tasks.keys())
    all_prereqs = set(prereqs.keys())
    root_tasks = [task for task in all_tasks if task not in all_prereqs]
    return root_tasks


def add_next_tasks(available_tasks, tasks_to_add):
    for task in tasks_to_add:
        insert_task(available_tasks, task)


def all_done(completed_tasks, prereqs):
    for task in prereqs:
        if task not in completed_tasks:
            return False
    return True


def get_next_task(available_tasks, prereqs, completed_tasks):
    if len(available_tasks) == 0:
        return False, ""
    for task in available_tasks:
        if task in prereqs:
            if all_done(completed_tasks, prereqs[task]):
                return True, task
        else:
            return True, task
    return False, ""


def ordered_instructions(tasks, prereqs):
    available_tasks = find_root_tasks(tasks, prereqs)
    completed_tasks = ""
    tasks_to_do, current_task = get_next_task(available_tasks, prereqs, completed_tasks)
    while tasks_to_do:
        completed_tasks += current_task
        available_tasks.remove(current_task)
        if current_task in tasks:
            add_next_tasks(available_tasks, tasks[current_task])
        tasks_to_do, current_task = get_next_task(available_tasks, prereqs, completed_tasks)
    return completed_tasks


class Worker:
    def __init__(self, finish_time, current_task):
        self.finish_time = finish_time
        self.current_task = current_task

    def is_finished(self, time):
        return self.finish_time <= time

    def start_task(self, time, new_task, available_tasks):
        self.finish_time = time + ord(new_task) - ord('A') + 1
        self.current_task = new_task
        available_tasks.remove(self.current_task)

    def finish_task(self, completed_tasks, available_tasks, tasks):
        if self.current_task != "":
            completed_tasks += self.current_task
            if self.current_task in tasks:
                add_next_tasks(available_tasks, tasks[self.current_task])
        return completed_tasks


def one_or_more_workers_are_busy(time, workers):
    for worker in workers:
        if not worker.is_finished(time):
            return True


def ordered_instructions_with_elves(number_of_elves, minimum_task_time, tasks, prereqs):
    available_tasks = find_root_tasks(tasks, prereqs)
    completed_tasks = ""
    ignore, first_task = get_next_task(available_tasks, prereqs, completed_tasks)

    workers = []
    for worker in range(number_of_elves+1):
        workers.append(Worker(0, ""))

    time = 0
    workers[0].start_task(time + minimum_task_time, first_task, available_tasks)
    while one_or_more_workers_are_busy(time - 1, workers):
        for worker in range(number_of_elves+1):
            if workers[worker].is_finished(time):
                print("worker", worker, "has finished", workers[worker].current_task)
                completed_tasks = workers[worker].finish_task(completed_tasks, available_tasks, tasks)
                ignore, new_task = get_next_task(available_tasks, prereqs, completed_tasks)
                if new_task == "":
                    print("worker", worker, "is awaiting a task as of time", time)
                else:
                    workers[worker].start_task(time + minimum_task_time, new_task, available_tasks)
        time += 1

    finished = False
    while not finished:
        finished = True
        for worker in range(number_of_elves+1):
            if not workers[worker].is_finished(time):
                finished = False
                break
        if finished:
            break
        time += 1

    return time - 1


def part1(filename):
    f = open(filename)
    instructions = f.readlines()
    tasks, prereqs = process_instructions(instructions)
    print(ordered_instructions(tasks, prereqs))


def part2(filename, number_of_elves, minimum_task_time):
    f = open(filename)
    instructions = f.readlines()
    tasks, prereqs = process_instructions(instructions)
    print(ordered_instructions_with_elves(number_of_elves, minimum_task_time, tasks, prereqs))


# part1("input.txt")
# part2("test.txt", 1, 0)
part2("input.txt", 4, 60)
