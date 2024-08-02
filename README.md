# vcs

A simple version control system written in Go with the [Cobra](https://github.com/spf13/cobra) library for creating CLI applications. Inspired by [Git](https://github.com/git/git) and [Build Your Own VCS](https://ryanheathcote.com/git/build-your-own-vcs) by Ryan Heathcote.

# features

### `init`

Creates the necessary directories and files for the repository, including the repository directory (`.vcs`), objects directory, and index file. If the repository already exists, a message is displayed indicating that the repository already exists.

### `add [filename]`

Allows users to add a file to the repository. It checks if the file is ignored based on the patterns specified in the .vcsignore file. If the file is not ignored, it adds the file to the index and updates the index file with the new file's information.

### `commit -m "commit message"`

Allows users to commit changes to the repository by providing a commit message.

### `status`

Displays the working tree status, showing the changes to be committed and the unstaged changes. This command also checks the .vcsignore file for ignored patterns and computes the SHA1 hash of each file to compare with the index.


# TODO

- [ ] fix status showing <file> to be committed after it is already committed

- [ ] implement unstaging changes
- [ ] implement diff

- [ ] implement better pattern handling in `.vcsignore`
- [ ] implement better pattern handling in `add`
    - e.g. `go run main.go add .` adds a file named `"."` rather than all unstaged files

- [ ] handle multiple files being added at the same time
    ```
    > go run main.go add LICENSE README.md
    File added: LICENSE
    ```

- [ ] implement commit history (git log)
