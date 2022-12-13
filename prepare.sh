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
puzzle=$year/$day/puzzle

[ -d $year/$day ] || mkdir -p $year/$day

if [ -f ${puzzle} ]
then
    echo "${puzzle} already exists"
else
    session=$(cat session)
    curl --cookie ${session} -o ${puzzle} https://adventofcode.com/${year}/day/${day}/input
fi

if [ ! -f $year/$day/a.jl ]
then
    cp templates/a.jl $year/$day
fi

echo "cd $year/$day" | pbcopy
echo "Now cd $year/$day (already in the clipboard for you!)"
