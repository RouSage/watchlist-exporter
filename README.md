# IMDb Watchlist Exporter

This is a simple script that will export your IMDb watchlist (in a CSV format) to a Notion database.

## Usage

1. Go to [IMDb](https://www.imdb.com/) and log in.
2. Go to the Watchlist page.

   a. Click on the "Export" button
   b. You can also export specific lists from the Your lists

3. IMDb will give you a link to download the CSV file.

   a. A CSV file should be in the following format:

   ```csv
   Position,Const,Created,Modified,Description,Title,Original Title,URL,Title Type,IMDb Rating,Runtime (mins),Year,Genres,Num Votes,Release Date,Directors,Your Rating,Date Rated
   ```

4. Create integration in Notion and connect it to a page (you can access Notion API only from a specific page).

   a. See more details [here](https://developers.notion.com/docs/create-a-notion-integration)

5. Copy `.env.example` to `.env` and fill in the values.

   a. You should have a Notion API key and a Notion page ID after creating the integration.

6. Install dependencies with `make deps`.
7. Now you can use the `go run main.go` command to export your watchlist to Notion.

   a. Use the `--path` flag to specify the path to the CSV file.

   b. Use the `--database-name` flag to specify the name of the database to create. If you don't provide this flag, the script will create a new database with the name `Watchlist DB`.

   c. Use the `--database-id` flag to specify the ID of the existing database. If you don't provide this flag, the script will create a new database. This also ignores the `--database-name` flag. Database ID is 32 characters long and can be found in the URL of the database in Notion.

## Example

```bash
# will export IMDb CSV to Notion's database with ID <database-id>
go run main.go --path ./data/watchlist.csv --database-id <database-id>

# will export IMDb CSV to a NEW Notion database with name "My Watchlist"
go run main.go --path ./data/watchlist.csv --database-name "My Watchlist"
```

Example of a resulting Notion database:

<img width="1111" alt="image" src="https://github.com/user-attachments/assets/29dd2d90-457c-4995-9949-b16a65b6fabc">


## Useful commands

- `make build` - build the binary
- `make test` - run tests
- `make clean` - remove the binary
- `make deps` - install dependencies
