# Cabin Bags

Check owned cabin bags (or prospective cabin bags) against the various supplied airlines' bag limits.

When adding new airlines, ensure you go from largest to smallest dimension.

```shell
❯ ./dist/cabinbags 
Ryanair
  Small fits within Small Ryanair limit
  Small fits within Large Ryanair limit
  Large does not fit within Small Ryanair limit
  Large fits within Large Ryanair limit
  Test does not fit within Small Ryanair limit
  Test does not fit within Large Ryanair limit
  Test does not fit within Ryanair limits
EasyJet
  Small fits within Small EasyJet limit
  Small fits within Large EasyJet limit
  Large does not fit within Small EasyJet limit
  Large fits within Large EasyJet limit
  Test does not fit within Small EasyJet limit
  Test does not fit within Large EasyJet limit
  Test does not fit within EasyJet limits
Jet2
  Small fits within Large Jet2 limit
  Large fits within Large Jet2 limit
  Test does not fit within Large Jet2 limit
  Test does not fit within Jet2 limits
```
