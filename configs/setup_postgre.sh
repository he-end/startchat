#!/bin/bash

echo "📦 Installing PostgreSQL..."
# sudo apt update -y
# sudo apt install -y postgresql postgresql-contrib

# ==============================
# 🔧 Input Credentials
# ==============================
read -p "Input DB USER             : " DB_USER
read -s -p "New password              : " DB_PASSWORD
echo
read -s -p "Confirm password          : " CONFIRM_PASSWORD
echo

while [ "$DB_PASSWORD" != "$CONFIRM_PASSWORD" ]; do
    echo "❌ Password not match, try again."
    read -s -p "New password              : " DB_PASSWORD
    echo
    read -s -p "Confirm password          : " CONFIRM_PASSWORD
    echo
done

read -p "New database name         : " DB_NAME
read -p "PostgreSQL superuser [default: postgres] : " ROOT_USER
ROOT_USER=${ROOT_USER:-postgres}

read -s -p "Set password for superuser (optional)     : " ROOT_PASSWORD
echo

# ==============================
# 🔐 Buat user & DB
# ==============================
echo "🔐 Creating user and database..."

sudo -u "$ROOT_USER" psql <<EOF
DO \$\$
BEGIN
   IF NOT EXISTS (
      SELECT FROM pg_catalog.pg_roles WHERE rolname = '$DB_USER'
   ) THEN
      CREATE ROLE "$DB_USER" WITH LOGIN PASSWORD '$DB_PASSWORD';
   END IF;
END
\$\$;

CREATE DATABASE "$DB_NAME" OWNER "$DB_USER";
GRANT ALL PRIVILEGES ON DATABASE "$DB_NAME" TO "$DB_USER";
EOF

# ==============================
# 🔐 Set password untuk superuser jika diisi
# ==============================
if [ -n "$ROOT_PASSWORD" ]; then
    echo "🔒 Setting password for superuser $ROOT_USER..."
    sudo -u "$ROOT_USER" psql -c "ALTER USER $ROOT_USER WITH PASSWORD '$ROOT_PASSWORD';"
fi

# ==============================
# ⚙️ Setup konfigurasi PostgreSQL
# ==============================
PG_VERSION=$(psql -V | awk '{print $3}' | cut -d. -f1)
PG_CONF_DIR="/etc/postgresql/$PG_VERSION/main"
HBA_FILE="$PG_CONF_DIR/pg_hba.conf"
CONF_FILE="$PG_CONF_DIR/postgresql.conf"

echo "⚙️ Configuring PostgreSQL to allow password login..."

# 1. Ubah pg_hba.conf jika file ada
if [ -f "$HBA_FILE" ]; then
    sudo sed -i 's/^\s*local\s\+all\s\+all\s\+peer\s*$/local all all md5/' "$HBA_FILE"
    sudo sed -i 's/^\s*host\s\+all\s\+all\s\+127\.0\.0\.1\/32\s\+.*/host all all 127.0.0.1\/32 md5/' "$HBA_FILE"
    sudo sed -i 's/^\s*host\s\+all\s\+all\s\+::1\/128\s\+.*/host all all ::1\/128 md5/' "$HBA_FILE"
else
    echo "❗ File $HBA_FILE not found. Skipping pg_hba.conf update."
fi

# 2. Ubah postgresql.conf jika ada
if [ -f "$CONF_FILE" ]; then
    sudo sed -i "s/^#\?listen_addresses\s*=.*/listen_addresses = '*'/" "$CONF_FILE"
else
    echo "❗ File $CONF_FILE not found. Skipping postgresql.conf update."
fi

# ==============================
# 🔄 Restart PostgreSQL
# ==============================
echo "🔁 Restarting PostgreSQL service..."
sudo systemctl restart postgresql

# ==============================
# 💾 Save to .dbconfig
# ==============================
CONFIG_FILE=".dbconfig"
cat > "$CONFIG_FILE" <<EOF
DB_NAME=$DB_NAME
DB_USER=$DB_USER
DB_PASSWORD=$DB_PASSWORD
HOST=127.0.0.1
PORT=5432
EOF

chmod 600 "$CONFIG_FILE"

echo "✅ PostgreSQL setup complete!"
echo "💾 Configuration saved to $CONFIG_FILE"
