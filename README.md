[![Build Status](https://travis-ci.org/rguilmont/classifile.svg?branch=master)](https://travis-ci.org/rguilmont/classifile)
[![codecov](https://codecov.io/gh/rguilmont/classifile/branch/master/graph/badge.svg)](https://codecov.io/gh/rguilmont/classifile)


# Classifile

Classify your files automatically, from their name, content or metadata, with regexp-based rules.

Supported format so far :

- PDF 
- DOC
- DOCX 
- XML 
- HTML 
- RTF 
- ODT 
- Pages

## Installation

Ubuntu :

```
$ sudo apt-get install poppler-utils wv unrtf tidy
```

Then grab the latest version from release.

## Usage

First thing to do, is to write your rules to automatically organize Your file.
Configuration files can be written in YAML or JSON. This is an example of configuration file : 

```
searchoperations:
  - filename: ".(doc|odt|pdf)$"
    directory: /home/user/Downloads
    rules:
      - conditions:
        - elem: path
          matches: "(?i)(file_to_move)"
          expected: yes
        actions:
          - operation: move
            destination: ./result_test_assets1/
      - conditions:
        - elem: content
          matches: "(?i)(lorem ipsum)"
          expected: yes
        - elem: type
          matches: "(?i)pdf$"
          expected: no
        actions:
          - operation: copy
            destination: ./result_test_assets1/
```

In this example, a search operation will be recursively executed on path `/home/user/Downloads`, looking for files with name matching `.(doc|odt|pdf)$`.
If a found file match 

- all of the first rule conditions ( `path` matching `(?i)(file_to_move)` ) , then it will be moved to `./result_test_assets1/`.
- all of the second rule conditions ( `content` matching `(?i)(lorem ipsum)` and `type` not matching `(?i)pdf$` ) , then it will be copied to `./result_test_assets1/`.

Note that : 
- The first matching rule will be applied to a file. 
- If a file is found by multiple search operations, it will be only processed by the first one.

### Dry-run

If you want to test your configuration, and check the actions that would be taken, simply run `classifile` in dry-run mode : 

```
classfile --dry-run -conf ./my-rules.yaml
```
