# Spawn files viewer
`*.swt` only.

## Reasons for creation
Reading plane text of `.swt` files is a pain for the eyes, so i created this tool is created to make modding a bit easier. 
Tool is targeting Action blocks with name `a_spawnUnitGroupToZone`

## Usage
`spawn-view <filename>`

```
$ spawn-view 3rd_mission.swt

LINE | TRIGGER           | WAVE             | UNIT TO SPAWN       | SPAWN POINT    | APPLIED OWNER
----------------------------------------------------------------------------------------------------
1696 | b1_armed          | guards           | Lgn_homunculus      | b1             | integrators
1710 | b1_armed          | guards           | Lgn_homunculus      | b1             | integrators
1724 | b1_armed          | guards           | Lgn_homunculus      | b2             | integrators
1738 | b1_armed          | guards           | Lgn_spider          | b3             | integrators
1752 | b1_armed          | guards           | Lgn_wolfpack        | b4             | integrators
1774 | b2_armed          | guards           | Lgn_homunculus      | b2             | integrators
1894 | b3_armed          | guards           | Lgn_spider          | b4             | integrators
1908 | b3_armed          | guards           | Lgn_wolfpack        | b2             | integrators
1930 | b4_armed          | guards           | Lgn_homunculus      | b3             | integrators
2234 | boss_b2           | prm=?            | Madrobot_damaged    | b2_spwn        | founders
2338 | boss_b3           | prm=?            | Madrobot_damaged    | b3_spwn        | founders
2442 | boss_b4           | prm=?            | Madrobot_damaged    | b4_spwn        | founders
2544 | guard_aggr_b1     | prm=?            | Madrobot_damaged    | b1_spwn        | founders
2646 | guard_aggr_b2     | prm=?            | Madrobot_damaged    | b2_spwn        | founders
2748 | guard_aggr_b3     | prm=?            | Madrobot_damaged    | b3_spwn        | founders
2850 | guard_aggr_b4     | prm=?            | Madrobot_damaged    | b4_spwn        | founders
2954 | legion_spawn1     | legion_group     | Lgn_homunculus      | legion_spawn1  | total_marauders
2968 | legion_spawn1     | legion_group     | Lgn_wolfpack        | legion_spawn1  | total_marauders
3009 | legion_spawn2     | legion_group     | Lgn_homunculus      | legion_spawn2  | total_marauders
3023 | legion_spawn2     | legion_group     | Lgn_wolfpack        | legion_spawn2  | total_marauders
...
```
## A bit of explanations

| COLUMN NAME       | MEANING |
| :---------------- | :------ |
| LINE              | Related to action name `a_spawnUnitGroupToZone` string number |
| TRIGGER           | Self explanatory |
| WAVE              | The sequence in one trigger to spawn units |
| UNIT TO SPAWN     | The actual unit `sysname` to spawn it |
| SPAWN POINT       | Area where unit will be spawned |
| APPLIED OWNER     | Faction on which unit will appear |
