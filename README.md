# Projekat-NAISP

## Structure

All files created by the app are located in .data folder
.data is generated when app is ran

## Running App

you can run app with

```make run```
or
```go run .```

if you use make run, .data will be removed before running and created again by the app

## Tasks

There are some Todo left in code that shows where things are meant to be implemented

### LRU

- [+] Create LRU cache
- [+] Create Add and Get methods in LRU cache
  
### Merkle Tree

- [+] Implement functions for merkle tree
- [+] Implement functions for writing Metadata.txt

### SSTable

- [X] Make function that returns bytes for sstable
- [X] Make function that returns bytes for index summary
- [X] Make function that returns bytes for index
- [X] Make function that creates one or more files based on config using above functions
  
#### Naming of ssTable

  What should be naming convention for sstable and other files?
  0_1_ssTable.bin ?
  0_1_ssTableIndex.bin ?
  0_1_ssTableSummary.bin ?

### BloomFilter

- [+] Make helper functions for k and m of bloom
- [+] Make function that returns bytes of populated bloomFilter
- [+] Make function that deserializes bloomFilter
- [+] Make function that checks if element is in bloomFilter
  
### Token Bucket

- [ ] Create Token Bucket(this will be added to App->operations to Handle functions)

### WAL

- [+] Create function for reading wal file
- [+] Opening file should as memory mapped file

### bTree

- [+] Implement bTree
- [+] Make bTree implement memtable interface

### skipList

- [+] Add missing functions to skipList to implement memtable interface

### LSM

- [ ] Fix naming of sstable files to implement LSM levels

### Compaction

- [ ] Implement compaction algorithm using levels in LSM

### Config

- [ ] Create default config structure
- [ ] Adapt app to use config.yaml file if exists
- [ ] Add default config to engine->constants
- [ ] Use engine->constants if config.yaml doesn't exist

### Other

- [+] Make function for deserializing Record (in types->record.go) (need this for wal read)
- [+] Change how wal file is created

### Note

- use English for variable and function name
- use descriptive variable names
- Only expose(first letter capital) functions from packages that are used in other packages
  
## Collaboration

### Clone repository

```git clone <put url here>```

### Switch to dev branch

```git switch dev```

### Create branch

Create Branch of feature you are working on

```git checkout -b <name>```

### Push branch with

```git push --set-upstream origin <name>```

### (Optional) Open Pull request to merge to dev

You can open pull request(merge request) on github page.
You don't to do it it you are not sure how.

### Making Changes

Make changes to your code, then commit them with

```git add .```

```git commit -m "<commit message here>"```

Push changes to pull request that is already open with

```git push```

That is it :)
