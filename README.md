Nurse Scheduling Solver written in Go
=====================================

__solving the nurse scheduling problem using a genetic algorithm approach using go__

Nurse Scheduling Problem
------------------------

>The nurse scheduling problem (NSP), also called the nurse rostering 
>problem (NRP), is the operations research problem of finding an optimal 
>way to assign nurses to shifts, typically with a set of hard constraints 
>which all valid solutions must follow, and a set of soft constraints 
>which define the relative quality of valid solutions.[1] Solutions to the 
>nurse scheduling problem can be applied to constrained scheduling 
>problems in other fields
 -- https://en.wikipedia.org/wiki/Nurse_scheduling_problem

Approach
--------

This program attempts to find a monthly roster that fits the requirements of a 
department (a) and the individual wishes of the employees regarding holidays 
and such (b).

The genetic algorithm works by mutating genomes that are then ranked by 
their fitness. The higher the fitness the better the current genome matches 
the given criteria. The genomes with the highest fitness is also called elite. 
After a given amount of genetic evolutions the algorithm finishes.  


Modelling the requirements
--------------------------

__Department requirements__

For each day the department has a few requirements such as:
 * staff count working in the morning 
 * staff count working in the afternoon
 * staff count working at night

__Employee requirement__

Each employee has the following requirements
 * working hours of one employee must not exceed 40 hours
 * on holidays the employee must not be scheduled to work
 
__the fitness function__
