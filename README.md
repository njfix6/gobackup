# gobackup

Gobackup is a file backup command line that saves jobs easily for reuse.

### Requirements
1. [Go](https://golang.org/)

### Install
1. Install [Go](https://golang.org/)
2. Run: `go get github.com/njfix6/gobackup` to install command line.

### Workflow
##### First Backup
Run: `gobackup <job> <folder1> <folder2>` . This will sync folder1 to folder2 and save the source and destination under the job's name. The config is located at `~/.gobackup/config.json`
##### Repeating backup
Run: `gobackup <job>` with the job name made earier to backup those folders over and over again.
