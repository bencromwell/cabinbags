# Cabin Bags

Check owned cabin bags (or prospective cabin bags) against the various supplied airlines' bag limits.

When adding new airlines, ensure you go from largest to smallest dimension.

```shell
❯ ./dist/cabinbags 
Ryanair
  Small (40x25x20) fits within Small (40x25x20) Ryanair limit
  Small (40x25x20) fits within Large (55x40x20) Ryanair limit
  Large (50x33x20) does not fit within Small (40x25x20) Ryanair limit
  Large (50x33x20) fits within Large (55x40x20) Ryanair limit
  Test (55x40x20) does not fit within Small (40x25x20) Ryanair limit
  Test (55x40x20) fits within Large (55x40x20) Ryanair limit
EasyJet
  Small (40x25x20) fits within Small (45x36x20) EasyJet limit
  Small (40x25x20) fits within Large (56x45x25) EasyJet limit
  Large (50x33x20) does not fit within Small (45x36x20) EasyJet limit
  Large (50x33x20) fits within Large (56x45x25) EasyJet limit
  Test (55x40x20) does not fit within Small (45x36x20) EasyJet limit
  Test (55x40x20) fits within Large (56x45x25) EasyJet limit
Jet2
  Small (40x25x20) fits within Large (56x45x25) Jet2 limit
  Large (50x33x20) fits within Large (56x45x25) Jet2 limit
  Test (55x40x20) fits within Large (56x45x25) Jet2 limit
```
