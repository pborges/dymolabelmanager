# Dymo LabelManager PNP
## Do not rely on this library as it is just what I use to print barcodes for my home inventory system

### usage 
dymobarcode <pdf|print> <text>
### ex

```
dymobarcode pdf "blargh"
```

will create a barcode PDF named blargh.pdf


```
dymobarcode print "blargh"
```

will create a barcode in the temporary directory then pipe it to lpr
