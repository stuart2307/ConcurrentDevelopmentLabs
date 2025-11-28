# Reusable Barrier
https://github.com/stuart2307/ConcurrentDevelopmentLabs
This repository contains my attempts at the Concurrent Development labs. This folder consists of a reusable barrier implementation, without the use of a built-in barrier.

The way I handled this was by using channels. The barrier channel has no available space, so any thread that tries to insert something into the channel must wait for another thread to try and take from the channel. The same applies the other way around. Any thread looking to read from a channel must wait for there to be something in the channel. This is how the barrier is implemented.

Each thread attempts to read from the barrier channel. The exception to this is the final thread to reach the barrier, where it inserts empty structs into the channel for every thread, save for itself. This forces all threads to wait until the final thread is finished its work, where it subsequently releases all other threads.

This differs from Lab 3 in that it uses 2 barriers. For the second barrier, the count goes down for each thread to hit the barrier, and the thread that sees 0 is the one to release all other threads.
