# VBS-Obfuscator-Go
VBS-Obfuscator-GO is a Go-based tool designed for obfuscating VBScript (VBS) files. It transforms readable VBScript code into a less recognizable form by employing random variable names and encoding character values using mathematical operations. This helps protect scripts from casual inspection and modification.

# Features
- Random Variable Naming: Generates unique, random variable names to replace original readable names.
- Mathematical Character Encoding: Obfuscates character values using various arithmetic operations such as addition, subtraction, and division.
- Random Capitalization: Applies random capitalization to VBScript keywords and variable names.

## Usage
- To obfuscate a VBScript file, use the following command:
```
go run obfuscator.go inputFile.vbs outputFile.vbs
```
Replace ``inputFile.vbs`` with your source VBScript file and ``outputFile.vbs`` with the target output file path.

## Example
- To obfuscate a script named example.vbs into example_obfuscated.vbs, execute:

```go run obfuscator.go example.vbs example_obfuscated.vbs```

## Contributing
- Contributions are welcome! Please fork the repository and submit a pull request with your improvements or bug fixes.
