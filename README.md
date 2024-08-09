# Executable Byte Converter

## Overview

**Executable Byte Converter** is a Go-based tool that reads an executable (.exe) file and converts its bytes into numeric values. The resulting numeric values are then saved to a text file in the same directory as the converter executable, making it easy to analyze or inspect the binary data.

### Developed by

- **real7lab**
- [Discord Server](https://discord.gg/visionn)

## Features

- Converts each byte of an executable file into its corresponding numeric value.
- Automatically generates a text file with the byte values, using the same base name as the original executable.
- Simple and intuitive command-line interface.

## Installation

### Prerequisites

- Go (version 1.16 or later)

### Steps

1. **Clone the Repository**

   ```bash
   git clone https://github.com/real7lab/Executable-Byte-Converter.git
   cd executable-byte-converter
   ```

2. **Build the Program**

   ```bash
   go build -o exe-byte-converter
   ```

3. **Run the Program**

   ```bash
   ./exe-byte-converter
   ```

## Configuration

### `exe-byte-converter`

The tool does not require any additional configuration. Simply run the executable and provide the path to the .exe file when prompted.

### Usage

1. **Launch the Program**: Run `./exe-byte-converter` in your terminal or command prompt.
2. **Enter the File Path**: When prompted, input the full path to the .exe file you want to convert.
3. **View the Output**: The program will create a `.txt` file with the numeric byte values in the same directory where the converter is executed.

## Error Handling

Errors encountered during file reading or writing are silently ignored. Ensure that the provided file path is correct and that you have the necessary permissions to read and write files in the directories involved.

## Troubleshooting

- **Issue: No output file is created**: Verify that the path to the executable file is correct and that the file exists.
- **Note:** Ensure that you have the necessary permissions to read the .exe file and write the output in the specified directory.

## License


This project is developed by [real7lab](https://discord.gg/visionn).

For more details, please refer to our [official Discord server](https://discord.gg/visionn).
