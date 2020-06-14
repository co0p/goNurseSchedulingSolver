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

This program attempts to find a roster that fits the requirements of a 
department (a) and tries to find a balanced shift assignment between the employees(b).

The genetic algorithm works by mutating genomes that are then ranked by 
their fitness. The higher the fitness the better the current genome matches 
the given criteria. The genomes with the highest fitness is also called elite. 
After a given amount of genetic evolutions the algorithm finishes.  

Modelling a roster as a bitset
------------------------------

- P = n employees (P1, P2, ... Pn)
- D = m days (D1, D2, ... Dm)
- S = given 3 shift types modelled as the following bits
    - S1 = Morning shift ```[0,0]```
    - S2 = Late shift   ```[0,1]```
    - S3 = Night shift  ```[1,0]```

A roster is a function of R(P, D, S) which generates a shift assignment:

     Roster           | Day 1           | Day 2           | Day 3
    +----------------------------------------------------------------------+
     employyee        |  P1   P2   P3   |  P1   P2   P3   |  P1   P2   P3
     random shift     |  S1   S2   S3   |  S2   S3   S1   |  S3   S1   S2
    +----------------------------------------------------------------------+
    resulting bitset  = [0,0][0,1][1,0]   [0,1][1,0][0,0]   [1,0][0,0][0,1]

The genetic algorithm will try to find a random shift assignment that satisfies the 
requirements.
 
Modelling the requirements
--------------------------

__Department requirements__

Every day there is a Morning, Late and Night shift. There must be at least on employee 
assigned to each shift for each day. Multiple employee assignments are allowed.

```Sum(Dn, Sm) >= 1```

Given the resulting bitset from example above, this requirement is true for all n,m:
    
    SUM(D1, S1) = 1         SUM(D2, S1) = 1         SUM(D3, S1) = 1
    SUM(D1, S2) = 1         SUM(D2, S2) = 1         SUM(D3, S2) = 1
    SUM(D1, S3) = 1         SUM(D2, S3) = 1         SUM(D3, S3) = 1
    
__Employee requirement__

For a given roster the distribution of work for the employees should be equally distributed.
It is ok of there is a slight variation (+,- tbd) regarding the distribution of work.

```SUMx(Px,Sm) ~ SUMy(Py,Sm), x,y=1..n, x!=y```
 
Given the resulting bitset from the example above, this requirement is true, because
all employees work the same amount of shifts for the roster of 3 days.
 
    m=1, SUM1(P1,S1) = 3 ~ SUM2(P2,S1) = 3 ~ SUM3(P3,S1) = 3 
    m=2, SUM1(P1,S2) = 3 ~ SUM2(P2,S2) = 3 ~ SUM3(P3,S2) = 3 
    m=3, SUM1(P1,S3) = 3 ~ SUM2(P2,S3) = 3 ~ SUM3(P3,S3) = 3 

__the fitness function__

The fitness function F(R) for a given roster R is being used to rank the chromosomes against each other. 
The higher the value, the better the chromosome. In our example we have two different kinds of rules:


F1(R): hard rule
 
- If the department's requirement are not met, the chromosome should not be considered at all ```F1(R) = 0```, otherwise ```F1(R) = 1```.
  
F2(R): soft rule

- Equal distribution of shift assignment should be rewarded, the more unbalanced the assignment, the lower the fitness value.

The final fitness function is ```F(R) = F1(R) * F2(R) ```
