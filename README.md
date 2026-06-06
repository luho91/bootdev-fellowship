# bootdev-fellowship
A tool to hack the bootdev fellowship achievement for those who can't afford it or don't have any friends.
## Usage
### Obtaining your access token
There are three ways to obtain the token.
1. If you have the Bootdev CLI installed and are logged in, there will be a `.bootdev.yaml` file in your home directory. This will hold an `access_token` (among other interesting things). bootdev-fellowship can try and read that file and get the token from there.
2. You can manually copy and paste a token. This is useful if something is wrong with the first option.
    1. You can find your access token (again) in your `.bootdev.yaml`.
    2. As a fallback, you can inspect a submit request in your browser (while being logged in) and read the `Authorization` header, it will start with `Bearer `.
3. WIP: You will be able to login using email and password.
