## Usage

```
 $ pwd
<prj_dir>/igc-adapter

 $ go run main.go
```

## Input

```
.
├── main.go
├── input1.json <- this
├── input2.json <- this
└── input3.json <- this
```

## Output

```
.
├── .index.json <- this
├── main.go
├── input1.json
├── input2.json
└── input3.json
```

## What's .index.json

It is Descriptor for ImplementationGuide of FHIR.
