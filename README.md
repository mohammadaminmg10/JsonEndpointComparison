# JSON Endpoint Comparison Tool

## Overview

This project provides tools to compare JSON responses from two different sources—either from two endpoints or two local JSON files. It is designed to assist developers and testers in identifying discrepancies in APIs or data files that are supposed to deliver the same output.

## Features

- **Dynamic Input Options**: Compare JSON data from URLs or local files.
- **Parameter Handling**: Supports custom parameters for API endpoints.
- - **Statistical Analysis**: Calculates and reports the percentage of differences, helping quantify the level of discrepancy between the compared data.
- **Detailed Comparison**: Outputs differences clearly to help identify mismatches easily.
- **Easy to Use**: Simple command-line interface for quick comparisons.

## Prerequisites

Before running the project, make sure you have the following installed:
- Go (1.15 or later)
- Access to terminal or command line interface

## Installation


Clone the repository to your local machine:

```bash
git clone https://github.com/mohammadaminmg10/JsonEndpointComparison.git
cd JsonEndpointComparison
```

## Usage

To use the JSON Endpoint Comparator, you need to specify the mode (either 'endpoints' or 'files'), and the two sources you want to compare.
- **Comparing Endpoints**:
  If you want to compare two endpoints, use the 'endpoints' mode and provide the URLs of the two endpoints. For example:
  ```bash
  go run cmd/main.go -mode=endpoints -firstURL=https://api.example.com/endpoint1 -secondURL=https://api.example.com/endpoint2
  ```

- **Comparing Files**:
- If you want to compare two local JSON files, use the 'files' mode and provide the paths to the two files. For example:
  ```bash
  go run cmd/main.go -mode=files
  ```
  
## Output
The differences between the two sources will be written to a file named result.txt in the project directory. Each difference is listed with the key and the differing values from the two sources.

## Contributing
Contributions are welcome! Please feel free to submit a Pull Request.
