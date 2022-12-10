# Module Imports
import mariadb
import sys# Connect to MariaDB Platform
try:
    conn = mariadb.connect(
        user="root",
        password="heaven",
        host="localhost",
        database="sm2k4"
        )
except mariadb.Error as e:
    print(f"Error connecting to MariaDB Platform: {e}")
    sys.exit(1)# Get Cursor
cur = conn.cursor()
cur.execute(
    "SELECT CoachName FROM club WHERE Gender='M'")

for (CoachName) in cur:
    print(f"First Name: {CoachName}")