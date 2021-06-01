CONTENT='Content-Type: application/json'
DATA='{"NameStudent": "Daniela", "Subject": "Mates", "Grade": 98.0}'

set -x
curl --header "$CONTENT" --request PUT --data "$DATA" 127.0.0.1:1789/estudiante/$1 && echo "\n"
