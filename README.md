# rand-csv

Generate random CSV data. For teaching a friend about parsing CSVs in Python.

## Usage
Print a csv to stdout
```bash
go run donors.go --rows 3 --min 10 --max 50
```
```csv
Emma Bailey,17
Megan Watson,29
Emma Bailey,48
```

Print csv to file with dates
```bash
go run donors.go --rows 10000 --min 3 --max 1000 --date --output donor_list.csv
```

The file `donor_list.csv` will contain some fake data
```
2020-05-02,Joe Wilson,128
2020-05-03,Emma Bailey,458
2020-04-17,Hannah Langdon,60
2020-04-23,Wendy Burgess,122
2020-04-28,Stephen Quinn,968
2020-04-24,Nathan Ross,423
2020-04-25,Diana Hughes,457
2020-04-16,Adrian Campbell,538
2020-05-03,Megan Watson,293
2020-03-29,Elizabeth Young,103
```
