# Tools
This sub-repo is for tools related to election, staking and voting.

# Dumper
Dumper dumps the information from staking contract (0x87c9dbff0016af23f5b1ab9b8e072124ab729193) and registration contract (0x95724986563028deb58f15c5fac19fa09304f32d) to a local csv file.

# Processor
Processor processes the csv produced by Dumper and generate another csv that breaks down the votes to each voters for each delegate. See the samples for more information.

# Generate breakdown of votes
```
go run tools/dumper/dumper.go > tools/processor/stats.csv
go run tools/processor/processor.go
```
# How to use the produced information

For example, one BP's breakdown looks like below:
```
>>>blockboost<<<
io1yfyjftv050x3pc6ee33k3eg23ldzm8pn229972    1491626522508296126545702
io1ge3ursda36ept6vqlq3g3kknd92rfuw2ecpehk    7531000000000000000000
io1t09saa20x8n89rn0zjusjnh4pw2lemlhty3lnc    4000000000000000000000000
io1cchxj9ffwm75teycynjmuknffvwdmc5sk94pn5    2000000000000000000000000
io1cse9e458vyqea0vdmlz0kzed8mv2xplme5qmes    860828544270782813896403
```
which translates to
```
io1yfyjftv050x3pc6ee33k3eg23ldzm8pn229972    17.84%
io1ge3ursda36ept6vqlq3g3kknd92rfuw2ecpehk    0.09%
io1t09saa20x8n89rn0zjusjnh4pw2lemlhty3lnc    47.84%
io1cchxj9ffwm75teycynjmuknffvwdmc5sk94pn5    23.92%
io1cse9e458vyqea0vdmlz0kzed8mv2xplme5qmes    10.29%
```

If this bp blockboost earned 125 IOTX in this epoch, and would distribute 80% (of 125 IOTX) to his/her voters. He/she will send 17.84, 0.09, 47.84, 23.92 and 10.29 IOTX to the five io addresses respectively.
