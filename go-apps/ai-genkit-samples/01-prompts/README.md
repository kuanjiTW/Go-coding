
Introduction
============

This example demonstrates a simple prompt request across multiple
LLM backends.


How to Run
==========

1. Download and install [ollama](https://ollama.com).

2. Install models in ollama.

   ```
   $ ollama run gpt-oss
   $ ollama llama3
   ```

3. Install Go dependencies.

   ```
   $ go mod tidy
   ```

4. Run the program.

   ```
   $ go run main.go
   ```
