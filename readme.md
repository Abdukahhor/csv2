csv2 reads CSV data, parses the CSV-encoded data and stores the result in the struct slice. 
    - Validation
    - Write JSON/XML format

Package csv reads and writes comma-separated values (CSV) files. There are many kinds of CSV files; this package supports the format described in RFC 4180.

A csv file contains zero or more records of one or more fields per record. Each record is separated by the newline character. The final record may optionally be followed by a newline character.
