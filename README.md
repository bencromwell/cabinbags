# Cabin Bags

Check owned cabin bags (or prospective cabin bags) against the various supplied airlines' bag limits.

When adding new airlines, ensure you go from largest to smallest dimension.

```shell
❯ ./dist/cabinbags 
Small (40x25x20)
  Ryanair
    - Small (40x25x20): fits
    - Large (55x40x20): fits
  EasyJet
    - Small (45x36x20): fits
    - Large (56x45x25): fits
  Jet2
    - Large (56x45x25): fits
  Lufthansa
    - Personal Item (40x30x20): fits
    - Carry-on (55x40x23): fits

Large (50x33x20)
  Ryanair
    - Small (40x25x20): does not fit
    - Large (55x40x20): fits
  EasyJet
    - Small (45x36x20): does not fit
    - Large (56x45x25): fits
  Jet2
    - Large (56x45x25): fits
  Lufthansa
    - Personal Item (40x30x20): does not fit
    - Carry-on (55x40x23): fits

Test (55x40x20)
  Ryanair
    - Small (40x25x20): does not fit
    - Large (55x40x20): fits
  EasyJet
    - Small (45x36x20): does not fit
    - Large (56x45x25): fits
  Jet2
    - Large (56x45x25): fits
  Lufthansa
    - Personal Item (40x30x20): does not fit
    - Carry-on (55x40x23): fits

https://www.rohan.co.uk/thule-subterra-ii-convertible-carry-on-40l-dark-slate/ (55x35x21)
  Ryanair
    - Small (40x25x20): does not fit
    - Large (55x40x20): does not fit
  EasyJet
    - Small (45x36x20): does not fit
    - Large (56x45x25): fits
  Jet2
    - Large (56x45x25): fits
  Lufthansa
    - Personal Item (40x30x20): does not fit
    - Carry-on (55x40x23): fits
```
