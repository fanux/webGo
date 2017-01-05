## Architecture
```
                                                                 ┌────────────────────┐
                                                                 │        AI          │
                                                                 └────────────────────┘
                                                                            |          
   ┌──────────────────────────────────────────────────────────────────────────────────┐
   │                                     MQ                                           │
   └──────────────────────────────────────────────────────────────────────────────────┘
          |              |               |                |                 |          
          |              |               |                |                 |          
   ┌─────────────┐ ┌─────────────┐ ┌──────────────┐ ┌────────────────┐ ┌──────────────┐                
   │    lhttp    │ │ lttp        │ │  Go worker   │ │ Go worker      │ │ Go worker    │                
   └─────────────┘ └─────────────┘ └──────────────┘ └────────────────┘ └──────────────┘                
          |              |                                                             
          |              |                                                             
   ┌─────────────┐ ┌────────────┐                                                    
   │  client1    │ │ client2    │                                                    
   └─────────────┘ └────────────┘                                                    
```
* client1 create room, publish `worker` to MQ
* One or more worker accept this work, create a chess manual and a play room
* client1,client2 and worker subscribe roomId
* client Lazi, publish to roomId
* worker receive the Lazi location, return back the resault(is the Lazi invalid, and eat chess pieces info)
* publish to client1 and client2

We can scale up the workers, and using multiple workers to ensuring play reliability.
