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

## Output Example:
```vbs
dIm rQYcRYotBATtLNAyJXDrlSaLspMYixuiHEjJKweCyfMsxZvoEdnguuLn, UfJnvKturBVaDzNiomRsmyVnAXPMXAJWIeaMgQayWvwLAFKRAeaRlPEm, nqPNyKVCmFYgJtlBxlJNFvkIrAXLVQOGYeDYwqoomKGUnLzhhaSryQxa
sUB vXdnFcILOlmwolUiPVPTwiFdQzGmxVfLuahXlNPWBVSVBxZleoywunKC
RQYCRyotBATTLnaYJxDRlSALsPMyiXUiHejJkwEcYfMsXzvOEdnguUlN = "587859/6757*91632/1104*6154-6055*9283-9169*1049-944*-941+1053*662360/5710*-2034+2080*3215-3146*630828/6372*957736/9209*431568/3888*21824/682*9779-9745*5017-4945*3077-2976*81540/755*-1628+1736*-7782+7893*5833-5789*-8495+8527*-2209+2296*1056609/9519*3814-3700*278748/2581*-740+840*57057/1729*267-233"
UfJNVKtUrbVadZnIoMRsMyvnaXPMXAJwIeamGQaYwVWlaFKraEArLPEM = Split(RQyCRYotbaTtLNaYJXdrLSaLSPMYIXuIHEJjkwecyFmsxzvoednGUULn, chr(eval(8809-8767)))
fOR EACh dzaPqAuHHeUBhiVVmBhBFjRadtyqCOyLuqXKXUUJLYCenKZnHQrlQAIT In uFjNvKtuRbVADzNIomrSmyvnAxPmXaJwiEAMgqAYWVwlAfkraEArLPEM
NQpNykvCmFYgjTlBxLJNfVkiRaxLVQOgYeDyWQOoMKguNLZHhaSRyqXa = nqpnYKvcmfYGJTlbxljNfvKIRaXlvQOGyeDyWqoomKgUnlZHHaSRyqXa & cHr(EVAL(DZapQaUhheUBHIvVmBhBfJRaDTyQCOYlUqxkXuuJlyCEnKznhQrlqaIt))
next
lyLxeYXJtRfahQqtahfvuUgHtgpVTrMJQNMHubOGKDJfQlRHIYBLBlpE
eval(eXEcUTE(NQPNykVCMfYGJTLBXlJNfvkIRaXLvQogyedywQoomKGUnlZhHASryqxA))
End Sub
vXdNFCILoLMWOLUipVPTwiFDqZGMxvfLUaHxLnpwbVSvbXzLEoyWUNkC
```
