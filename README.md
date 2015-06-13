(;,;)
=======

qthulhu
-------
- [ ] Append data to partition
- [ ] Handle partition
- [ ] Generate sequence


Design:
Each partition has a db
Track the last key and increment?

qthulhu



append(partition, value)

  Store:
   - open(partition)
   - generate_key
   - put(key, value)
