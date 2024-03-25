# enhaned data for riskctrl

## init db

```
psql   -U postgres -h localhost -p 5432 -c "create database enhanced_riskctrl"
```

```
psql   -U postgres -h localhost -p 5432 -d enhanced_riskctrl -f ./manifest/migration/enhanced_riskctrl.sql
```
