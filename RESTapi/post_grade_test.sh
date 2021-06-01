CONTENT='Content-Type: application/json'
DATA='{"NameStudent": "Daniela", "Subject": "Mates", "Grade": 100.0}'

set -x
curl --header "$CONTENT" --request POST --data "$DATA" 127.0.0.1:1789/calificacion && echo "\n"
