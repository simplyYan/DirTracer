# DirTracer

**DirTracer** is a command-line utility written in Go that generates a visual representation of a directory tree. It allows users to easily view the structure of their directories and customize the output according to their preferences.

## Features

- Generates a tree-like structure of directories and files.
- Includes options to include or exclude files from the output.
- Works on Windows, Linux, and macOS.

## Installation

### Prerequisites

- Ensure you have [Go](https://golang.org/dl/) installed on your machine.
- A terminal or command prompt for running commands.

### Steps to Install

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/simplyYan/DirTracer.git
   cd DirTracer
   ```

2. **Build the Executable:**

   Choose the appropriate command based on your operating system:

   **For Windows:**

   ```bash
   GOOS=windows GOARCH=amd64 go build -o DirTracer.exe
   ```

   **For Linux:**

   ```bash
   GOOS=linux GOARCH=amd64 go build -o DirTracer
   ```

   **For macOS:**

   ```bash
   GOOS=darwin GOARCH=amd64 go build -o DirTracer
   ```

3. **Add to PATH (Optional):**

   To make the `DirTracer` command accessible from anywhere in your terminal, move the executable to a directory that is included in your `PATH`. For example:

   - **Linux/macOS:**
     ```bash
     mv DirTracer /usr/local/bin/
     ```

   - **Windows:**
     Move `DirTracer.exe` to a directory listed in your `PATH`, or add the directory containing `DirTracer.exe` to your system `PATH`.

## Usage

To use **DirTracer**, open your terminal and run the following command:

```bash
./DirTracer -dir "/path/to/your_directory" -files true
```

### Command-Line Options

- `-dir`: Specify the directory to trace. The default is the current directory (`.`).
  
- `-files`: Include files in the output. The default is `true`. Set to `false` to exclude files from the output.

### Example

```bash
./DirTracer -dir "/home/user/Documents" -files true
```

This command will display the directory structure of `Documents`, including all files.

## Output Format

The output will be structured in a tree format, similar to:

```
/my_directory
├── documents
│   ├── letter.txt
│   ├── resume.pdf
│   └── grades/
│       ├── grade1.txt
│       └── grade2.txt
├── images
│   ├── photo1.jpg
│   ├── photo2.png
│   └── wallpapers/
│       ├── wallpaper1.jpg
│       └── wallpaper2.png
└── music
    ├── song1.mp3
    ├── song2.mp3
    └── playlists/
        ├── playlist1.m3u
        └── playlist2.m3u
```

## Contribution

Contributions are welcome! If you have suggestions or improvements, feel free to submit a pull request or open an issue.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to the Go community for providing excellent documentation and support.

## Contact

For any questions or feedback, please reach out to [your-email@example.com].
