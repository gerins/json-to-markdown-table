### Main Structure
| Parameter  | Required | Data Type | Description      | Example         |
|------------|----------|-----------|------------------|-----------------|
| quiz        | M        | Object      | Field for 'quiz'              | Refer to sub-structure           |
### quiz Structure
| Parameter  | Required | Data Type | Description      | Example         |
|------------|----------|-----------|------------------|-----------------|
| sport        | M        | Object      | Field for 'sport'              | Refer to sub-structure           |
| maths        | M        | Object      | Field for 'maths'              | Refer to sub-structure           |
### sport Structure
| Parameter  | Required | Data Type | Description      | Example         |
|------------|----------|-----------|------------------|-----------------|
| q1        | M        | Object      | Field for 'q1'              | Refer to sub-structure           |
### q1 Structure
| Parameter  | Required | Data Type | Description      | Example         |
|------------|----------|-----------|------------------|-----------------|
| question        | M        | String      | Field for 'question'              | Which one is correct team name in NBA?           |
| options        | M        | Array String      | Field for 'options'              | Refer to list           |
| answer        | M        | String      | Field for 'answer'              | Huston Rocket           |


### maths Structure
| Parameter  | Required | Data Type | Description      | Example         |
|------------|----------|-----------|------------------|-----------------|
| q1        | M        | Object      | Field for 'q1'              | Refer to sub-structure           |
| q2        | M        | Object      | Field for 'q2'              | Refer to sub-structure           |
### q1 Structure
| Parameter  | Required | Data Type | Description      | Example         |
|------------|----------|-----------|------------------|-----------------|
| question        | M        | String      | Field for 'question'              | 5 + 7 = ?           |
| options        | M        | Array String      | Field for 'options'              | Refer to list           |
| answer        | M        | String      | Field for 'answer'              | 12           |

### q2 Structure
| Parameter  | Required | Data Type | Description      | Example         |
|------------|----------|-----------|------------------|-----------------|
| question        | M        | String      | Field for 'question'              | 12 - 8 = ?           |
| options        | M        | Array String      | Field for 'options'              | Refer to list           |
| answer        | M        | String      | Field for 'answer'              | 4           |