# HVPPS share reward for various a and X values.

## HVPPS description
HVPPS (High-Variance Pay Per Share) aka PoT (Pay on Target) is a configurable superset of PPS system which allows you to customize miner's variance. By playing with ```a``` and ```X``` variables you can literally turn it into either plain PPS or something moderate between PPS and solo mining.

## Examples

### Common variables

shareDiff = 4294967296

netDiff = 3.211 P

fee = 1.5%

### Moderate variance configuration
a=0.800000 X=5.000000
```
5G share reward: 0.00000092873202164 WSH
20 G share reward: 0.00000281538902210 WSH
80 G share reward: 0.00000853466356393 WSH
320 G share reward: 0.00002587226190688 WSH
1280 G share reward: 0.00007843003197071 WSH
5120 G share reward: 0.00023775539753990 WSH
NetDiff/4 share reward: 0.01356371537117023 WSH
NetDiff/3 share reward: 0.01707377905157401 WSH
NetDiff/2 share reward: 0.02361580011352108 WSH
Maximum share reward: 0.14900562527402308 WSH
Block solving share reward: 0.04111749618302889 WSH
```
### Big variance configuration
a=0.900000 X=4.000000
```
5G share reward: 0.00000056835189606 WSH
20 G share reward: 0.00000197911625307 WSH
80 G share reward: 0.00000689168307577 WSH
320 G share reward: 0.00002399823433467 WSH
1280 G share reward: 0.00008356670567262 WSH
5120 G share reward: 0.00029099617078442 WSH
NetDiff/4 share reward: 0.02752115093104003 WSH
NetDiff/3 share reward: 0.03565426228924745 WSH
NetDiff/2 share reward: 0.05135628356744854 WSH
Maximum share reward: 0.33371411516488336 WSH
Block solving share reward: 0.09583421378229816 WSH
```

## References

https://bitcointalk.org/index.php?topic=131376.0 - Original thread on bitcointalk.org

https://play.golang.org/p/boRVu--wlgH - Script used to calculate rewards in this note
