# CLI scraper for re-sep.

## Usage

Push a single article into a new sqlite database
```sh
re-sep-cli -s <url> -o test.db
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
- Update existing database (when needed)
- Crawl and scrape all articles (when needed)
- libsql support (if necessary)
