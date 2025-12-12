
Introduction
============

This folder contains some adapted examples from
[Genki samples](https://github.com/firebase/genkit/tree/main/go/samples)
[Genkit](https://github.com/firebase/genkit) is an open-source
framework for building full-stack AI-powered applications, built
and used in production by Google's Firebase.


How to Run
==========

1. Install Genkit CLI.

   ```
   $ curl -sL cli.genkit.dev | bash
   ```

2. Export the Gemini API key. Replace `your_api_key` with
   your API key. You may apply for an API key at
   [Google AI Studio](https://aistudio.google.com/apikey).

   ```
   $ export GEMINI_API_KEY=your_api_key
   ```

3. For most samples, execute the following commands unless
   explicitly specified in ths sample's README.md file.

   ```
   $ go mod tidy
   $ genkit start -o -- go run main.go
   ```

   A browser will be launched showing the webpage of Genkit tools.
   Test your inputs in the browser or use the following command
   (with `flow-name` replaced by the flow name and with `your-input`
   replaced by your input) instead.

   ```
   $ genkit flow:run flow-name 'your-input'
   ```
