# CLI scraper for re-sep.

## Usage
```sh
CLI for the scraper of re-sep

Usage:
  re-sep-cli [flags] <url>

Flags:
  -h, --help         Print this help message
  -a, --all          Scrape all available articles
  -s, --single       Scrape a single article
  -o, --out string   Specify output path
      --yes          Assume yes
  -v, --verbose      Verbose output
  -w, --worker int   Number of scraping workers (default 4)
      --sleep int    Adjust worker sleep time after each job (default -1)

```

#### Push a single article into a new sqlite database
```sh
re-sep-cli -s <url> -o test.db
```
You can also run the scraper on existing database. If you do so, it will update
the records there using `ON CONFLICT DO UPDATE`

#### Scrape all articles
```sh
re-sep-cli -a <url> -o test.db
```

After that the file is your. You can do whatever you want with it.

### Examples:
#### Deploy to turso:
```sh
turso db create articles --from-file ./test.db
```

Maybe you can email the service provider your database or mail them an USB
flash drive containing the database or something. It's just a file.


## TODO
- libsql support (if necessary)
