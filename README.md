# thoughtster

Tweeting is hard; you have to open a new browser tab, type `twitter.com`, navigate to _What's happening_, type out your message, and send it. Instead, wouldn't it be easier if you could just...

```bash
thoughtster --tweet "...do it from a command line"
```

And that's basically what this project is all about.

## Usage

If you want to do the same, create a new app on Twitter's developer portal and copy consumer key, consumer secret, access token and access token secret, placing them in a `.env` file at the root of the project

```bash
API_KEY=XXX
API_KEY_SECRET=XXX
ACCESS_TOKEN=XXX
ACCESS_TOKEN_SECRET=XXX
```

Since the project uses zero dependencies, just make sure you have Golang installed on your system. Create and install the executable:

```bash
go install github.com/sunderee/thoughtster
```

and verify it works

```bash
# Print command's path
which thoughtster

# Display help message
thoughtster --help
```

## License

Project is open-sourced under MIT license.
