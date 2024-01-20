# NonoGo
This is a fun project which I have mainly used to teach myself Go programming.

## Puzzle Files
The puzzle files in the `data` folder are from: [https://github.com/mikix/nonogram-db](https://github.com/mikix/nonogram-db)
Each puzzle file is a YAML file which contains the puzzle definition. The required format should be clear from an examination of any one of the puzzle files in the `data` directory.

## Example Usage
Compile the code, and then run the executable, pointing it to a puzzle file.  See the `data` directory for examples.

```bash
nonogo --puzzle=./data/candle.nonogram.yaml --debug=true|false
```