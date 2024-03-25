# enhaned data for riskctrl

## init db

```
psql   -U postgres -h localhost -p 5432 -c "create database enhanced-riskctr"
```

```
psql   -U postgres -h localhost -p 5432 -d enhanced-riskctr -f ./manifest/migration/enhanced-riskctr.sql
```
