# Term Helper CLI

`term-helper` is a CLI application built using [Cobra](https://github.com/spf13/cobra) in Go that helps users write terminal commands. It utilizes the Groq API to assist with Linux terminal commands by interpreting user prompts.

## Features

- Provides assistance in writing Linux terminal commands based on user input.
- Interacts with the Groq API using a pre-trained language model (`llama3-8b-8192`).
- Allows users to input prompts as flags or arguments.
- Provides a detailed response from the API, including tokens and message content.

## Prerequisites

To use this CLI, you'll need:

- Go (1.19 or later) installed on your system.
- A valid API key from Groq. Set the environment variable `GROQ_API_KEY` to authenticate API requests.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/basola21/term-helper.git
   cd term-helper
   ```

2. Build the project:

   ```bash
   go build -o term-helper
   ```

3. Export your Groq API key:

   ```bash
   export GROQ_API_KEY=your_api_key_here
   ```

4. Run the command-line tool:

   ```bash
   ./term-helper --prompt "Show me how to list all files"
   ```

## Usage

You can run `term-helper` in two ways:

### Using a Prompt Argument

Run the application by providing a prompt as an argument:

```bash
./term-helper "How do I list all files in a directory?"
```

### Using a Prompt Flag

You can also pass the prompt using a flag:

```bash
./term-helper --prompt "How do I change file permissions in Linux?"
```

### Available Commands and Options

- `-p`, `--prompt`: Provide a prompt to query the Groq API.

### Example

```bash
./term-helper --prompt "How do I create a new directory in Linux?"
```

### Output

The application will return the response from the Groq API, including information about the message and tokens:

```bash
Response from Groq:
ID: 12345
Model: llama3-8b-8192
Message: Use the command `mkdir` followed by the directory name.
```

## Environment Variables

Make sure to set your Groq API key in the environment:

```bash
export GROQ_API_KEY=your_api_key_here
```

This is necessary for the application to authenticate with the Groq API.

## Error Handling

- If no prompt is provided, the application will prompt you to input one.
- If the API key is not set, an error message will ask you to set the `GROQ_API_KEY` environment variable.
- All HTTP errors or issues with the API response will be printed to the console.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Feel free to open issues or submit pull requests to improve this project. Contributions are welcome!

---

Enjoy using `term-helper` to get quick assistance with Linux terminal commands!
