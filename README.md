gscan
=====

a port scan with go program language


    Usage of ./gscan:
      -h="help": help doc
      -ip="127.0.0.1": IP range of scan. Example:
      	192.168.1.1
    		192.168.1.1, 192.168.1.5
    		192.168.1.1-192.168.1.100
      -p="1-1024": Port range of scan. Example:
    		135
    		135, 445, 3389
    		1-1024
      -w="connect": Way of scan. Expample
    		connect
    		syn
    		fin
