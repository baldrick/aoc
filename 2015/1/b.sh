floor=0
n=0
while read -n1 c; do
    (( n = n + 1 ))
    if [[ "$c" == "(" ]]
    then
        (( floor = floor + 1 ))
    else
        (( floor = floor - 1 ))
    fi
    if [[ $floor -eq -1 ]]
    then
        echo "Entered basement at position $n"
        exit 0
    fi
done < $1