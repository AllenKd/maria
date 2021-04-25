# Maria

CLI tool for some routine job

## Build

```
$ make build
```

## Commands

```
$ ./maria --help
  NAME:
     Maria - A CLI for routine job
  
  USAGE:
     maria [global options] command [command options] [arguments...]
  
  VERSION:
     0.1.0
  
  COMMANDS:
     clean, c  Clean folder
     help, h   Shows a list of commands or help for one command
  
  GLOBAL OPTIONS:
     --help, -h     show help (default: false)
     --version, -v  print the version (default: false)
```

### Clean

Clean misc files and folders except the main(largest) file for each folder, root files are keep the same.

#### Example

Test folder before clean: 

```
$ tree --du -h ./testFolder
  ./testFolder
  ├── [143K]  fakeDir1
  │   └── [143K]  largestFileInDir1.mp3 <--- will be moveout
  ├── [4.5M]  fakeDir2
  │   ├── [3.0M]  fakeSubDir2
  │   │   ├── [2.0M]  largestFileInDir2.jpg <--- will be moveout
  │   │   └── [1.0M]  smallFileInSubDir2.jpg
  │   ├── [1.4M]  largerFileInDir2.mp3
  │   └── [200K]  smallFileInDir2.mp3
  ├── [1.8M]  fakeDir3
  │   ├── [977K]  largestFileInDir3.jpg <--- will be moveout
  │   └── [879K]  smallFileInDir3.jpg
  ├── [784K]  rootFile1.jpg
  ├── [800K]  rootFile2.mp3
  └── [400K]  rootFile3.mp3
```

Test folder after clean:

```
$ tree --du -h ./testFolder
  ./testFolder
  ├── [143K]  largestFileInDir1.mp3
  ├── [2.0M]  largestFileInDir2.jpg
  ├── [977K]  largestFileInDir3.jpg
  ├── [784K]  rootFile1.jpg
  ├── [800K]  rootFile2.mp3
  └── [400K]  rootFile3.mp3
```

#### Usage

```
$ ./maria clean --help
  NAME:
     maria clean - Clean folder
  
  USAGE:
     maria clean [command options] [arguments...]
  
  OPTIONS:
     --root-folder value, -r value  root folder path (default: "./testFolder")
     --help, -h                     show help (default: false)
```