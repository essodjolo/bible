import os
import yaml

data = "../../data/kjv.yml"

with open(data, 'r') as kjv_bible:
    kjv = yaml.safe_load(kjv_bible)
    print(kjv['books']['John'][3][16])