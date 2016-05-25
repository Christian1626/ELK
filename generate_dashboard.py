import json
import re
from pprint import pprint
from collections import OrderedDict

with open('data.json') as data_file:    
    data = json.load(data_file)



# s = "Example String logstash-canalv2-* adasd"
# data = re.sub('logstash-canalv2-\*', 'logstash-canalv2-mg', data)
# data = re.sub('\', \'\_type\'', '(mg)\', \'_type\'', str(data))
# data = re.sub('\",\\\\n    \"panelIndex\"','(mg)\",\\\\n    \"panelIndex\"', data)
# data = re.sub('\"id\": \".*\"', '\"id\": \".*(mg)\"', str(data))

print(data) 