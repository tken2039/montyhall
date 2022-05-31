# BUILD

```
❯ make build
```

# SAMPLE 1

```
❯ ./bin/montyhall play --try-count 373981 --detail

- Preferences -
Number of doors: 3
Number of trials: 373981
Changes after Monty opened the door: yes
---------------

- Work Detail -
Start verification.
Number of workers: 4
HitRate (worker[4]): 49393 / 73981
HitRate (worker[3]): 66574 / 100000
HitRate (worker[1]): 66864 / 100000
HitRate (worker[2]): 66800 / 100000
---------------

- Result -
hit count: 249631
hit rate: 66.75 %
----------
```

# SAMPLE 2

```
❯ ./bin/montyhall play --doors 4 --try-count 1601 --change=false

- Preferences -
Number of doors: 4
Number of trials: 1601
Changes after Monty opened the door: no
---------------

- Result -
hit count: 396
hit rate: 24.73 %
----------
```