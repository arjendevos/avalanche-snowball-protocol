# Simple implementation of the Snowball Protocol

This is not a 100% accurate implementation but more a playing around implementation.

Nodes should query at the same time, which doesn't happen here.
This result in the first nodes doing all the work.

The real implementation looks more like this:

1. Nodes receive transactions
2. Nodes send queries (20) to other nodes with their transaction preference
3. Nodes send back their transaction preference
4. ````
   If 14 nodes have same preference:
       if preference == current node preference:
           success++
   else:
       current node preference = preference
       success = 1
   else:
       success = 0```

   ````

5. if success reaches a threshold (20), then we got a preference
