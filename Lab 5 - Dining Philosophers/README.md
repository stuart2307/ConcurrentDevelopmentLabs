# Dining Philosophers
https://github.com/stuart2307/ConcurrentDevelopmentLabs
This repository contains the code for my attempts at the Concurrent Development Labs.
This folder contains the well-known Dining Philosophers problem, and a known solution.

Deadlocks occur in this problem when there is a circular wait. Everyone is trying to grab their fork, but it's in use by someone else. The thing to note here is that this only occurs when everyone is trying to grab a fork on the same side.

The implemented solution is for one person to grab the forks in the opposite order. Say there are 5 people, and 5 forks.


         E
      f     f
    A         D
     f       f
      B  f  C

f = fork

A grabs the right, B grabs the right, C grabs the right, D grabs the right, and E grabs the right. None of them can grab the other fork. Deadlock. With the opposite order for, say, person E, This is alleviated. A grabs right, B grabs right, C grabs right, D grabs right, E tries to grab left, but waits because it is taken. A can still grab its left fork now, because E never took it. This lets A eat, be free, which lets B eat, be free, which lets C eat, so on so forth.
