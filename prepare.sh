if [ -z "$1" ]
then
    year=$(date | cut -d ' ' -f 5)
    day=$(date | cut -d ' ' -f 3)
else
    year=$1
    day=$2
    if [ -z "$day" ]
    then
        echo "If year is specified, day must also be specified..."
        exit 1
    fi
fi
puzzle=$year/$day/puzzle.txt

[ -d $year/$day ] || mkdir -p $year/$day

if [ -f ${puzzle} ]
then
    echo "${puzzle} already exists"
else
    session=$(cat session)
    curl --cookie ${session} -o ${puzzle} https://adventofcode.com/${year}/day/${day}/input
fi

if [ ! -f $year/$day/$template ]
then
    sed "s/{{DAY}}/$day/" <templates/a.go.template | sed "s/{{YEAR}}/$year/" >$year/$day/a.go
    sed "s/{{DAY}}/$day/" <templates/template.BUILD.bazel | sed "s/{{YEAR}}/$year/" >$year/$day/BUILD.bazel
fi

which pbcopy 2>/dev/null
if [ $? -eq 0 ]
then
    echo "cd $year/$day" | pbcopy
    echo "Now cd $year/$day (already in the clipboard for you!)"
else
    echo "Now cd $year/$day"
fi
