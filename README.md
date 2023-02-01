# bbScan
automate bbScope and send it to Osmedeus

# Install:
```bash
go install github.com/xhzeem/bbScope@latest
```

# Usage
## Single:
```bash
bbscan xhzeem.json
```

## Multiple:
```bash
for FILE in $(ls); bbscan $FILE; done
```
