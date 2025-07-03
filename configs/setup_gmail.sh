#!/bin/bash

EMAIL_OK=""

while [ "$EMAIL_OK" != "ok" ]
do
    read -p "[*] input username your gmail [ e.g user@gmail.com ] : " GMAIL_USERNAME
    if [ -z $GMAIL_USERNAME ]; then
        echo "please try again"
    fi
    EMAIL_OK="ok"
done

echo -n "[*] input your APP GOOGLE PASSWORD [e.g asdd woeur weif asdn ] : "

pwdOK=""
while [ "$pwdOK" != "ok" ];
do
    read -a  myarr
    if [ ${#myarr[@]} != 4 ]; then
        echo "the password must 4 length, pls try again"
        continue
    elif [ -z $myarr ]; then
        echo "pls insert the APP GOOGLE PASSWORD"
        continue
    fi
    
    err="err"
    for word in "${myarr[@]}"; do
        if [ ${#word} != 4 ]; then
            echo "${#word}"
            continue
        fi
        err="nil"
    done
    if [ $err == "nil" ]; then
        pwdOK="ok"

    else
        echo "pls try again"
    fi

done


GMAIL_PASSWORD_APP_GOOGLE="${myarr[@]}"
EMAIL_HOST="smtp.gmail.com"
EMAIL_PORT="587"

CONFIG_FILE=".emailconfig"
cat > "$CONFIG_FILE" <<EOF
GMAIL_USERNAME="$GMAIL_USERNAME"
GMAIL_PASSWORD="${myarr[@]}"
HOST="smtp.gmail.com"
PORT="$EMAIL_PORT"
EOF
chmod 600 "$CONFIG_FILE"

echo "✅ Gmail setup complete!"
echo "💾 Configuration saved to $CONFIG_FILE"
