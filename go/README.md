experimental support for the [feather file format](github.com/wesm/feather) in Go

the `fbs` folder was generated using `flatc`, the flatbuffers schema generator

`flatc -g metadata.fbs`

the generator will yield faulty code, because 'type' in PrimitiveArray will result in the use of a reserved keyword `type`, you need to edit it before use (here corrected to `dtype`)