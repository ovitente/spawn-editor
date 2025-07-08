# Spawn files viewer
`*.swt` only.

## Reasons for creation
Reading plain text of `.swt` files is a pain for the eyes, so I created this tool to make modding a bit easier. 

It is targeting `<Action>` blocks with name `a_spawnUnitGroupToZone`

## Installation

### Prerequisites
This tool requires Go programming language to be installed on your system.

#### Linux/Mac
Install Go using your package manager

#### Windows
1. Download Go from [https://go.dev/dl/](https://go.dev/dl/)
2. Run the installer and follow the instructions
3. Open Command Prompt and verify installation:
```cmd
go version
```

### Building the binary
```bash
# Clone or download the source code
git clone <repository-url>
cd spawn-editor

# Build the binary
go build -o spawn-view main.go

# Make it executable (Linux/macOS)
chmod +x spawn-view
```

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

## Editing examples

Here is basic schema for spawning with some explanations
source
```
<Action guid="199" disabled="0"><Name>a_spawnUnitGroupToZone</Name>
  <Param>prm=?</Param>
  <Param>army</Param>
  <Param>tank</Param>
  <Param>Fnd_bradley</Param>
  <Param>brd_atgm_rotary_cannon_blocks</Param>
  <Param>cop_1_run</Param>
  <Param>player</Param>
  <Param>prm=?</Param>
  <Param>prm=?</Param>
  <Param>prm=?</Param>
  <Param>prm=?</Param>
  <Param>prm=?</Param>
</Action>
```

### Known definitions
| Parameter | Definition |
| :------ | :---------------- |
| <Action guid="199" disabled="0"><Name>a_spawnUnitGroupToZone</Name> | Name and unique ID of the action to do. ID must be unique! |
| `prm=?` | |
| `army` | |
| `tank` | |
| `Fnd_bradley` | Unit sysname |
| `brd_atgm_cannon_ceramic` | Names of loadout, taken from *upgrade_presets.xml files |
| `cop_1_run` | Spawning zone or point. Make sure it exist in your current *.swt file. Otherwise unit will not be spawned. |
| `player` | Who owns unit after spawn |
| `prm=?` | |
| `prm=?` | |
| `prm=?` | |
| `prm=?` | |
| `prm=?` | |
