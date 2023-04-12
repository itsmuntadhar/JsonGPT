# JsonGPT

JsonGPT is a simple library for generating JSON from a template.

You can use it as a cli tool to generate once or run a local server to generate on demand.

## Why?

Well, starting from "why spend 10 minutes to do something while you can spend a couple of hours to automate it?". I was working on a project and starting with the frontend and wanted to test and generate JSON. So I used GPT to describe to it the shape I wanted and it generates it for me.

Also an obligatory *why not?*

## Installation

For now, you can just clone the repo and build it yourself.

Looking into publishing binaries.

## How to use it?

### CLI

`./jsongpt once -k <openai_api_key> <template>`

you can omit the `-k` flag if you have the `OPENAI_API_KEY` environment variable set.

***Flags***

`-k` OpenAI API Key

`-m` Model to use (default: gpt-3.5-turbo)

`-l` How many objects to generate (default: 1)

`-L` Language (default: english)

### Server

`./jsongpt server -p <port>`

***Flags***

`-p` Port to listen on (default: 8080)

`-k` OpenAI API Key

you can omit the `-k` flag if you have the `OPENAI_API_KEY` environment variable set.

***Endpoints***

- `POST: /`:

  - body:

    ```json
    {
        "gpt_model": "<model to use>",
        "max_tokens": "<max tokens to generate>",
        "language": "<language to use>",
        "length": "<how many objects to generate>",
        "system_prompt": "<system prompt to use>",
        "model": "<your actual json object to mock>"
    }
    ```

    all optional except `model`

  - response:

      kinda your model :/

## Example Request

```json
{
    "model": {
        "id": "random v4 uuid",
        "fullname": "a random person name",
        "email": "a random email",
        "phone": "a random phone number",
        "bio": "30-40 words of social media bio",
        "username": "a random username that matches [a-zA-Z0-9_]{3,15}",
    }
}
```
