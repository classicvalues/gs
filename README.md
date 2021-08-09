# Go-Gulfstream project generator
Standard Tooling for Go-Gulfstream Development

## Table of contents
- [Create a new manifest file](#1-create-a-new-manifest-file-for-project)
- [Initialize a new project](#2-initialize-a-new-project)
- [Add mutations](#3-add-mutations)
- [Apply changes](#4-apply-changes-to-the-project)
- [Example manifest file](docs/manifest.md)

### 1. Create a new manifest file for project
With empty manifest file
```shell script
$ gs manifest path/to/project
```

With data example
```shell script
$ gs manifest -d path/to/project 
```

### 2. Initialize a new project
```shell script
$ gs init path/to/project
```

### 3. Add mutations
Update manifest file
```shell script
$ gs add path/to/project
$ gs apply path/to/project
```
OR short entry
```shell script
$ gs add -a path/to/project 
```

### 4. Apply changes to the project
Edit the ```path/to/project/gulfstream.yml``` manifest file and then execute the command:
```shell script
$ gs apply path/to/project  
```

With interactive mode
```shell script
$ gs apply -i path/to/project  
```