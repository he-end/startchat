#!/bin/bash

# sudo apt update && sudo apt upgrade
# sudo apt install mariadb-server
sudo apt install -y postgresql-common
# sudo /usr/share/postgresql-common/pgdg/apt.postgresql.org.sh


# Input DB user
read -p "Input DB USER             : " DB_USER

# Input password (silent)
read -s -p "New password              : " DB_PASSWORD
echo

# Konfirmasi password
ok=true
while $ok; do
    read -s -p "Confirm password          : " CONFIRM_PASSWORD
    echo
    if [ "$DB_PASSWORD" != "$CONFIRM_PASSWORD" ]; then
        echo "Password not match, please try again."
    else
        ok=false
    fi
done

# Input host
read -p "Input host [default: localhost] : " DB_HOST
DB_HOST=${DB_HOST:-localhost}

# Input DB name
read -p "New database name         : " DB_NAME

# Input root user
read -p "Root user [default: root] : " ROOT_USER
ROOT_USER=${ROOT_USER:-root}

# Input root password (silent)
read -s -p "Root password             : " ROOT_PASSWORD
echo

# Tentukan perintah mysql, pakai sudo
if [ -z "$ROOT_PASSWORD" ]; then
    MARIADB_COMMAND="mysql -u $ROOT_USER"
else
    MARIADB_COMMAND="mysql -u $ROOT_USER --password=\"$ROOT_PASSWORD\""
fi

# SQL Commands
CREATE_USER_COMMAND="CREATE USER IF NOT EXISTS '$DB_USER'@'$DB_HOST' IDENTIFIED BY '$DB_PASSWORD';"
CREATE_NEW_DB="CREATE DATABASE IF NOT EXISTS $DB_NAME;"
GRANT_PRIVILEGES_COMMAND="GRANT ALL PRIVILEGES ON \`$DB_NAME\`.* TO '$DB_USER'@'$DB_HOST';"
FLUSH_PRIVILEGES_COMMAND="FLUSH PRIVILEGES;"

# Eksekusi semua perintah SQL lewat here-document
eval $MARIADB_COMMAND <<EOF
$CREATE_NEW_DB
$CREATE_USER_COMMAND
$GRANT_PRIVILEGES_COMMAND
$FLUSH_PRIVILEGES_COMMAND
EOF

# Output konfirmasi
echo
echo "✅ User '$DB_USER' created success and privilefes access to database '$DB_NAME' from host '$DB_HOST'."
# Simpan konfigurasi ke file (misalnya .dbconfig di direktori saat ini)
CONFIG_FILE=".dbconfig"

cat > "$CONFIG_FILE" <<EOF
DB_NAME=$DB_NAME
DB_USER=$DB_USER
DB_PASSWORD=$DB_PASSWORD
ROOT_USER=$ROOT_USER
ROOT_PASSWORD=$ROOT_PASSWORD
EOF

echo "💾 Konfigurasi disimpan ke $CONFIG_FILE"

