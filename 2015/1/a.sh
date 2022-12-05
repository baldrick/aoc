up=$(sed 's/)//g' <"$1" | wc -m)
(( up = up - 1 ))

down=$(sed 's/(//g' <"$1" | wc -m)
(( down = down - 1 ))

(( floor = up - down))
echo $floor