bit=$1
carryIn=$2
y=x$bit
x=y$bit

echo "Checking bit $bit with carry in $carryIn"

puzzle=puzzle.txt
inputs=/tmp/$$.in
trap 'rm -f "$inputs"' EXIT

grep $bit puzzle.txt | grep -v \: >$inputs

# x and y = n
N=$(grep AND $inputs | sed 's/ -> [a-z0-9]\+//')
if [[ "$N" != "$x AND $y" ]] && [[ "$N" != "$y AND $x" ]]
then
    echo "Can't find $x AND $y"
    exit 1
fi
N=$(grep AND $inputs | sed 's/.* -> //')
echo n=$N

# x xor y = sumbits
SUMBITS=$(grep XOR $inputs | grep -v '\-> z' | sed 's/.* -> //')
echo sumbits=$SUMBITS

# sumbits xor carryIn = sum
SUM=$(grep $SUMBITS $puzzle | grep "XOR" | grep $carryIn | sed 's/.* -> //')
if [[ "$SUM" != "z$bit" ]]
then
    echo "Incorrect sum output $SUM"
    exit 1
fi

# sumbits AND carryIn = m
M=$(grep $SUMBITS $puzzle | grep "AND" | grep $carryIn | sed 's/.* -> //')
echo m=$M

# m OR n = carryOut
carryOut=$(grep $M $puzzle | grep " OR " | grep $N | sed 's/.* -> //')
echo "carry out = $carryOut"