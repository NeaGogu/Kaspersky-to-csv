# Kaspersky-to-csv

A simple command line functionality to convert your Kaspersky Password Manager exported file to CSV format

## Why?

Because when you want to ditch KPM and move to another Password Manager you encounter an annoying issue - KPM does not export to a CSV format.

Every password manager supports importing your passwords from another password manager (Lastpass, Google Passwords, etc.) by feeding them a CSV file with the required information. 

___Almost every___ password manager allows you to export your valuable information in case you decide not to use them anymore... but not KPM, which exports to a more human-readable format like:

```bash
  Website name: github
  Website URL: https://github.com/NeaGogu/Kaspersky-to-csv
  Login name: aNiceName
  Login: aNiceUsername
  Password: aNicePassword
  Comment: 
```

## Usage

Clone the github project, `cd` into that directory and type the following:
```bash
./bin/ksp -src passwords.txt -des KSP_parsed.csv
```

Windows users have to use the following command:
```bash
./bin/ksp_win64.exe -src passwords.txt -des KSP_parsed.csv
```

MacOS users have to use the following command:
```bash
./bin/ksp_macos -src passwords.txt -des KSP_parsed.csv
```

The `src` tag is required, the rest of the tags are optional. For usage help you can also use
```bash
./bin/ksp --help
-------------------------------------------
OUTPUT:

Usage of ./ksp:
  -delimiter string
        Delimiter in the CSV file (default ",")
  -des string
        Destination file (default "./KSP_parsed.csv")
  -src string
        REQUIRED: Kaspersky Source file
```
