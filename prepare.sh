day=$(date | cut -d ' ' -f 2)
year=$(date | cut -d ' ' -f 4)
puzzle=$year/$day/puzzle

[ -d $year/$day ] || mkdir -p $year/$day

if [ -f ${puzzle} ]
then
    echo "${puzzle} already exists"
else
    session=$(cat session)
    curl --cookie ${session} -o ${puzzle} https://adventofcode.com/${year}/day/${day}/input
fi
echo "cd $year/$day" | pbcopy
echo "Now cd $year/$day (already in the clipboard for you!)"
