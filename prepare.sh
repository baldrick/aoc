init() {
    if [ -z "$1" ]
    then
        year=$(date | cut -d ' ' -f 4)
        day=$(date | cut -d ' ' -f 2)
    else
        year=$1
        day=$2
        if [ -z "$day" ]
        then
            echo "If year is specified, day must also be specified..."
            exit 1
        fi
    fi

    [ -d $year/$day ] || mkdir -p $year/$day
}

fileReplace() {
    search=$1
    replace=$2
    infile=$3
    outfile=$4
    force=$5
    if [[ -f $outfile && "$force" != "force" ]]
    then
        echo "$outfile already exists, delete it if you want this script to replace it"
    else
        sed "s/$search/$replace/g" <"$infile" >"$outfile"
    fi
}

updateApp() {
    clidays="$(seq 1 25)"
    deps=""
    imports=""
    cmds=""
    for d in $clidays
    do
        if [[ $d -le $day ]]
        then
            imports="$imports\\n    day$d \"github.com\/baldrick\/aoc\/2023\/$d\""
            cmds="$cmds\\n            *day$d.A, *day$d.B,"
            deps="$deps\\n        \"\/\/2023\/$d\:${d}\","
        fi
    done
    fileReplace "{{DEPS}}" "$deps" templates/template.aoc.BUILD.bazel ${year}/BUILD.bazel force
    fileReplace "{{IMPORTS}}" "$imports" templates/aoc.go.template ${year}/aoc.go.2 force
    fileReplace "{{CMDS}}" "$cmds" ${year}/aoc.go.2 ${year}/aoc.go force
    rm ${year}/aoc.go.2
}

createCode() {
    tmp=/tmp/prepare_$$

    fileReplace "{{DAY}}" "$day" templates/a.go.template $tmp
    fileReplace "{{YEAR}}" "$year" $tmp $year/$day/a.go

    fileReplace "{{DAY}}" "$day" templates/a_test.go.template $tmp force
    fileReplace "{{YEAR}}" "$year" $tmp $year/$day/a_test.go

    fileReplace "{{DAY}}" "$day" templates/template.BUILD.bazel $tmp force
    fileReplace "{{YEAR}}" "$year" $tmp $year/$day/BUILD.bazel

    rm $tmp
}

getPuzzle() {
    puzzle=$year/$day/puzzle.txt
    if [ -f ${puzzle} ]
    then
        echo "${puzzle} already exists"
    else
        session=$(cat session)
        tmp=/tmp/curl.$$
        curl --cookie ${session} -o ${puzzle} https://adventofcode.com/${year}/day/${day}/input 2>$tmp
        if [[ $? -ne 0 ]]
        then
            echo "Failed to retrieve puzzle:"
            cat $tmp
        fi
        rm $tmp
    fi
}

showInstructions() {
    testcmd="blaze test --test_output=all $year/$day/${day}_test"
    which pbcopy 1>/dev/null 2>/dev/null
    if [ $? -eq 0 ]
    then
        echo $testcmd | pbcopy
        echo "To test, run: $testcmd (already in the clipboard for you)"
    else
        echo "To test, run: $testcmd"
    fi
    echo "To run part A of the puzzle, run: blaze run $year -- day${day}A"
}

init $@
echo "Preparing $year/$day"
updateApp
createCode
getPuzzle
showInstructions
