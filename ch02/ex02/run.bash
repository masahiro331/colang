# go build -o cf .
cat << EOS | ./cf
10
100
1000
EOS

# go build -o cf .
./cf 10 100 1000
