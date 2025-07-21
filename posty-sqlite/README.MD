# Posty Notes

Make Migration
```
goose -dir=assets/migration create posts sql
```
Run Migration
```
goose -dir=assets/migrations sqlite3 app.db up
```
