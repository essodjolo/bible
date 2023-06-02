#!/bin/bash

sed -i '' 's/\([[:digit:]]*:[[:digit:]]*\)/ \1/' data/kjv.txt

sed -i '' 's/^\([[:digit:]]\)\([[:alpha:]]*\)/\1 \2/' data/kjv.txt

sed -i '' 's/^\([[:alpha:]]*\) \([[:digit:]]*:[[:digit:]]*\) \(.*\)$/\1 \2 "\3"/' data/kjv.txt

sed -i '' 's/^\([[:digit:]] [[:alpha:]]* [[:digit:]]*:[[:digit:]]*\) \(.*\)$/\1 "\2"/' data/kjv.txt