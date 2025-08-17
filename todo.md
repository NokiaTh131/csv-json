| #  | Task                            | Description                                             | Suggestion / Tips                                                                      |
| -- | ------------------------------- | ------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| 1  | **Read input CSV file**         | Open a CSV file and read all rows                       | Use `os.Open()` and `encoding/csv.NewReader`. Handle file-not-found errors gracefully. |
| 2  | **Parse CSV headers**           | Identify the first row as field names                   | Store headers in a slice of strings; ensures JSON objects have proper keys.            |
| 3  | **Convert rows to map/struct**  | Map each row to a JSON object                           | Use `map[string]string` per row, or define a dynamic struct if fixed schema.           |
| 4  | **Build JSON array**            | Aggregate all row maps into a slice                     | Example: `[]map[string]string` → makes `json.Marshal` straightforward.                 |
| 5  | **Marshal JSON**                | Convert Go slice of maps into JSON string               | Use `encoding/json.MarshalIndent` for pretty output.                                   |
| 6  | **Write to output file**        | Save JSON to user-specified output file                 | Use `os.WriteFile`. Handle errors and permissions.                                     |
| 7  | **Add CLI flags**               | Accept `-f <input.csv>` and `-o <output.json>`          | Use `flag` package. Set defaults (`output.json`) and validation for file extensions.   |
| 8  | **Optional: Handle large CSV**  | Process CSV line by line instead of reading all at once | Useful for huge CSVs to avoid memory issues. Use streaming with `csv.Reader.Read()`.   |

## Checklist

- [x] Read input CSV file
- [x] Parse CSV headers
- [x] Convert rows to map/struct
- [x] Build JSON array
- [x] Marshal JSON
- [x] Write to output file
- [x] Add CLI flags
- [ ] Optional: Handle large CSV
