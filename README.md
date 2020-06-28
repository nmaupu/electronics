# Eagle

`eagle` directory contains my own eagle library I use for my projects.
Each device contains a `MOUSER_REF` which is used when generating a *Bill Of Materials* to retrieve part's information from Mouser via their API.

# BOM Maker
Bom-Maker is a tool to generate BOM and/or Mouser cart from a CSV.

## Usage

1. Generate a BOM using the `ULP` script `mybom` available [here](eagle/ulps/mybom.ulp)
![mybom export](doc/img/mybom.png =800x)

2. use `bom-maker` to generate either the bom or the mouser cart directly
```
# BOM in HTML format
$ cat /path/to/mybom.csv | ./bin/bom-maker -k <mouser_searchapi_key> generate -o html > /tmp/out.html

# Create a Mouser cart
$ cat /path/to/mybom.csv | ./bin/bom-maker -k <mouser_searchapi_key> cart -c <mouser_cart_api_key>
```

## CSV Format

CSV format is as follow:
```
Qty;Value;Device;MouserRef;Package;Parts
```

The CSV can contain more columns but the above ones are the mandatory ones.
To generate this CSV, this [ulp](eagle/ulps/mybom.ulp) can be used (see above for more information).

`MOUSER_REF` is the Mouser reference indicated in their website. Everything is based on this reference for the BOM to be generated correctly. It has to be added to every parts as an attribute !
An attribute can be added in the part itself directly in the library or added in the schematic.
