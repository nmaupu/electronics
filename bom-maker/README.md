Bom-Maker is a tool to generate BOM from a CSV.

# CSV Format

CSV format is as follow:
```
Qty;Value;Device;MouserRef;Package;Parts
```

The CSV can contain more columns but the above ones are the mandatory ones.
To generate this CSV, this [ulp](../eagle/ulps/mybom.ulp) can be used using the following options:

![mybom export](doc/img/mybom.png =800x)

`MOUSER_REF` is the Mouser reference indicated in their website. Everything is based on this reference for the BOM to be generated correctly. It has to be added to every parts as an attribute !
An attribute can be added in the part itself directly in the library or added in the schematic.
