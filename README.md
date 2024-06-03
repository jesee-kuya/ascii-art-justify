# ASCII Art Generator

This project is an ASCII Art Generator that converts input text into ASCII art using different fonts and optional color coding. 

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Functions](#functions)
  - [Art](#art)
  - [Ascii](#ascii)
  - [Slice](#slice)
  - [RgbToAnsiConv](#rgbtoansiconv)
  - [HexToRgb](#hextorgb)
  - [IsFlagPassed](#isflagpassed)
  - [GetFileName](#getfilename)
  - [Reader](#reader)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. Clone the repository:
    ```sh
    git clone  https://github.com/jesee-kuya/ascii-art-color.git
    ```

2. Navigate to the project directory:
    ```sh
    cd ascii-art-color
    ```

3. Build the project:
    ```sh
    go build -o ascii-art-color
    ```

## Usage

Run the binary with the required arguments:

```sh
./ascii-art-color -filename=<font_file> -color=<color_code> <letters_to_color> "Your text here"
```

### Parameters

- **-filename**: The font file to use (default: `standard`).
- **-color**: The color to use for the text.
- **-letters**: The specific letters to color.
- **"Your text here"**: The text you want to convert to ASCII art.

## Functions

### **Art**
```go
func Art(argsPassed []string, content []string, lettersTocolor string, colorCode string, indexToStartDisplay int)
```
#### **Art** 
Takes a slice of strings parsed as arguments from a certain index, joins the words, checks if there is a non-ASCII character, and calls the Ascii function to print the words in ASCII art.

### **Ascii**
```go
func Ascii(fileArr []string, wordsArr []string, lettersToColor string, colorCode string)
```

#### **Ascii**
Prints ASCII art from a given array of characters. It takes four arguments: 
- **-fileArr** (a slice of strings representing the file -array)
- **-wordsArr** (a slice of strings representing - the words to be printed), 
- **-lettersToColor** (a string -representing the letters to be colored)
- **-colorCode** (a string representing the color to be applied).

### **Slice**
```go
func Slice(word string) []string
```
- **-Slice** splits a string at \n and returns a slice of strings.

 #### **RgbToAnsiConv**
```go
func RgbToAnsiConv(colorflag string) (string, error)
```
- **-RgbToAnsiConv** converts an RGB color to an ANSI escape sequence. It takes a string representing an RGB color in the format rgb(R, G, B) and returns a string that represents the ANSI escape sequence.

#### **HexToRgb**
```go
func HexToRgb(color string) (r, g, b uint8, err error)
```
- **HexToRgb**:  converts a hex color code to RGB.  It takes a hex - **color**:  string and returns the red, green, and blue components as uint8 values.

#### **IsFlagPassed**
```go
func IsFlagPassed(name string) bool
```
- **IsFlagPassed**: checks if a flag with the given name is passed in the command line.
- **Returns**: a boolean value indicating whether the flag is passed.

#### **GetFileName**
```go
func GetFileName(name string) ([]string, error)
```
- **GetFileName**: retrieves the content of a specified ASCII art file. It takes the **Name** of the file as a parameter and **Returns** the content of the file as a slice of strings and an error if the file is not found or cannot be read.

#### **Reader**
```go
func Reader(filename string, sepp string) ([]string, error)
```
- **Reader**: checks if a flag with the given name is passed in the command line.
- **Returns**:  an error if the file is not valid.

# **Examples**
## **Example 1: Basic Usage**
```go
./ascii-art-color -filename=standard.txt "Hello, World!"
```
## **Example 2: Using Color**
```go
./ascii-art-color -filename=standard.txt -color="rgb(255, 0, 0)" "Hello, World!"
```

## **Example 3: Coloring Specific Letters**
```go
./ascii-art-color -filename=standard.txt -color="rgb(0, 255, 0)" "HW" "Hello, World!"
```

# **Contributing**
We welcome contributions to improve the ASCII Art Generator. If you find a bug or have a feature request, please open an issue. If you would like to contribute code, please fork the repository and submit a pull request.

### **Steps to Contribute**

1. **Fork the repository.**
```sh
   https://github.com/jesee-kuya/ascii-art-color.git
```
2. **Create a new branch**:
```sh
   git checkout -b feature-branch
```
3. **Make your changes.**
4. **Commit your changes**
```sh
git commit -am 'Add new feature
```
5. **Push to the branch**
```sh
git push origin feature-branch
```
5. **Open a pull request**

## License

This project is licensed under the MIT License. See the [LICENSE]( https://github.com/jesee-kuya/ascii-art-color.git/main/LICENSE)  file for details.

